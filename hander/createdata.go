package hander

import (
	"encoding/json"
	"errors"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"strings"
	"time"
)

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		m := Errordetail{Errorcode: 400, Errordesc: "Cannot unmarshal this json"}
		e, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(e))
		if err != nil {
			panic(err)
		}
	}

	//json.NewEncoder(w).Encode(article)
	log.Println("Create")
	log.Printf("%+v\n", article)
	out, err := json.Marshal(article)
	if err != nil {
		//tell := Errordetail{Errorcode: http.StatusBadRequest, Errordesc: "Error json cant marshal"}
		m := Errordetail{Errorcode: 400, Errordesc: "Error json cant marshal"}
		e, err := json.Marshal(m)
		if err != nil {
			panic(err)
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(e))
		if err != nil {
			panic(err)
		}
	}
	log.Println(out)
	db := SqlHandler{Conn: SqliteHandler.Conn}
	_, err = db.Inserter(string(out))

	if err != nil {
		log.Println("Error")
		log.Println((string(err.Error())))
		if (string(err.Error())) == "Error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Someing Wrong in code"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
		if (string(err.Error())) == "duplicate key" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "duplicate key value violates unique constraint (id)"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
		if (string(err.Error())) == "out of range" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "data out of range"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
		if (string(err.Error())) == "value too long" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "value too long for type character varying"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
		if (string(err.Error())) == "wrong format" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "ISBN wrong format (XXX-X-XX-XXXXX-X)"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
			//http.Error(w, string(e), http.StatusBadRequest)
		}
		if (string(err.Error())) == "wrong format2" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "ISBN wrong format2 (XXX-X-XX-XXXXX-X)"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
		if (string(err.Error())) == "string lenght not match" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "string lenght not match (XXX-X-XX-XXXXXX-X)"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
		if (string(err.Error())) == "cant convert isbn format" {

			//w.WriteHeader(http.StatusBadRequest)

			m := Errordetail{Errorcode: 400, Errordesc: "cant convert isbn format (XXX-X-XX-XXXXXX-X)"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}
			w.Header().Set("Content-Type", "application/json")
			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}
		}
	} //end of err
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
