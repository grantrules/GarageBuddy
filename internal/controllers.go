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

func ResetPassUser(cc *CustomContext, r ResetPassForm) error {
	u, err := GetUserByEmail(cc.db, r.Email)
	if err != nil {
		return errors.New("couldn't find user by email")
	}
	resetKey := utils.RandomString(128)

	// insert reset key into db table reset_password_tokens
	_, err = cc.db.Exec("INSERT INTO reset_password_tokens (user_id, reset_key) VALUES ($1, $2)", u.ID, resetKey)

	if err != nil {
		return errors.New("failed to insert reset key into db")
	}

	// send email to user with reset key
	err = utils.SendEmail(u.Email, "Password Reset", "Click here to reset your password: http://localhost:8081/reset/"+resetKey)

	return err
}

func GetCarsByUserId(cc *CustomContext, userId int) ([]Car, error) {

	cars, err := ListCarsByUserId(cc.db, userId)
	if err != nil {
		return nil, err
	}
	return cars, nil
}

func CreateCarController(cc *CustomContext, cf CarForm, userId int) (int64, error) {
	c := new(Car)
	c.Name = cf.Name
	c.Make = cf.Make
	c.Model = cf.Model
	c.Year = cf.Year
	c.UserID = userId

	return CreateCar(cc.db, *c)
}
