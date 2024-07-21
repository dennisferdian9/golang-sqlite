package main

import (
	"bytes"
	"encoding/json"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/DATA-DOG/go-sqlmock"

	"github.com/dennisferdian9/golang-sqlite/config"
	"github.com/dennisferdian9/golang-sqlite/models"
	router "github.com/dennisferdian9/golang-sqlite/routes"
	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

func TestPingRoute(t *testing.T) {
	router := router.SetupRouter()

	w := httptest.NewRecorder()
	req, _ := http.NewRequest("GET", "/ping", nil)
	router.ServeHTTP(w, req)

	pingResponse := gin.H{
		"message": "pong",
	}
	pingResponseJson, _ := json.Marshal(pingResponse)

	assert.Equal(t, 200, w.Code)
	assert.Equal(t, string(pingResponseJson), w.Body.String())
}

// func TestgetUser(t *testing.T)  {
// 	db, mock, err := sqlmock.New()
// 	assert.NoError(t, err)
// 	defer db.Close()

// 	// rows := sqlmock.NewRows([]string{"username", "name"}).
// 	// 	AddRow("dennis", "Dennis Ferdian")
// }

func TestPostUserTest(t *testing.T) {
	// chec sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	originalDB := config.DB
	config.DB = db
	defer func() { config.DB = originalDB }()

	// rows := sqlmock.NewRows([]string{"username", "name"}).
	// 	AddRow("dennis", "Dennis Ferdian")
	username := "test_name"
	name := "test name"

	// Create an example user for testing
	mock.ExpectExec("INSERT INTO users \\(username, name\\) VALUES \\(\\?, \\?\\)").
		WithArgs(username, name).
		WillReturnResult(sqlmock.NewResult(1, 1))

	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()

	exampleUser := []byte(`{"username":"test_name","name":"test name"}`)
	// userJson, _ := json.Marshal(exampleUser)
	req, err := http.NewRequest(http.MethodPost, "/api/user", bytes.NewBuffer(exampleUser))
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()

	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusCreated, w.Code)
	assert.Contains(t, w.Body.String(), "Post Success")

	// Compare the response body with the json data of exampleUser
	// responseMessage := gin.H{
	// 	"message": "Post Success",
	// 	"data":    &exampleUser,
	// }
	// responseMessageJson, _ := json.Marshal(responseMessage)

	// t.Log("Response Body:", w.Body.String())

	// assert.Equal(t, string(responseMessageJson), w.Body.String())

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetUsers(t *testing.T) {
	// chec sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	originalDB := config.DB
	config.DB = db
	defer func() { config.DB = originalDB }()

	var userDB []models.Users
	user := models.Users{
		Username: "dennis",
		Name:     "Dennis Ferdian",
	}

	rows := sqlmock.NewRows([]string{"username", "name"}).
		AddRow(user.Username, user.Name)
	// Create an example user for testing
	mock.ExpectQuery("SELECT username, name FROM users").
		WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()

	// userJson, _ := json.Marshal(exampleUser)
	req, err := http.NewRequest(http.MethodGet, "/api/user", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	userDB = append(userDB, user)

	responseMessageJson, _ := json.Marshal(userDB)

	assert.Contains(t, w.Body.String(), string(responseMessageJson))

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}

func TestGetOneUser(t *testing.T) {
	// chec sqlmock
	db, mock, err := sqlmock.New()
	assert.NoError(t, err)
	defer db.Close()

	originalDB := config.DB
	config.DB = db
	defer func() { config.DB = originalDB }()

	var userDB []models.Users
	user := models.Users{
		Username: "dennis",
		Name:     "Dennis Ferdian",
	}

	rows := sqlmock.NewRows([]string{"username", "name"}).
		AddRow(user.Username, user.Name)
	// Create an example user for testing
	mock.ExpectQuery("SELECT username, name FROM users").
		WillReturnRows(rows)

	gin.SetMode(gin.TestMode)
	router := router.SetupRouter()

	// userJson, _ := json.Marshal(exampleUser)
	req, err := http.NewRequest(http.MethodGet, "/api/user", nil)
	assert.NoError(t, err)
	req.Header.Set("Content-Type", "application/json")

	w := httptest.NewRecorder()
	router.ServeHTTP(w, req)

	assert.Equal(t, http.StatusOK, w.Code)
	userDB = append(userDB, user)

	responseMessageJson, _ := json.Marshal(userDB)

	assert.Contains(t, w.Body.String(), string(responseMessageJson))

	err = mock.ExpectationsWereMet()
	assert.NoError(t, err)
}
