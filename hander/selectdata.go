package hander

import (
	"encoding/json"
	"errors"
	"fmt"
	"log"
	"net/http"
	"time"

	"github.com/gorilla/mux"
)

func ReturnAllArticles(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Select all")
	// vars := mux.Vars(r)
	//     key := vars["id"]
	db := SqlHandler{
		Conn:    SqliteHandler.Conn,
		Command: SqliteHandler,
	}
	ans, err := db.selectall2()
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
	if err != nil {
		m := Errordetail{Errorcode: 400, Errordesc: "Writer func dead"}
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
}

func ReturnSingleArticle(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Println("Select one")
	vars := mux.Vars(r)
	key := vars["id"]
	db := SqlHandler{
		Conn:    SqliteHandler.Conn,
		Command: SqliteHandler,
	}
	ans, err := db.selectone(string(key))
	if err != nil || ans == "" {
		log.Println(err.Error())
		if (string(err.Error())) == "not found" {
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
		if (string(err.Error())) == "id not found" {
			m := Errordetail{Errorcode: 400, Errordesc: "id not found"}
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
		if (string(err.Error())) == "cannot convert isbn" {
			m := Errordetail{Errorcode: 400, Errordesc: "cannot convert isbn"}
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
		if (string(err.Error())) == "Error" {
			m := Errordetail{Errorcode: 400, Errordesc: "Error"}
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

	} else {

		w.WriteHeader(http.StatusOK)
		_, err = w.Write([]byte(ans))
		if err != nil {
			panic(err)
		}

	}

}

func (db *SqlHandler) selectone(n string) (string, error) {
	var err error
	var text string

	row, err := db.Command.dbget1(n)
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

func (db *SqlHandler) dbget1(n string) (Article, error) {
	var id, title, desc1, content, isbn, times, update string
	row := db.Conn.QueryRow(`SELECT id, title, desc1, content, isbn, createtime, recentupdate FROM articleinfo where id = $1`, n)
	var ans Article
	// isbn2 := isbn[:3]+"-"+isbn[3:4]+"-"+isbn[4:6]+"-"+isbn[7:12]+"-"+isbn[12:]
	log.Println(row)
	err := row.Scan(&id, &title, &desc1, &content, &isbn, &times, &update)
	if err != nil {
		//return nil
		return ans, errors.New("not found")
	}
	if id == "" {
		//return nil
		return ans, errors.New("id not found")
	}

	ans.Id = id
	ans.Title = title
	ans.Desc = desc1
	ans.Content = content
	ans.ISBN = isbn
	ans.Time1, _ = time.Parse(time.RFC3339, times)
	ans.Recentupdate, _ = time.Parse(time.RFC3339, update)
	return ans, nil
}

func (db *SqlHandler) selectall2() (string, error) {
	var set []Article
	set, _ = db.Command.dbgetall()
	s, err := json.Marshal(set)
	if err != nil {
		//log.Println("Marshal error")
		return "", errors.New("Marshal error")
	}
	log.Println(string(s))

	return string(s), nil
}
func (db *SqlHandler) dbgetall() ([]Article, error) {
	var err error
	var id, title, desc1, content, isbn, date, update string
	var data Article
	var set []Article
	rows, err := db.Conn.Query(`SELECT id, title, desc1, content, isbn, createtime, recentupdate FROM articleinfo ORDER BY id ASC`)
	if err != nil {
		log.Println("Failed to run query", err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&id, &title, &desc1, &content, &isbn, &date, &update)
		if err != nil {
			return set, errors.New("Scan row error")
		}
		data.Id = id
		data.Title = title
		data.Desc = desc1
		data.Content = content

		data.ISBN, _ = toisbn(isbn)
		log.Println(data.ISBN)
		times := date
		data.Time1, err = time.Parse(time.RFC3339, times) //2015-09-15T14:00:13Z
		if err != nil {
			return set, errors.New("Time Format error")
		}
		update := update
		data.Recentupdate, err = time.Parse(time.RFC3339, update) //2015-09-15T14:00:13Z
		if err != nil {
			return set, errors.New("Time Format error")
		}
		set = append(set, data)

	}
	return set, nil
}
