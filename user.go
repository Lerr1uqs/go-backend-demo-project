package main

import (
	"encoding/json"
	"net/http"
)

type User struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User

func AddUser(w http.ResponseWriter, r *http.Request) {
	// 鉴权
	var newUser User
	_ = json.NewDecoder(r.Body).Decode(&newUser)
	users = append(users, newUser)
	json.NewEncoder(w).Encode(newUser)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(users)
}
