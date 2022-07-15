package model

import (
	"fmt"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/helper"
)

type SingUp struct {
	Email     string `json:"email"`
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
	Password  string `json:"password"`
}

func (s SingUp) Validate() error {

	if !helper.EmailCheck(s.Email) {
		return fmt.Errorf("invalid email")
	}
	if len(s.Password) < 4 {
		return fmt.Errorf("invalid password, Password should be more than 4 characters")
	}
	if len(s.FirstName) < 1 || len(s.LastName) < 1 {
		return fmt.Errorf("invalid firstname/lastname, please enter firstname/lastname")
	}

	return nil
}
