package main

import (
	"database/sql"
	"errors"
	"log"
)

type User struct {
	ID       string
	Name     string
	Email    string
	Password string
}

func createUser(db *sql.DB, u User) (int64, error) {
	result, err := db.Exec("INSERT INTO Users (Name, Email, Password) VALUES ($1, $2, $3)", u.Name, u.Email, "fuck")
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

func getUserByEmail(db *sql.DB, email string) (User, error) {
	var u User
	row := db.QueryRow("SELECT * FROM users WHERE Email = ?", email)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return u, errors.New("unable to find user by email")
	}
	return u, nil

}
