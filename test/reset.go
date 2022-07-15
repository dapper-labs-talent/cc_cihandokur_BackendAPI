package test

import (
	"fmt"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/db"
	. "github.com/onsi/gomega"
)

func ResetAll() {
	ResetUserTable()
}

func ResetUserTable() {

	fmt.Println("called ResetUserTable")

	cmd := "TRUNCATE users;"
	res := db.DB.Exec(cmd)
	Expect(res.Error).NotTo(HaveOccurred())
}
