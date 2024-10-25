package internal

import (
	"database/sql"
	"errors"
	"log"
)

func CreateUser(db *sql.DB, u User) (int64, error) {
	var id int64
	err := db.QueryRow("INSERT INTO Users (name, email, password) VALUES ($1, $2, $3) RETURNING id", u.Name, u.Email, u.Password).Scan(&id)
	if err != nil {
		log.Print(err)
		return -1, err
	}
	return id, err
}

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var u User
	row := db.QueryRow("SELECT * FROM users WHERE Email = $1", email)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return u, errors.New("unable to find user by email")
	}
	return u, nil

}
