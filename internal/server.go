package internal

import (
	"log"

	"github.com/grantrules/garagebuddy/internal/utils"
	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/boj/redistore.v1"
)

// write a function that returns a new handler that handles CustomContext
func CustomContextHandler(h func(*CustomContext) error) echo.HandlerFunc {
	return func(c echo.Context) error {
		cc := c.(*CustomContext)
		return h(cc)
	}
}

func StartServer() {
	db, err := utils.ConnectDB()
	if err != nil {
		log.Fatal(err)
	}

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
	e.GET("/", Hello)
	e.GET("/test", Test)

	e.POST("/login", CustomContextHandler(Login))
	e.GET("/logout", CustomContextHandler(Logout))
	e.POST("/register", CustomContextHandler(Register))
	e.POST("/resetPass", CustomContextHandler(ResetPass))

	e.GET("/oauth2/login/google", CustomContextHandler(OauthGoogleLogin))
	e.GET("/oauth2/callback/google", CustomContextHandler(OauthGoogleCallback))

	e.GET("/mycars", CustomContextHandler(MyCars))

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
