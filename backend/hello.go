package main

import (
	"database/sql"
	"errors"
	"log"
	"net/http"
	"strings"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"golang.org/x/crypto/scrypt"
)

var db *sql.DB

type Login struct {
	Email    string `json:"email" xml:"email" form:"email" query:"email"`
	Password string `json:"password" xml:"password" form:"password" query:"password"`
}

type Register struct {
	Name            string `json:"name" xml:"name" form:"name" query:"name"`
	Email           string `json:"email" xml:"email" form:"email" query:"email"`
	Password        string `json:"password" xml:"password" form:"password" query:"password"`
	PasswordConfirm string `json:"password-confirm" xml:"password-confirm" form:"password-confirm" query:"password-confirm"`
}

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func createUser(u User) (int64, error) {
	result, err := db.Exec("INSERT user VALUES (null, ?, ?, ?)", u.Name, u.Email, u.Password)
	if err != nil {
		log.Print(err)
		return -1, err
	}
	id, err := result.LastInsertId()
	if err != nil {
		log.Print(err)
	}
	return id, err
}

func getUserByEmail(email string) (User, error) {
	var u User
	row := db.QueryRow("SELECT * FROM user WHERE Email = ?", email)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return u, errors.New("unable to find user by email")
	}
	return u, nil

}

func hashPassword(password string) (string, error) {
	salt := []byte("poopoo peepee")

	hashedPass, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)

	return string(hashedPass[:]), err
}

func loginUser(l Login) (User, error) {
	u, err := getUserByEmail(l.Email)
	if err != nil {
		return u, errors.New("login failed")
	}

	hashedPassword, err := hashPassword(l.Password)
	if err != nil {
		return u, errors.New("login failed")
	}

	if u.Password != hashedPassword {
		return u, errors.New("login failed")
	}

	return u, nil

}

func registerUser(r Register) (int64, error) {
	if r.Name == "" {
		return -1, errors.New("name cannot be empty")
	}
	if !strings.Contains(r.Email, "@") {
		return -1, errors.New("invalid email")
	}
	if r.Password != r.PasswordConfirm {
		return -1, errors.New("passwords don't match")

	}
	hashedPass, err := hashPassword(r.Password)
	if err != nil {
		return -1, errors.New("error")
	}

	u := new(User)
	u.Name = r.Name
	u.Email = r.Email
	u.Password = hashedPass

	return createUser(*u)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, WorldPoop!")
}

func login(c echo.Context) error {
	var l Login
	if err := c.Bind(l); err != nil {
		return err
	}
	u, err := loginUser(l)
	if err != nil {
		return c.JSON(http.StatusForbidden, "Login failed")
	}
	u.Password = ""
	return c.JSON(http.StatusCreated, u)
}

func logout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

func register(c echo.Context) error {
	r := new(Register)
	if err := c.Bind(r); err != nil {
		return err
	}
	userId, err := registerUser(*r)
	if err != nil {
		log.Print(err)

		return c.JSON(http.StatusNotAcceptable, err)
	} else {
		return c.JSON(http.StatusCreated, userId)
	}
}

func resetPass(c echo.Context) error {
	return c.String(http.StatusOK, "Reset pass")
}

func main() {
	connStr := "postgres://carmaint:example@postgres/carmaint"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Query("SELECT 1")

	e := echo.New()

	e.Use(middleware.Logger())

	// Routes
	e.GET("/", hello)
	e.POST("/login", login)
	e.GET("/logout", logout)
	e.POST("/register", register)
	e.POST("/resetPass", resetPass)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
