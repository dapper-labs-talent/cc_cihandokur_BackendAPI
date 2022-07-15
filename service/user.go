package service

import (
	"errors"
	"fmt"
	"log"
	"time"

	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/db"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/helper"
	"github.com/dapper-labs-talent/cc_cihandokur_BackendAPI/model"
	"gorm.io/gorm"
)

type UserService struct{}

func (u UserService) CreateNewUser(singup *model.SingUp) (resp model.Response, err error) {

	err = singup.Validate()
	if err != nil {
		return
	}

	user := model.User{}

	if result := db.DB.Where("email = ?", singup.Email).First(&user); !errors.Is(result.Error, gorm.ErrRecordNotFound) {
		err = fmt.Errorf("user already exists. ")
		return
	}

	user.Password = helper.HashString(&singup.Password)
	user.Email = singup.Email
	user.FirstName = singup.FirstName
	user.LastName = singup.LastName
	user.CreatedDate = time.Now()
	user.UpdatedDate = time.Now()

	token, err := helper.GenerateJWT(singup.Email)
	if err != nil {
		return
	}

	user.LatestJwtToken = token

	if result := db.DB.Create(&user); result.Error != nil {
		err = result.Error
		return
	}

	resp.Token = token

	return resp, nil
}

func (u UserService) Login(login *model.Login) (resp model.Response, err error) {

	err = login.Validate()
	if err != nil {
		return
	}

	user := model.User{}

	if result := db.DB.Where("email = ?", login.Email).First(&user); result.Error != nil && result.RowsAffected < 1 {
		log.Println(result.Error.Error(), result.RowsAffected)
		err = fmt.Errorf("user not found. ")
		return
	}

	passMatch := helper.CheckStringHash(login.Password, user.Password)

	if !passMatch {
		err = fmt.Errorf("incorrect credentials ")
		return
	}

	token, err := helper.GenerateJWT(login.Email)

	if err != nil {
		return
	}

	user.LatestJwtToken = token
	user.UpdatedDate = time.Now()

	if result := db.DB.Save(&user); result.Error != nil {
		err = result.Error
		return
	}

	resp.Token = token

	return resp, nil
}

func (u UserService) GetUsers() (resp model.UserList, err error) {

	userList := model.UserList{}

	db.DB.Find(&userList.Users)

	return userList, nil
}

func (u UserService) Update(email string, update *model.UserUpdate) error {

	err := update.Validate()
	if err != nil {
		return err
	}

	if !helper.EmailCheck(email) {
		return fmt.Errorf("invalid email. ")
	}

	user := model.User{}

	if result := db.DB.Where("email = ?", email).First(&user); result.Error != nil && result.RowsAffected < 1 {
		return fmt.Errorf("user not found. ")
	}

	user.FirstName = update.FirstName
	user.LastName = update.LastName

	if result := db.DB.Save(&user); result.Error != nil {
		return result.Error
	}

	return nil
}
