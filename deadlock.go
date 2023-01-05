package main

import (
	"database/sql"
	"fmt"
	"log"
	"os"

	_ "github.com/go-sql-driver/mysql"
)

func main() {
	db, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	defer db.Close()

	if err = DropTableIfExists(db); err != nil {
		log.Fatal(err)
	}
	if err = CreateTable(db); err != nil {
		log.Fatal(err)
	}
	if err = InsertSeed(db); err != nil {
		log.Fatal(err)
	}

	rows, err := db.Query("SELECT * FROM t1")
	if err != nil {
		log.Fatal(err)
	}
	defer rows.Close()

	// SQLの実行
	for rows.Next() {
		fmt.Println(rows)
	}
}

func DropTableIfExists(db *sql.DB) error {
	bytes, err := os.ReadFile("t1_drop.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	return err
}

func CreateTable(db *sql.DB) error {
	bytes, err := os.ReadFile("t1.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	return err
}

func InsertSeed(db *sql.DB) error {
	bytes, err := os.ReadFile("seed.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	return err
}
