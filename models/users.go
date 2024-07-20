package models

import (
	"database/sql"

	"github.com/dennisferdian9/golang-sqlite/config"
)

type Users struct {
	Username string `json:"username" form:"username"`
	Name     string `json:"name" form:"name"`
}

// Mock data for tasks
var UsersData = []Users{
	{Username: "dennisferdian", Name: "Dennis Ferdian"},
	{Username: "deny", Name: "Deny"},
	{Username: "akmal", Name: "Akmal"},
	{Username: "firly", Name: "Firly"},
	{Username: "destia", Name: "Destia"},
	{Username: "rino", Name: "Rino"},
}

func GetUsers() ([]Users, error) {
	query := `SELECT username, name FROM users`
	rows, err := config.DB.Query(query)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []Users
	for rows.Next() {
		var user Users
		if err := rows.Scan(&user.Username, &user.Name); err != nil {
			return nil, err
		}
		users = append(users, user)
	}
	return users, nil
}

func PostUser(username string, name string) (string, error) {
	query := `INSERT INTO users (username, name) VALUES (?, ?)`
	_, err := config.DB.Exec(query, username, name)
	if err != nil {
		return "Error Update", err
	}
	return "Update Success", nil

}

func GetOneUsers(username string) (*Users, error) {
	query := `SELECT username, name FROM users WHERE username = ?`
	row := config.DB.QueryRow(query, username)

	var user Users
	err := row.Scan(&user.Username, &user.Name)

	if err != nil {
		if err == sql.ErrNoRows {
			return nil, nil // No user found
		}
		return nil, err
	}

	return &user, nil
}
