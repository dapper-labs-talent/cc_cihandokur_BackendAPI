package test

import (
	"encoding/json"

	. "github.com/onsi/gomega"
)

func ToJson(obj interface{}) (str string) {
	strByte, err := json.Marshal(obj)
	Expect(err).NotTo(HaveOccurred())
	return string(strByte)
}
