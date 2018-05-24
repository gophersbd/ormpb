package e2e

import (
	"fmt"
	"io/ioutil"
	"time"

	"github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	. "github.com/onsi/gomega"
)

var _ = Describe("PostgresMigration", func() {
	var (
		mTime = time.Now().Format("20060102")
	)

	Describe("Test", func() {
		BeforeEach(func() {
			sql, err := ioutil.ReadFile(fmt.Sprintf("migration/postgres/%s_example_down.sql", mTime))
			Expect(err).ShouldNot(HaveOccurred())
			_, err = postgresClient.Exec(string(sql))
		})

		Context("For Create", func() {
			It("with example_up", func() {
				sql, err := ioutil.ReadFile(fmt.Sprintf("migration/postgres/%s_example_up.sql", mTime))
				Expect(err).ShouldNot(HaveOccurred())
				_, err = postgresClient.Exec(string(sql))
				Expect(err).ShouldNot(HaveOccurred())

				tableName := pq.QuoteIdentifier("examples")

				_, err = postgresClient.Exec(fmt.Sprintf("INSERT INTO %s(name, email) VALUES ($1, $2)", tableName), "shahriar", "shahriar052@gmail.com")
				Expect(err).ShouldNot(HaveOccurred())

				var name string
				err = postgresClient.QueryRow(fmt.Sprintf("SELECT name from %s WHERE user_id = $1", tableName), 1).Scan(&name)
				Expect(err).ShouldNot(HaveOccurred())

				Expect(name).Should(Equal("shahriar"))
			})
		})
	})
})
