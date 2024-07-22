package models

import (
	"database/sql"

	"github.com/dennisferdian9/golang-sqlite/config"
)

type Users struct {
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
	Password string `json:"password" form:"password"`
}

func GetUsers() ([]Users, error) {
	query := `SELECT username, name, password FROM users`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Users
	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.Username, &user.Name, &user.Password); err != nil {
			return nil, err
		}

		users = append(users, user)
	}
	return users, nil
}

func PostUser(username string, name string, password string) (string, error) {
	query := `INSERT INTO users (username, name, password) VALUES (?, ?, ?)`
	println(password)
	_, err := config.DB.Exec(query, username, name, password)
	println(err)

	if err != nil {
		return "Error Update", err
	}
	return "Update Success", nil

}

func GetOneUsers(username string) (*Users, error) {
	query := `SELECT username, name, password FROM users WHERE username = ?`
	row := config.DB.QueryRow(query, username)

	var user Users
	err := row.Scan(&user.Username, &user.Name, &user.Password)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}

	return &user, nil
}
