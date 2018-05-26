package e2e

import (
	"testing"
	"time"

	"github.com/gophersbd/ormpb/tests/e2e/helper"
	_ "github.com/lib/pq"
	. "github.com/onsi/ginkgo"
	"github.com/onsi/ginkgo/reporters"
	. "github.com/onsi/gomega"
)

const (
	TIMEOUT = 20 * time.Minute
)

var (
	root *helper.Framework
)

func TestE2e(t *testing.T) {
	RegisterFailHandler(Fail)
	SetDefaultEventuallyTimeout(TIMEOUT)

	junitReporter := reporters.NewJUnitReporter("junit.xml")
	RunSpecsWithDefaultAndCustomReporters(t, "e2e Suite", []Reporter{junitReporter})
}

var _ = BeforeSuite(func() {
	pgClient, err := helper.GetPostgresClient()
	Expect(err).ShouldNot(HaveOccurred())
	msClient, err := helper.GetMySQLClient()
	Expect(err).ShouldNot(HaveOccurred())
	root = helper.New(pgClient, msClient)
})
