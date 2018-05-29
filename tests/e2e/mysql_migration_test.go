package e2e

import (
	"database/sql"
	"fmt"
	"io/ioutil"
	"time"

	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("MySQLMigration", func() {
	var (
		client  *sql.DB
		mTime   string
		upSQL   []byte
		downSQL []byte
	)

	BeforeEach(func() {
		client = root.MySQLClient
		mTime = time.Now().Format("20060102")
		var err error
		downSQL, err = ioutil.ReadFile(fmt.Sprintf("../../examples/mysql/migrations/%s_examples_down.sql", mTime))
		Expect(err).ShouldNot(HaveOccurred())
		upSQL, err = ioutil.ReadFile(fmt.Sprintf("../../examples/mysql/migrations/%s_examples_up.sql", mTime))
		Expect(err).ShouldNot(HaveOccurred())
	})

	Describe("Test", func() {
		BeforeEach(func() {
			client.Exec(string(downSQL))
		})

		AfterEach(func() {
			client.Exec(string(downSQL))
		})

		Context("For Create", func() {
			It("with example_up", func() {
				_, err := client.Exec(string(upSQL))
				Expect(err).ShouldNot(HaveOccurred())

				tableName := "examples"

				_, err = client.Exec(fmt.Sprintf("INSERT INTO %s (name, email) VALUES (%s, %s)", tableName, "'shahriar'", "'shahriar052@gmail.com'"))
				Expect(err).ShouldNot(HaveOccurred())

				var name string
				var point float64
				err = client.QueryRow(fmt.Sprintf("SELECT name,point from %s WHERE user_id = %d", tableName, 1)).Scan(&name, &point)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(name).Should(Equal("shahriar"))
				Expect(point).Should(Equal(17.33))
			})
		})

		Context("For Delete", func() {
			It("with example_down", func() {
				_, err := client.Exec(string(upSQL))
				Expect(err).ShouldNot(HaveOccurred())

				tableName := "examples"

				_, err = client.Query(fmt.Sprintf("SELECT * from %s", tableName))
				Expect(err).ShouldNot(HaveOccurred())

				_, err = client.Exec(string(downSQL))
				Expect(err).ShouldNot(HaveOccurred())

				_, err = client.Query(fmt.Sprintf("SELECT * from %s", tableName))
				Expect(err.Error()).Should(HaveSuffix("doesn't exist"))
			})
		})
	})
})
