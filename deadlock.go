package main

import (
	"crypto/md5"
	"database/sql"
	"fmt"
	"log"
	"math/rand"
	"os"
	"sort"
	"strings"
	"sync"
	"time"

	_ "github.com/go-sql-driver/mysql"
)

var db *sql.DB

func main() {
	src, err := sql.Open("mysql", "root@tcp(127.0.0.1:3306)/test")
	if err != nil {
		log.Fatal("unable to use data source name", err)
	}
	db = src
	defer db.Close()

	if err = DropTableIfExists(); err != nil {
		log.Fatal(err)
	}
	if err = CreateTable(); err != nil {
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

	wg := &sync.WaitGroup{}
	for i := 0; i < 100; i++ {
		wg.Add(1)
		go BatchInsert(wg)
	}
	wg.Wait()

}

func BatchInsert(wg *sync.WaitGroup) {
	defer wg.Done()

	s := []string{}
	for i := 0; i < 1000; i++ {
		rand.Seed(time.Now().UnixNano())
		hex := md5.Sum([]byte(fmt.Sprintf("%d", rand.Intn(1000))))
		s = append(s, fmt.Sprintf("('%s')", fmt.Sprintf("%x", hex)))
	}
	sort.Strings(s)
	q := "INSERT IGNORE INTO t1 (name) VALUES " + strings.Join(s, ",")
	if _, err := db.Exec(q); err != nil {
		log.Fatal(err)
	}
}

func DropTableIfExists() error {
	bytes, err := os.ReadFile("sql/t1_drop.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	return err
}

func CreateTable() error {
	bytes, err := os.ReadFile("sql/t1.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	return err
}

func InsertSeed() error {
	bytes, err := os.ReadFile("sql/seed.sql")
	if err != nil {
		return err
	}
	_, err = db.Exec(string(bytes))
	return err
}
