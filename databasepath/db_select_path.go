package databasepath

import (
	"encoding/json"
	"errors"
	"log"

	"time"
)

func (db *SqlHandler) Dbget1(n string) (Article, error) {
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

func (db *SqlHandler) Selectall2() (string, error) {
	var set []Article
	var err error
	set, err = db.Command.Dbgetall()
	if err != nil {
		return "", errors.New("Marshal error")
	}
	s, err := json.Marshal(set)
	if err != nil {
		//log.Println("Marshal error")
		return "", errors.New("Marshal error")
	}
	log.Println(string(s))

	return string(s), nil
}
func (db *SqlHandler) Dbgetall() ([]Article, error) {
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
