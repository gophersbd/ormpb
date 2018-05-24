package e2e

import (
	"testing"
	"time"

	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
	"os"
)

const (
	TIMEOUT     = 20 * time.Minute
	PostgresURL = "postgres_url"
)

var (
	postgresClient *sql.DB
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(TIMEOUT)

	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "e2e Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {

	postgresURL, found := os.LookupEnv(PostgresURL)
	Expect(found).Should(BeTrue())

	connStr := fmt.Sprintf("postgres://postgres@%s/?sslmode=disable", postgresURL)
	var err error
	postgresClient, err = sql.Open("postgres", connStr)
	Expect(err).ShouldNot(HaveOccurred())
})
