package main

import (
	"database/sql"
	"encoding/base64"
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

func HashPassword(password string) (string, error) {
	salt := []byte("poopoo peepee")

	hashedPass, err := scrypt.Key([]byte(password), salt, 32768, 8, 1, 32)

	encodedStr := base64.StdEncoding.EncodeToString(hashedPass)

	return encodedStr, err
}

func loginUser(cc *CustomContext, l Login) (User, error) {
	u, err := getUserByEmail(cc.db, l.Email)
	if err != nil {
		return u, errors.New("login failed")
	}

	hashedPassword, err := HashPassword(l.Password)
	if err != nil {
		return u, errors.New("login failed")
	}

	if u.Password != hashedPassword {
		return u, errors.New("login failed")
	}

	return u, nil

}

func registerUser(cc *CustomContext, r Register) (int64, error) {
	if r.Name == "" {
		return -1, errors.New("name cannot be empty")
	}
	if !strings.Contains(r.Email, "@") {
		return -1, errors.New("invalid email")
	}
	if r.Password != r.PasswordConfirm {
		return -1, errors.New("passwords don't match")

	}
	hashedPass, err := HashPassword(r.Password)
	if err != nil {
		return -1, errors.New("error")
	}

	u := new(User)
	u.Name = r.Name
	u.Email = r.Email
	u.Password = hashedPass

	return createUser(cc.db, *u)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, WorldPoop!")
}

func login(c echo.Context) error {
	cc := c.(*CustomContext)
	var l Login
	if err := c.Bind(l); err != nil {
		return err
	}
	u, err := loginUser(cc, l)
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
	cc := c.(*CustomContext)
	r := new(Register)
	if err := c.Bind(r); err != nil {
		return err
	}
	userId, err := registerUser(cc, *r)
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

type CustomContext struct {
	echo.Context
	db *sql.DB
}

func main() {
	connStr := "postgres://carmaint:example@postgres/carmaint"
	db, err := sql.Open("pgx", connStr)
	if err != nil {
		log.Fatal(err)
	}

	db.Query("SELECT 1")

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, db}
			return next(cc)
		}
	})
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
