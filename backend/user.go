package models

type User struct {
	modelImpl
	Username string
	FullName string
	Email    string
	Password string
}

func newUser(userName, fullName, email string) *User {
	u := &User{
		Username: userName,
		FullName: fullName,
		Email: email,
	}
	u.SetId(userName)
	return u
}

func (u *User) GetId() string {
	return u.UserName
}