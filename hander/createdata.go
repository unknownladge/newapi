package hander

import (
	"encoding/json"
	"io/ioutil"
	"log"
	"net/http"

	dbp "github.com/unknownladge/newapi/databasepath"
)

func CreateNewArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var article dbp.Article
	err := json.Unmarshal(reqBody, &article)
	if err != nil {
		m := dbp.Errordetail{Errorcode: 400, Errordesc: "Cannot unmarshal this json"}
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
		m := dbp.Errordetail{Errorcode: 400, Errordesc: "Error json cant marshal"}
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
	db := dbp.SqlHandler{Conn: dbp.SqliteHandler.Conn}
	_, err = db.Inserter(string(out))

	if err != nil {
		log.Println("Error")
		log.Println((string(err.Error())))
		if (string(err.Error())) == "Error" {
			m := dbp.Errordetail{Errorcode: 400, Errordesc: "Someing Wrong in code"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "duplicate key value violates unique constraint (id)"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "data out of range"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "value too long for type character varying"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "ISBN wrong format (XXX-X-XX-XXXXX-X)"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "ISBN wrong format2 (XXX-X-XX-XXXXX-X)"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "string lenght not match (XXX-X-XX-XXXXXX-X)"}
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

			m := dbp.Errordetail{Errorcode: 400, Errordesc: "cant convert isbn format (XXX-X-XX-XXXXXX-X)"}
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
