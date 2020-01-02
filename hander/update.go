package hander

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func UpdateArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	log.Println("Update")
	vars := mux.Vars(r)
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {

		m := Errordetail{Errorcode: 400, Errordesc: "Json cannot unmarshal"}
		e, err := json.Marshal(m)
		if err != nil {
			log.Println("error")
		}
		w.Header().Set("Content-Type", "application/json")
		w.WriteHeader(http.StatusBadRequest)
		_, err = w.Write([]byte(e))
		if err != nil {
			panic(err)
		}
	}
	//log.Println("Create")

	id := vars["id"]
	title := article.Title
	desc := article.Desc
	content := article.Content
	isbn := article.ISBN
	log.Println(r)
	log.Println(vars)
	log.Println(id, " ", title, " ", desc, " ", content, " ", isbn)

	_, err = updatedata(id, title, desc, content, isbn)
	if err != nil {
		if (string(err.Error())) == "notfound" {

			m := Errordetail{Errorcode: 400, Errordesc: "Article not found"}
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
		if (string(err.Error())) == "Columerror" {

			m := Errordetail{Errorcode: 400, Errordesc: "Colum error or not found"}
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

			m := Errordetail{Errorcode: 400, Errordesc: "cant convert isbn format"}
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

			m := Errordetail{Errorcode: 400, Errordesc: "string lenght not match"}
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
		if (string(err.Error())) == "nothing update" {

			m := Errordetail{Errorcode: 400, Errordesc: "nothing update (row not found or not yet created)"}
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
	}
	w.WriteHeader(http.StatusOK)

}
