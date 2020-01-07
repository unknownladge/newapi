package hander

import (
	"encoding/json"
	"errors"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func DeleteArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// once again, we will need to parse the path parameters
	vars := mux.Vars(r)

	id := vars["id"]
	db := SqlHandler{Conn: SqliteHandler.Conn}
	_, err := db.deletedata(string(id))
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

func (db *SqlHandler) deletedata(id string) (string, error) {
	var err error

	sqlStatement := `DELETE FROM articleinfo WHERE id = $1;`
	_, err = db.Conn.Exec(sqlStatement, id)
	if err != nil {
		return "Error", errors.New("errors")
	}
	return "OK", nil
}
