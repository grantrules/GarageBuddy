package main

import (
	"database/sql"
	"log"
	"net/http"

	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	_ "github.com/lib/pq"
)

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

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, WorldPoop!")
}

func login(c echo.Context) error {
	u := new(Login)
	if err := c.Bind(u); err != nil {
		return err
	}
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
	return c.JSON(http.StatusCreated, r)
}

func resetPass(c echo.Context) error {
	return c.String(http.StatusOK, "Reset pass")
}

func main() {

	connStr := "postgres://carmaint:example@postgres/carmaint"
	db, err := sql.Open("postgres", connStr)
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
