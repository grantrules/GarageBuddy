package internal

import (
	"database/sql"
	"log"

	_ "github.com/jackc/pgx/v5/stdlib"
	"github.com/labstack/echo-contrib/session"
	"github.com/labstack/echo/v4"
	"github.com/labstack/echo/v4/middleware"
	"gopkg.in/boj/redistore.v1"
)

func StartServer() {
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
	e.GET("/", Hello)
	e.GET("/test", Test)

	e.POST("/login", Login)
	e.GET("/logout", Logout)
	e.POST("/register", Register)
	e.POST("/resetPass", ResetPass)

	// Start server
	e.Logger.Fatal(e.Start(":8080"))
}
