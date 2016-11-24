package main

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func main() {
	fmt.Println("Creating database...")

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

	sql = `CREATE TABLE aircnc.apartment (
		id SERIAL PRIMARY KEY,
		titulo VARCHAR(64) NOT NULL,
		valor VARCHAR(64) NOT NULL, 
		descricao VARCHAR(64) NOT NULL, 
		cidade VARCHAR(64) NOT NULL, 
		uf VARCHAR(64) NOT NULL
	);`
	_, _ = db.Exec(sql)

	sql = `INSERT INTO aircnc.apartment (titulo,valor,descricao,cidade,uf) VALUES ('Neoway','100000','Empresa','Florianópolis','SC');`
	_, _ = db.Exec(sql)

	sql = `CREATE TABLE aircnc.booking (
		id SERIAL PRIMARY KEY,
		user_id INT,
		apartment_id INT
	);`
	_, _ = db.Exec(sql)

	sql = `INSERT INTO aircnc.booking (user_id,apartment_id) VALUES (
		(select id from aircnc.user where email = 'matheusvill@gmail.com'),
		(select id from aircnc.apartment where titulo = 'Neoway')
	);`
	_, _ = db.Exec(sql)
}
