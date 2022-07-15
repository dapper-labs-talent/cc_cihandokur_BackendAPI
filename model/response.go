package model

type Response struct {
	Token string `json:"token"`
}

type UserList struct {
	Users []User `json:"users"`
}
