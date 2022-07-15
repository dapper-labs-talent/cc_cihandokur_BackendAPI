package controller_test

import (
	"testing"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/config"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/db"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
)

func TestController(t *testing.T) {
	config.LoadConfiguration()
	db.New()
	RegisterFailHandler(Fail)
	RunSpecs(t, "Controller Suite")
}
