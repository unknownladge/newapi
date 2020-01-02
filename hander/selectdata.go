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
	ans, err := selectall2()
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
	ans, err := selectone(string(key))
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

func selectone(n string) (string, error) {
	var id, title, desc1, content, isbn, times, update string
	var data Article

	row := SqliteHandler.Conn.QueryRow(`SELECT id, title, desc1, content, isbn, createtime, recentupdate FROM articleinfo where id = $1`, n)

	// isbn2 := isbn[:3]+"-"+isbn[3:4]+"-"+isbn[4:6]+"-"+isbn[7:12]+"-"+isbn[12:]
	log.Println(row)
	err := row.Scan(&id, &title, &desc1, &content, &isbn, &times, &update)
	if err != nil {
		return "", errors.New("not found")
	}
	if id == "" {
		//log.Println("Not Found!")
		return "", errors.New("id not found")
	}
	data.Id = id
	data.Title = title
	data.Desc = desc1
	data.Content = content

	data.ISBN, err = toisbn(isbn)
	if err != nil {
		return "", errors.New("cannot convert isbn")
	}
	//	data.ISBN = isbn[:3] + "-" + isbn[3:4] + "-" + isbn[4:6] + "-" + isbn[7:12] + "-" + isbn[12:]
	log.Println(times)
	data.Time1, err = time.Parse(time.RFC3339, times) //2015-09-15T14:00:13Z
	if err != nil {
		panic(err)
		//	log.Panicln("Convert time error")
	}
	data.Recentupdate, err = time.Parse(time.RFC3339, update) //2015-09-15T14:00:13Z
	if err != nil {
		panic(err)
		//	log.Panicln("Convert time error")
	}
	// times = strings.Replace(times, "T", " ", -1)
	// times = strings.Replace(times, "Z", "", -1)
	// settimes := strings.Split(times," ")
	// log.Println(times)
	// data.Time1 = settimes[0]+" "+settimes[1]
	log.Println(data.Time1)
	if err != nil {
		log.Println(err)
		log.Println("Error")
		return "", errors.New("Error")
	} else {
		log.Println(err)
		s, err := json.Marshal(data)
		if err != nil {
			log.Println("Marshal error")
			return "", errors.New("Marshal error")
		}
		log.Println(string(s))

		return string(s), nil
	}
}
func selectall2() (string, error) {
	var data Article
	var set []Article

	var id, title, desc1, content, isbn, date, update string
	rows, err := SqliteHandler.Conn.Query(`SELECT id, title, desc1, content, isbn, createtime, recentupdate FROM articleinfo ORDER BY id ASC`)
	if err != nil {
		log.Println("Failed to run query", err)
	}
	defer rows.Close()
	for rows.Next() {

		err = rows.Scan(&id, &title, &desc1, &content, &isbn, &date, &update)
		if err != nil {
			return "", errors.New("Scan row error")
		}
		data.Id = id
		data.Title = title
		data.Desc = desc1
		data.Content = content
		isbn := string(isbn)
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
		data.ISBN = isbn[:3] + "-" + isbn[3:4] + "-" + isbn[4:6] + "-" + isbn[7:12] + "-" + isbn[12:]
		times := date
		data.Time1, err = time.Parse(time.RFC3339, times) //2015-09-15T14:00:13Z
		if err != nil {
			return "", errors.New("Time Format error")
		}
		update := update
		data.Recentupdate, err = time.Parse(time.RFC3339, update) //2015-09-15T14:00:13Z
		if err != nil {
			return "", errors.New("Time Format error")
		}
		set = append(set, data)

	}
	s, err := json.Marshal(set)
	if err != nil {
		//log.Println("Marshal error")
		return "", errors.New("Marshal error")
	}
	log.Println(string(s))

	return string(s), nil
}
