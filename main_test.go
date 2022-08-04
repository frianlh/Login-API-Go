package main

import (
	handler "Login-API/handler"
	"bytes"
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/stretchr/testify/assert"
)

// Success
func TestLoginShouldBeSuccess(t *testing.T) {
	// Sample User Input
	var jsonUser = []byte(`{"username":"user1","password":"pwuser1"}`)

	request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler.Login(response, request)

	// Response Status
	assert.Equal(t, http.StatusOK, response.Code)

	// Validation Response Body - JSON
	assert.Equal(t, "application/json", response.HeaderMap.Get("Content-Type"))

	// Response Body
	assert.JSONEq(t, `{"message": "login success"}`, response.Body.String())

	// Cookies
	var token = ""
	for _, cookie := range response.Result().Cookies() {
		if cookie.Name == "Token" {
			token = cookie.Value
		}
	}
	assert.NotEqual(t, token, "")
}

// Error
func TestLoginShouldBeError(t *testing.T) {
	// Sample User Input
	var jsonUser = []byte(`{"username":"user","password":"pwuser"}`)

	request, err := http.NewRequest(http.MethodPost, "/login", bytes.NewBuffer(jsonUser))
	if err != nil {
		t.Fatal(err)
	}

	response := httptest.NewRecorder()
	handler.Login(response, request)

	// Response Status
	assert.NotEqual(t, http.StatusOK, response.Code)

	// Validation Response Body - JSON
	assert.Equal(t, "application/json", response.HeaderMap.Get("Content-Type"))

	// Response Body
	assert.JSONEq(t, `{"message": "login failed"}`, response.Body.String())

	// Cookies
	var token = ""
	for _, cookie := range response.Result().Cookies() {
		if cookie.Name == "Token" {
			token = cookie.Value
		}
	}
	assert.Equal(t, token, "")
}
