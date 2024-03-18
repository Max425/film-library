package domain

import (
	"fmt"
)

type User struct {
	id       int
	name     string
	mail     string
	password string
	salt     string
	role     int
}

func NewUser(id int, name, mail, password, salt string, role int) (*User, error) {
	if id < 0 {
		return nil, fmt.Errorf("invalid user ID: %d", id)
	}

	if name == "" {
		return nil, fmt.Errorf("%w: title is required", ErrRequired)
	}

	if mail == "" {
		return nil, fmt.Errorf("%w: title is required", ErrRequired)
	}

	if password == "" {
		return nil, fmt.Errorf("%w: title is required", ErrRequired)
	}

	if role < 0 {
		return nil, fmt.Errorf("invalid user role: %d", role)
	}

	return &User{
		id:       id,
		name:     name,
		mail:     mail,
		password: password,
		salt:     salt,
		role:     role,
	}, nil
}

func (u *User) Sanitize() {
	u.password = ""
	u.salt = ""
}

func (u *User) ID() int {
	return u.id
}

func (u *User) Name() string {
	return u.name
}

func (u *User) Mail() string {
	return u.mail
}

func (u *User) Password() string {
	return u.password
}

func (u *User) Salt() string {
	return u.salt
}

func (u *User) Role() int {
	return u.role
}

func (u *User) SetSalt(salt string) {
	u.salt = salt
}

func (u *User) SetPassword(password string) {
	u.password = password
}
