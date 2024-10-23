package internal

import (
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/sessions"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"golang.org/x/oauth2"
	"golang.org/x/oauth2/google"
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

func Login(c *CustomContext) error {
	l := new(LoginForm)
	if err := c.Bind(l); err != nil {
		return err
	}
	u, err := LoginUser(c, *l)
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

func Logout(c *CustomContext) error {
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	sess.Values["foo"] = nil

	return c.String(http.StatusOK, "Logout")
}

func Register(c *CustomContext) error {
	r := new(RegisterForm)
	if err := c.Bind(r); err != nil {
		return err
	}
	userId, err := RegisterUser(c, *r)
	if err != nil {
		log.Print(err)

		return c.JSON(http.StatusNotAcceptable, err)
	} else {
		return c.JSON(http.StatusCreated, userId)
	}
}

func ResetPass(c *CustomContext) error {
	r := new(ResetPassForm)
	if err := c.Bind(r); err != nil {
		return err
	}
	err := ResetPassUser(c, *r)
	if err != nil {
		return c.String(http.StatusOK, "Reset pass")
	} else {
		return c.String(http.StatusInternalServerError, "Reset pass")
	}
}

var googleOauthConfig = &oauth2.Config{
	ClientID:     os.Getenv("GOOGLE_AUTH_CLIENT_ID"),
	ClientSecret: os.Getenv("GOOGLE_AUTH_CLIENT_SECRET"),

	RedirectURL: os.Getenv("GOOGLE_AUTH_REDIRECT_URL"),
	Scopes: []string{
		"https://www.googleapis.com/auth/userinfo.email",
		"https://www.googleapis.com/auth/userinfo.profile",
	},
	Endpoint: google.Endpoint,
}

func OauthGoogleLogin(c *CustomContext) error {
	url := googleOauthConfig.AuthCodeURL("abc123")
	return c.Redirect(http.StatusTemporaryRedirect, url)
}

func OauthGoogleCallback(c *CustomContext) error {
	if c.FormValue("state") != "abc123" {
		return c.String(http.StatusBadRequest, "invalid oauth state")

	}

	code := c.FormValue("code")
	token, err := googleOauthConfig.Exchange(context.Background(), code)
	if err != nil {
		return c.String(http.StatusBadRequest, "Could not get token")
	}

	client := googleOauthConfig.Client(context.Background(), token)
	response, err := client.Get("https://www.googleapis.com/oauth2/v2/userinfo")
	if err != nil {
		return c.String(http.StatusInternalServerError, "Failed to get user info")
	}
	defer response.Body.Close()

	if response.StatusCode != http.StatusOK {
		return c.String(http.StatusInternalServerError, "Failed to get user info")
	}

	// Display user info
	var user struct {
		Email string `json:"email"`
		Name  string `json:"name"`
	}
	if err := json.NewDecoder(response.Body).Decode(&user); err != nil {
		return c.String(http.StatusInternalServerError, "Failed to decode user info")
	}

	return c.String(http.StatusOK, fmt.Sprintf("User: %s, Email: %s", user.Name, user.Email))

}

func MyCars(c *CustomContext) error {
	// get user id from session
	sess, err := session.Get("session", c)
	if err != nil {
		return err
	}
	userId := sess.Values["user_id"].(int)

	cars, err := GetCarsByUserId(c, userId)
	if err != nil {
		return c.JSON(http.StatusInternalServerError, err)
	}
	return c.JSON(http.StatusOK, cars)
}

type CustomContext struct {
	echo.Context
	db *sql.DB
}
