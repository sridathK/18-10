package models

import "errors"

var users = map[uint64]User{
	123: {
		FName:    "Bob",
		LName:    "abc",
		Password: "someSecretPassword",
		Email:    "bob@email.com",
	},
}

func FetchUser(userId uint64) (User, error) {
	u, ok := users[userId]
	if !ok {
		return User{}, errors.New("user not there")
	}
	return u, nil

}
