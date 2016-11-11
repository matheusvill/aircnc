package storage

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func InsertUser(email, password string) int64 {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = db.Close() }()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO user(email, password) values(?,?)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(email, password)
	if err != nil {
		fmt.Println(err)
	}

	id, err := res.LastInsertId()
	if err != nil {
		fmt.Println(err)
	}

	tx.Commit()

	return id
}

func GetUser(id int) map[string]interface{} {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = db.Close() }()

	var email, password string
	stmt, err := db.Prepare("SELECT email, password FROM aircnc.user WHERE id=?")
	if err != nil {
		fmt.Println(err)
	}

	err = stmt.QueryRow(id).Scan(&email, &password)
	if err != nil {
		fmt.Println(err)
	}

	user := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	return user
}

func CreateDb() {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = db.Close() }()

	sql := `DROP DATABASE IF EXISTS aircnc;`
	_, _ = db.Exec(sql)

	sql = `CREATE DATABASE aircnc;`
	_, _ = db.Exec(sql)

	sql = `CREATE TABLE aircnc.user (
		id SERIAL PRIMARY KEY,
		email VARCHAR(64) NOT NULL,
		password VARCHAR(64) NOT NULL
	);`
	_, _ = db.Exec(sql)

	sql = `INSERT INTO aircnc.user (email, password) VALUES ('matheusvill@gmail.com','1234');`
	_, _ = db.Exec(sql)
}
