package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/gorilla/mux"

	//thenew "github.com/beaw/newapi/hander" //server
	thenew "github.com/unknownladge/newapi/hander" //server
)

type Article struct {
	Id      string `json:"Id"`
	Title   string `json:"Title"`
	Desc    string `json:"Desc"`
	Content string `json:"Content"`
	Isbn    string `json:Isbn`
}

var Articles []Article

func main() {

	Articles = []Article{
		Article{Id: "1", Title: "Hello 1", Desc: "Article Description", Content: "Article Content"},
		Article{Id: "2", Title: "Hello 2", Desc: "Article Description", Content: "Article Content"},
	}
	thenew.ConnDB()
	handleRequests()
	defer thenew.CloseDB()
}

func handleRequests() {
	// creates a new instance of a mux router
	myRouter := mux.NewRouter().StrictSlash(true)
	// replace http.HandleFunc with myRouter.HandleFunc

	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/article", thenew.Od.ReturnAllArticles).Methods("GET")
	myRouter.HandleFunc("/article/{id}", thenew.Od.ReturnSingleArticle).Methods("GET")
	myRouter.HandleFunc("/article", thenew.CreateNewArticle).Methods("POST")
	myRouter.HandleFunc("/article/{id}", thenew.DeleteArticle).Methods("DELETE")
	myRouter.HandleFunc("/article/{id}", thenew.UpdateArticle).Methods("PUT")
	// finally, instead of passing in nil, we want
	// to pass in our newly created router as the second
	// argument
	log.Fatal(http.ListenAndServe(":10000", myRouter))
}
func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "Welcome to the HomePage!")
	//fmt.Println("Endpoint Hit: homePage")
}
