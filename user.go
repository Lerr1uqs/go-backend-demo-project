package main

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/google/uuid"
)

type UserInput struct {
	Username string `json:"username"`
	Password string `json:"password"`
}

type User struct {
	UserId   string `json:"userid"`
	Username string `json:"username"`
	Password string `json:"password"`
}

var users []User

func AddUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[D] AddUser")
	// TODO: user 字段合法性
	var inputUser UserInput
	_ = json.NewDecoder(r.Body).Decode(&inputUser)

	userID := uuid.New().String()
	var user = User{
		UserId:   userID,
		Username: inputUser.Username,
		Password: inputUser.Password,
	}

	// users = append(users, inputUser)
	// json.NewEncoder(w).Encode(inputUser)
	// TODO: https return OK
	// if db == nil {
	// 	fmt.Println("[E] db is nil")
	// 	return
	// }

	if err := InsertUser(user); err != nil {
		log.Fatalf("%v", err)
	}

	w.WriteHeader(http.StatusOK)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	fmt.Println("[D] GetUser")
	json.NewEncoder(w).Encode(users)
}
