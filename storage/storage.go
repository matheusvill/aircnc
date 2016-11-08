package storage

import (
    "database/sql"
    "fmt"
    _ "github.com/mattn/go-sqlite3"
)


type Storage interface {
	// inserir os dados e voltar a chave do storage
	InsertUser(email, password string) int
	// pegar todos os dados do user pelo id (que é chave)
	GetUser(id int) map[string]interface{}
	// atualizar todos os dados do user
	UpdateUser(email, password string) bool

	// inserir os dados e voltar a chave do storage
	InsertHome(name, address, country, city string, user int) int
	// pegar todos os dados do home pelo id
	GetHome(id string) map[string]interface{}
	// atualizar todos os dados do home
	UpdateHome(email, password string) bool
}


func (self *Storage) InsertUser(email, password string) int{

    db, err := sql.Open("sqlite3", "aircnc.db")
    checkErr(err)	

	tx, err := db.Begin()
	checkErr(err)

    // insert
    stmt, err := db.Prepare("INSERT INTO user(email, password) values(?,?)")
    checkErr(err)

    res, err := stmt.Exec(email, password)
    checkErr(err)

    id, err := res.LastInsertId()
    checkErr(err)

	tx.Commit()

	return id
}

func (self *Storage) GetUser(id int) map[string]interface{}{


	var email, password string

    db, err := sql.Open("sqlite3", "aircnc.db")
    checkErr(err)	

    // select
    stmt, err := db.Prepare("SELECT email, password FROM user WHERE id=?")
    checkErr(err)

	err = stmt.QueryRow(id).Scan(&email, &password)
	checkErr(err)

	user := map[string]interface{}{
		"email" : email,
		"password": password, 
	}


	return user
}



func checkErr(err error) {
    if err != nil {
        panic(err)
    }
}

func main(){

    db, err := sql.Open("sqlite3", "aircnc.db")
    checkErr(err)	


	sql := `CREATE TABLE user (
				id INTEGER PRIMARY KEY AUTOINCREMENT,
				username VARCHAR(64) NOT NULL,
				password VARCHAR(64) NOT NULL
			);`


	_, err = db.Exec(sql)
	checkErr(err)	

}