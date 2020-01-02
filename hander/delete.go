package hander

import (
	"encoding/json"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	id := vars["id"]
	_, err := deletedata(string(id))
	if err != nil {
		if (string(err.Error())) == "errors" {
			m := Errordetail{Errorcode: 400, Errordesc: "Cannot Delete Somthing error"}
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
	log.Println("test")

	///////////

}
