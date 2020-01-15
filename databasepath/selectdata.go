package databasepath

import (
	"encoding/json"
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	// dbp "github.com/unknownladge/newapi/databasepath"
)

var db = SqlHandler{
	Conn:    SqliteHandler.Conn,
	Command: SqliteHandler,
}
var Od = Order{
	Dborder: db,
}

func (db Order) ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Select all")

	// vars := mux.Vars(r)
	// key := vars["id"]
	//  remove to out function  ////////////////////////////////
	ans, err := db.Dborder.Selectall2()
	if err != nil {
		if (string(err.Error())) == "Marshal error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Marshal error"}
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

		if (string(err.Error())) == "Time Format error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Time Format error"}
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
		if (string(err.Error())) == "Scan row error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Scan row error"}
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
	}
	_, err = w.Write([]byte(ans))

}

func (db Order) ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Select one")
	vars := mux.Vars(r)
	key := vars["id"]

	ans, err := db.Dborder.Selectone(key)

	if err != nil || ans == "" {
		log.Println(err.Error())
		if (string(err.Error())) == "Article not found" {
			m := Errordetail{Errorcode: 400, Errordesc: "Article not found"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}

		}
		if (string(err.Error())) == "id not found" {
			m := Errordetail{Errorcode: 400, Errordesc: "id not found"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}

		}
		if (string(err.Error())) == "cannot convert isbn" {
			m := Errordetail{Errorcode: 400, Errordesc: "cannot convert isbn"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}

		}
		if (string(err.Error())) == "Marshal error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Marshal error"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}

		}
		if (string(err.Error())) == "Error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Error"}
			e, err := json.Marshal(m)
			if err != nil {
				panic(err)
			}

			w.WriteHeader(http.StatusBadRequest)
			_, err = w.Write([]byte(e))
			if err != nil {
				panic(err)
			}

		}

	} else {

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(ans))
		if err != nil {
			panic(err)
		}

	}

}
