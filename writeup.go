package main

import (
	"encoding/json"
	"net/http"
)

type Writeup struct {
	WID          string `json:"wid"`
	Username     string `json:"username"`
	Title        string `json:"title"`
	Content      string `json:"content"`
	Resource     string `json:"resource"`
	ExternalLink string `json:"exlink"`
}

var wps []Writeup

func GetWriteups(w http.ResponseWriter, r *http.Request) {
	json.NewEncoder(w).Encode(wps) // encode wps to w
}

func AddWriteup(w http.ResponseWriter, r *http.Request) {
	var newWriteup Writeup
	_ = json.NewDecoder(r.Body).Decode(&newWriteup)
	wps = append(wps, newWriteup)
	json.NewEncoder(w).Encode(newWriteup)
}
