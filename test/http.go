package test

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/config"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/router"
	. "github.com/onsi/gomega"
)

var (
	testRouter = router.RegisterRoutes()
)

func GetWithToken(url string, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("GET", url, nil)
	req.Header.Set("Content-Type", "application/json")
	setToken(req, token)
	return serve(req)
}

func PostWithoutToken(url, data string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("POST", url, bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/json")
	return serve(req)
}

func PutWithToken(url, data string, token string) *httptest.ResponseRecorder {
	req := httptest.NewRequest("PUT", url, bytes.NewBufferString(data))
	req.Header.Set("Content-Type", "application/json")
	setToken(req, token)
	return serve(req)
}

func ParseJSON(resp *httptest.ResponseRecorder, code int) map[string]interface{} {

	res := make(map[string]interface{})
	if len(resp.Body.Bytes()) > 0 {
		err := json.Unmarshal(resp.Body.Bytes(), &res)
		Expect(err).NotTo(HaveOccurred())
	}
	Expect(resp.Code).To(Equal(code))
	return res
}

func serve(req *http.Request) *httptest.ResponseRecorder {
	resp := httptest.NewRecorder()
	testRouter.ServeHTTP(resp, req)
	return resp
}

func setToken(req *http.Request, token string) {
	if len(token) == 0 {
		return
	}
	req.Header.Set(config.Config.Jwt.Header, token)
}
