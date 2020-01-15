package databasepath

import (
	"database/sql"
	"encoding/json"
	"errors"
	"fmt"
	"strings"
	"time"

	//"github.com/jmoiron/jsonq"
	"log"

	_ "github.com/lib/pq" // here
)

// const (
// 	host     = "172.31.25.45"
// 	port     = 5555
// 	user     = "postgres"
// 	password = "password"
// 	dbname   = "pasitbeaw"
// )

const (
	host     = "127.0.0.1"
	port     = 5432
	user     = "postgres"
	password = "1"
	dbname   = "task"
)

type Dbaseapi interface {
	Dbget1(string) (Article, error)
	Inserter(string) (string, error)
	Updatedata(string, string, string, string, string) (string, error)
	Deletedata(string) (string, error)
	Dbgetall() ([]Article, error)
	Selectone(string) (string, error)
	Selectall2() (string, error)
}
type Order struct {
	Dborder SqlHandler
}
type SqlHandler struct {
	Conn    *sql.DB
	Command Dbaseapi
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
func (db *SqlHandler) Inserter(n string) (string, error) {
	log.Println("insert")
	log.Println("'" + n + "'")

	var jq *Article = &Article{}

	err := json.Unmarshal([]byte(n), jq)
	if err != nil {
		log.Println(err.Error())
	}
	log.Print(jq)
	acc1 := jq.Id
	acc2 := jq.Title
	acc3 := jq.Desc
	acc4 := jq.Content
	acc5 := jq.ISBN
	acc5, err = setisbnformat(acc5)
	if err != nil {
		if string(err.Error()) == "string lenght not match" {
			return "", errors.New("string lenght not match")
		}
		if string(err.Error()) == "cant convert isbn format" {
			return "", errors.New("cant convert isbn format")
		}
	}
	t := time.Now().Format(time.RFC3339)
	log.Println(acc1 + acc2)
	log.Print(acc3)
	log.Print(acc4)
	log.Print(acc5)

	response := fmt.Sprintf(`INSERT INTO articleinfo(id,title,desc1,content,isbn,createtime,recentupdate) VALUES ('%s', '%s', '%s','%s','%s','%s','%s')`, acc1, acc2, acc3, acc4, acc5, t, t)
	log.Println("///")
	log.Println(response)
	log.Println("///")
	_, err = SqliteHandler.Conn.Exec(response)
	if err != nil {
		log.Println(err)
		if strings.Contains(string(err.Error()), "duplicate") {
			return "", errors.New("duplicate key")
		}
		if strings.Contains(string(err.Error()), "out of range") {
			//
			return "", errors.New("out of range")
		}
		if strings.Contains(string(err.Error()), "value too long") {
			//
			return "", errors.New("value too long")
		}

	}
	return "OK", nil
}
func (db *SqlHandler) Deletedata(id string) (string, error) {
	var err error

	sqlStatement := `DELETE FROM articleinfo WHERE id = $1;`
	_, err = db.Conn.Exec(sqlStatement, id)
	if err != nil {
		return "Error", errors.New("errors")
	}
	return "OK", nil
}
func (db SqlHandler) Selectone(n string) (string, error) {
	var err error
	var text string

	row, err := db.Command.Dbget1(n)
	if err != nil {
		return "", errors.New("not found")
	}
	//var id, title, desc1, content, isbn, times, update string
	text, err = createstring_selectone(row)
	return text, nil
}
func createstring_selectone(row Article) (string, error) {
	var err error
	var data Article
	data.Id = row.Id
	data.Title = row.Title
	data.Desc = row.Desc
	data.Content = row.Content

	data.ISBN, err = toisbn(row.ISBN)
	if err != nil {
		return "", errors.New("cannot convert isbn")
	}
	//	data.ISBN = isbn[:3] + "-" + isbn[3:4] + "-" + isbn[4:6] + "-" + isbn[7:12] + "-" + isbn[12:]
	//log.Println(times)
	data.Time1 = row.Time1 //2015-09-15T14:00:13Z

	data.Recentupdate = row.Recentupdate //2015-09-15T14:00:13Z

	log.Println(data.Time1)

	log.Println(err)
	s, _ := json.Marshal(data)

	log.Println(string(s))

	return string(s), nil

}
func (db *SqlHandler) Updatedata(id string, title string, desc string, content string, isbn string) (string, error) {
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
