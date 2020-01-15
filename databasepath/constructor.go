package databasepath

import (
	"time"
)

// normal article
type Article struct {
	Id           string    `json:"Id"`
	Title        string    `json:"Title"`
	Desc         string    `json:"Desc"`
	Content      string    `json:"Content"`
	ISBN         string    `json:"ISBN"`
	Time1        time.Time `json:"Time"`
	Recentupdate time.Time `json:"Recentupdate"`
}

type ArticleDB struct {
	Id           string    `json:"Id"`
	Title        string    `json:"Title"`
	Desc         string    `json:"Desc"`
	Content      string    `json:"Content"`
	ISBN         int       `json:"ISBN"`
	Time1        time.Time `json:"Time"`
	Recentupdate time.Time `json:"Recentupdate"`
}

type Errordetail struct {
	Errorcode int
	Errordesc string
}
