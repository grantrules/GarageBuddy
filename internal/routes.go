package internal

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/sessions"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
)

func Hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, WorldPoop!")
}

func Test(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("foo=%v\n", sess.Values["foo"]))
}

func Login(c echo.Context) error {
	cc := c.(*CustomContext)
	l := new(LoginForm)
	if err := c.Bind(l); err != nil {
		return err
	}
	u, err := LoginUser(cc, *l)
	if err != nil {
		log.Print(err)
		return c.JSON(http.StatusForbidden, "Login failed")
	}

	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Options = &sessions.Options{
		Path:     "/",
		MaxAge:   86400 * 7,
		HttpOnly: true,
	}
	sess.Values["foo"] = "bar"
	if err := sess.Save(c.Request(), c.Response()); err != nil {
		return err
	}

	u.Password = ""
	return c.JSON(http.StatusCreated, u)
}

func Logout(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["foo"] = nil

	return c.String(http.StatusOK, "Logout")
}

func Register(c echo.Context) error {
	cc := c.(*CustomContext)
	r := new(RegisterForm)
	if err := c.Bind(r); err != nil {
		return err
	}
	userId, err := RegisterUser(cc, *r)
	if err != nil {
		log.Print(err)

		return c.JSON(http.StatusNotAcceptable, err)
	} else {
		return c.JSON(http.StatusCreated, userId)
	}
}

func ResetPass(c echo.Context) error {
	return c.String(http.StatusOK, "Reset pass")
}

type CustomContext struct {
	echo.Context
	db *sql.DB
}
