package model

import (
	"fmt"
	"time"
)

// db Entity
type User struct {
	Id             int64     `json:"-"`
	Email          string    `json:"email" gorm:"unique"`
	FirstName      string    `json:"firstname"`
	LastName       string    `json:"lastname"`
	Password       string    `json:"-"`
	LatestJwtToken string    `json:"-"`
	CreatedDate    time.Time `json:"-"`
	UpdatedDate    time.Time `json:"-"`
}

type UserUpdate struct {
	FirstName string `json:"firstname"`
	LastName  string `json:"lastname"`
}

func (s UserUpdate) Validate() error {

	if len(s.FirstName) < 1 || len(s.LastName) < 1 {
		return fmt.Errorf("invalid firstname/lastname, please enter firstname/lastname")
	}

	return nil
}
