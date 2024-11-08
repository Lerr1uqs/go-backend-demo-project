package main

import (
	"database/sql"

	_ "github.com/mattn/go-sqlite3"

	// _ "database/sql"
	"fmt"
	"log"
)

func CreateDB() {
	// 连接到 SQLite 数据库（如果不存在则创建）
	db, err := sql.Open("sqlite3", "mydb.db")
	if err != nil {
		log.Fatal(err)
	}
	defer db.Close()

	// 创建表
	createTableSQL := `CREATE TABLE IF NOT EXISTS writeups (
        wid TEXT PRIMARY KEY,
        username TEXT NOT NULL,
        title TEXT NOT NULL,
        content TEXT NOT NULL,
        resource TEXT,
        exlink TEXT
    );`

	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}

	createTableSQL = `CREATE TABLE IF NOT EXISTS users (
        userid TEXT PRIMARY KEY,
        username TEXT NOT NULL,
        password TEXT NOT NULL
    );`
	
	_, err = db.Exec(createTableSQL)
	if err != nil {
		log.Fatalf("Failed to create table: %v", err)
	}


	fmt.Println("Table created successfully.")
}
