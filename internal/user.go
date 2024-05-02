package internal

import (
	"database/sql"
	"errors"
	"log"
)

func CreateUser(db *sql.DB, u User) (int64, error) {
	result, err := db.Exec("INSERT INTO Users (Name, Email, Password) VALUES ($1, $2, $3)", u.Name, u.Email, u.Password)
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

func GetUserByEmail(db *sql.DB, email string) (User, error) {
	var u User
	row := db.QueryRow("SELECT * FROM users WHERE Email = $1", email)
	if err := row.Scan(&u.ID, &u.Name, &u.Email, &u.Password); err != nil {
		return u, errors.New("unable to find user by email")
	}
	return u, nil

}
