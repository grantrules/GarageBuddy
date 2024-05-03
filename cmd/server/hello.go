package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"
	"net/http"
	"strings"

	"github.com/gorilla/sessions"
	"github.com/grantrules/garagebuddy/internal"
	"github.com/grantrules/garagebuddy/internal/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/boj/redistore.v1"
)

func loginUser(cc *CustomContext, l internal.LoginForm) (internal.User, error) {
	u, err := internal.GetUserByEmail(cc.db, l.Email)
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

func registerUser(cc *CustomContext, r internal.RegisterForm) (int64, error) {
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

	u := new(internal.User)
	u.Name = r.Name
	u.Email = r.Email
	u.Password = hashedPass

	return internal.CreateUser(cc.db, *u)
}

func hello(c echo.Context) error {
	return c.String(http.StatusOK, "Hello, WorldPoop!")
}

func test(c echo.Context) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	return c.String(http.StatusOK, fmt.Sprintf("foo=%v\n", sess.Values["foo"]))
}

func login(c echo.Context) error {
	cc := c.(*CustomContext)
	l := new(internal.LoginForm)
	if err := c.Bind(l); err != nil {
		return err
	}
	u, err := loginUser(cc, *l)
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

func logout(c echo.Context) error {
	return c.String(http.StatusOK, "Logout")
}

func register(c echo.Context) error {
	cc := c.(*CustomContext)
	r := new(internal.RegisterForm)
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

	store, err := redistore.NewRediStore(10, "tcp", "redis:6379", "", []byte("secret-key"))
	if err != nil {
		log.Fatalf("redis couldn't connect: %s", err)
	}

	e := echo.New()
	e.Use(func(next echo.HandlerFunc) echo.HandlerFunc {
		return func(c echo.Context) error {
			cc := &CustomContext{c, db}
			return next(cc)
		}
	})

	e.Use(middleware.Logger())
	e.Use(session.Middleware(store))

	// Routes
	e.GET("/", hello)
	e.GET("/test", test)

	e.POST("/login", login)
	e.GET("/logout", logout)
	e.POST("/register", register)
	e.POST("/resetPass", resetPass)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
