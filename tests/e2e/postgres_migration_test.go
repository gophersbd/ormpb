package e2e

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	"github.com/lib/pq"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PostgresMigration", func() {
	var (
		client  *sql.DB
		mTime   string
		upSQL   []byte
		downSQL []byte
	)

	BeforeEach(func() {
		client = root.PostgresClient
		mTime = time.Now().Format("20060102")
		var err error
		downSQL, err = ioutil.ReadFile(fmt.Sprintf("../../examples/postgres/migrations/%s_example_down.sql", mTime))
		Expect(err).ShouldNot(HaveOccurred())
		upSQL, err = ioutil.ReadFile(fmt.Sprintf("../../examples/postgres/migrations/%s_example_up.sql", mTime))
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Test", func() {
		BeforeEach(func() {
			_, err := client.Exec(string(downSQL))
			Expect(err).ShouldNot(HaveOccurred())
		})

		AfterEach(func() {
			_, err := client.Exec(string(downSQL))
			Expect(err).ShouldNot(HaveOccurred())
		})

		Context("For Create", func() {
			It("with example_up", func() {
				_, err := client.Exec(string(upSQL))
				Expect(err).ShouldNot(HaveOccurred())

				tableName := pq.QuoteIdentifier("examples")

				_, err = client.Exec(fmt.Sprintf("INSERT INTO %s(name, email) VALUES ($1, $2)", tableName), "shahriar", "shahriar052@gmail.com")
				Expect(err).ShouldNot(HaveOccurred())

				var name string
				var point float64
				err = client.QueryRow(fmt.Sprintf("SELECT name,point from %s WHERE user_id = $1", tableName), 1).Scan(&name, &point)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(name).Should(Equal("shahriar"))
				Expect(point).Should(Equal(17.33))
			})
		})

		Context("For Delete", func() {
			It("with example_down", func() {
				_, err := client.Exec(string(upSQL))
				Expect(err).ShouldNot(HaveOccurred())

				tableName := pq.QuoteIdentifier("examples")

				_, err = client.Query(fmt.Sprintf("SELECT * from %s", tableName))
				Expect(err).ShouldNot(HaveOccurred())

				_, err = client.Exec(string(downSQL))
				Expect(err).ShouldNot(HaveOccurred())

				_, err = client.Query(fmt.Sprintf("SELECT * from %s", tableName))
				Expect(err.Error()).Should(HaveSuffix("does not exist"))
			})
		})
	})
})
