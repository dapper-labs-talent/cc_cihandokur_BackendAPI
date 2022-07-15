package helper

import (
	"log"

	"golang.org/x/crypto/bcrypt"
)

const defaultCost = 10

//CheckStringHash compares hash with password
func CheckStringHash(str, hash string) bool {
	err := bcrypt.CompareHashAndPassword([]byte(hash), []byte(str))
	return err == nil
}

//HashString hashes user password
func HashString(str *string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(*str), defaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}

//CreateHashedPassword creating hash for password
func CreateHashedPassword(str string) string {
	bytes, err := bcrypt.GenerateFromPassword([]byte(str), defaultCost)
	if err != nil {
		log.Fatal(err)
	}
	return string(bytes)
}
