package User

import (
	"backend/Database"
)

func newUser(userName, email, password string) *User {
	u := &User{
		Username: userName,
		Email: email,
		Password: password,
	}
	return u
}

