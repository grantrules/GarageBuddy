package internal

import (
	"errors"
	"strings"

	"github.com/grantrules/garagebuddy/internal/utils"
)

func LoginUser(cc *CustomContext, l LoginForm) (User, error) {
	u, err := GetUserByEmail(cc.db, l.Email)
	if err != nil {
		return u, errors.New("login failed 0 couldn't find user")
	}
	hashedPassword, err := utils.HashPassword(l.Password)
	if err != nil {
		return u, errors.New("login failed - password couldn't be hashed???")
	}

	if u.Password != hashedPassword {
		return u, errors.New("login failed - hashed passwords didn't match")
	}

	return u, nil

}

func RegisterUser(cc *CustomContext, r RegisterForm) (int64, error) {
	if r.Name == "" {
		return -1, errors.New("name cannot be empty")
	}
	if !strings.Contains(r.Email, "@") {
		return -1, errors.New("invalid email")
	}
	if r.Password != r.PasswordConfirm {
		return -1, errors.New("passwords don't match")

	}
	hashedPass, err := utils.HashPassword(r.Password)
	if err != nil {
		return -1, errors.New("error")
	}

	u := new(User)
	u.Name = r.Name
	u.Email = r.Email
	u.Password = hashedPass

	return CreateUser(cc.db, *u)
}
