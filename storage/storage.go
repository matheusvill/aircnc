package storage

import (
	"database/sql"
	"fmt"
	"log"

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

	stmt, err := db.Prepare("INSERT INTO aircnc.user (email, password) values ($1, $2)")
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

func InsertApartment(titulo, valor, descricao, cidade, uf string) int64 {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = db.Close() }()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO aircnc.apartment (titulo,valor,descricao,cidade,uf) values ($1, $2, $3, $4, $5)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(titulo, valor, descricao, cidade, uf)
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

func InsertBooking(user, apartment int) int64 {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}
	defer func() { _ = db.Close() }()

	tx, err := db.Begin()
	if err != nil {
		fmt.Println(err)
	}

	stmt, err := db.Prepare("INSERT INTO aircnc.booking (user_id,apartment_id) values ($1, $2)")
	if err != nil {
		fmt.Println(err)
	}

	res, err := stmt.Exec(user, apartment)
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

	var email, password string
	err = db.QueryRow("SELECT email, password FROM aircnc.user WHERE id = $1", id).Scan(&email, &password)
	if err != nil {
		log.Fatal(err)
	}

	user := map[string]interface{}{
		"email":    email,
		"password": password,
	}

	return user
}

func GetApartment(id int) map[string]interface{} {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	var titulo, valor, descricao, cidade, uf string
	err = db.QueryRow("SELECT titulo, valor, descricao, cidade, uf FROM aircnc.apartment WHERE id = $1", id).Scan(&titulo, &valor, &descricao, &cidade, &uf)
	if err != nil {
		log.Fatal(err)
	}

	ap := map[string]interface{}{
		"titulo":    titulo,
		"valor":     valor,
		"descricao": descricao,
		"cidade":    cidade,
		"uf":        uf,
	}

	return ap
}

func GetBooking(id int) map[string]interface{} {
	db, err := sql.Open("postgres", "postgres://root@localhost:26257?sslmode=disable")
	if err != nil {
		fmt.Println(err)
	}

	var user, apartment int
	err = db.QueryRow("SELECT user_id,apartment_id FROM aircnc.booking WHERE id = $1", id).Scan(&user, &apartment)
	if err != nil {
		log.Fatal(err)
	}

	book := map[string]interface{}{
		"user":      user,
		"apartment": apartment,
	}

	return book
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
