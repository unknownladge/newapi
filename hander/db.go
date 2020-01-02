package hander

import (
	"database/sql"
	"errors"
	"fmt"
	"strings"
	"time"

	//"github.com/jmoiron/jsonq"
	"log"

	_ "github.com/lib/pq" // here
)

const (
	host     = "172.31.25.45"
	port     = 5555
	user     = "postgres"
	password = "password"
	dbname   = "pasitbeaw"
)

// const (
// 	host     = "localhost"
// 	port     = 5555
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "pasitbeaw"
// )

type SqlHandler struct {
	Conn *sql.DB
}

var SqliteHandler = new(SqlHandler)

func ConnDB() {
	psqlInfo := fmt.Sprintf("host=%s port=%d user=%s "+"password=%s dbname=%s sslmode=disable", host, port, user, password, dbname)
	db, err := sql.Open("postgres", psqlInfo)
	if err != nil {
		panic(err)
	}

	err = db.Ping()
	if err != nil {
		panic(err)
	}

	log.Println("Successfully connected! DB")
	SqliteHandler.Conn = db
}
func CloseDB() {
	SqliteHandler.Conn.Close()
	log.Println("DB Close")
}

func deletedata(id string) (string, error) {
	sqlStatement := `DELETE FROM articleinfo WHERE id = $1;`
	_, err := SqliteHandler.Conn.Exec(sqlStatement, id)
	if err != nil {
		return "Error", errors.New("errors")
	}
	return "OK", nil
}
func updatedata(id string, title string, desc string, content string, isbn string) (string, error) {
	timeupdate := time.Now().Format(time.RFC3339)
	isbn, err := setisbnformat(isbn)
	if err != nil {
		return "", err
	}
	sqlStatement := `UPDATE articleinfo SET title = $2, desc1 = $3, content = $4,isbn = $5,recentupdate = $6 WHERE id = $1;`
	res, err := SqliteHandler.Conn.Exec(sqlStatement, id, title, desc, content, isbn, timeupdate)
	if err != nil {

		log.Println(err)
		if strings.Contains(string(err.Error()), "out of range") {
			//
			return "", errors.New("out of range")
		}
		if strings.Contains(string(err.Error()), "value too long") {
			//
			return "", errors.New("value too long")
		}
		return "", errors.New("notfound")
	}
	count, err := res.RowsAffected()
	if err != nil {
		return "", errors.New("Columerror")
	}
	log.Println(count)
	if count == 0 {
		return "", errors.New("nothing update")
	}

	return "OK", nil
}
func toisbn(isbn string) (string, error) {
	//XXX-X-XX-XXXXXX-X
	if len(isbn) < 13 {
		log.Println("Error text")
		missing := 13 - len(isbn)
		log.Println(len(isbn))
		log.Println(missing)
		for i := 0; i < missing; i++ {
			isbn = "0" + isbn
			log.Println(isbn)
		}

	}
	if len(isbn) > 13 {
		log.Println("text out of range ")
		return "", errors.New("text out of range")
	}

	return isbn[:3] + "-" + isbn[3:4] + "-" + isbn[4:6] + "-" + isbn[6:12] + "-" + isbn[12:], nil
}
func setisbnformat(text string) (string, error) {
	if len(text) < 16 {
		return "", errors.New("string lenght not match")
	}
	if strings.Contains(text, "-") && len(text) == 17 {
		text = strings.Replace(text, "-", "", -1)

	} else {
		return "", errors.New("cant convert isbn format")
	}
	return text, nil
}
