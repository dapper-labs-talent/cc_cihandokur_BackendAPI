package controller_test

import (
	"net/http"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/model"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/test"
	. "github.com/onsi/ginkgo/v2"
	. "github.com/onsi/gomega"
	. "github.com/onsi/gomega/gstruct"
)

var _ = Describe("User Controller", func() {

	test.ResetUserTable()
	email := "cihandokur@gmail.com"
	pass := "Zup3rp@ss"
	var singUpData map[string]interface{}
	var loginData map[string]interface{}

	tokenKeys := Keys{
		"token": Not(BeEmpty()),
	}

	failedUserUpdateKeys := Keys{
		"Message": Equal("token contains an invalid number of segments"),
		"Status":  Equal("403"),
	}

	userData := make(map[string]interface{})
	userData["email"] = "cihandokur@gmail.com"
	userData["firstname"] = "Cihan"
	userData["lastname"] = "Dokur"

	usersContainer := Keys{
		"users": Not(BeEmpty()),
	}

	It("singup", func() {

		singUp := model.SingUp{
			Email:     email,
			FirstName: "Cihan",
			LastName:  "Dokur",
			Password:  pass,
		}

		jsonSingUp := test.ToJson(&singUp)
		resp := test.PostWithoutToken("/signup", jsonSingUp)
		singUpData = test.ParseJSON(resp, http.StatusOK)
		Expect(singUpData).To(MatchAllKeys(tokenKeys))
	})

	It("Login", func() {

		login := model.Login{
			Email:    email,
			Password: pass,
		}

		jsonLogin := test.ToJson(&login)
		resp := test.PostWithoutToken("/login", jsonLogin)
		loginData = test.ParseJSON(resp, http.StatusOK)
		Expect(loginData).To(MatchAllKeys(tokenKeys))
	})

	It("Get Users", func() {

		resp := test.GetWithToken("/users", loginData["token"].(string))
		getUserData := test.ParseJSON(resp, http.StatusOK)
		Expect(getUserData).To(MatchAllKeys(usersContainer))

		respUserData := getUserData["users"]
		Expect(respUserData).Should(ContainElement(userData))
	})

	It("Update current logged in user information", func() {

		userUpdate := model.UserUpdate{
			FirstName: "Dokur",
			LastName:  "Cihan",
		}

		jsonUserUpdate := test.ToJson(&userUpdate)
		resp := test.PutWithToken("/users", jsonUserUpdate, loginData["token"].(string))
		result := test.ParseJSON(resp, http.StatusOK)
		Expect(result).Should(BeEmpty())

		resp = test.PutWithToken("/users", jsonUserUpdate, "fakeTokenInfo")
		result = test.ParseJSON(resp, http.StatusForbidden)
		Expect(result).To(MatchAllKeys(failedUserUpdateKeys))
	})

})
