package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

type Article struct {
	Id      string `json:"id"`
	Title   string `json:"title"`
	Desc    string `json:"desc"`
	Content string `json:"content"`
}

var articles []Article //We're using this as a "memory database"

func homePage(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintf(w, "My Index")
}

func getArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r) //args sent, like id for updating or querying resources

	key := vars["id"] //specific arg we want

	for _, article := range articles {
		if article.Id == key {
			json.NewEncoder(w).Encode(article) //returns in JSON our article
		}
	}
}

func allArticles(w http.ResponseWriter, r *http.Request) {

	json.NewEncoder(w).Encode(articles)
}

func createArticle(w http.ResponseWriter, r *http.Request) {

	reqBody, _ := ioutil.ReadAll(r.Body) //request body in bytes
	var article Article
	json.Unmarshal(reqBody, &article) //assign this body to struct
	articles = append(articles, article)
	json.NewEncoder(w).Encode(article)
}

func updateArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	reqBody, _ := ioutil.ReadAll(r.Body)

	var updatedArticle Article //we will store our updated article here

	json.Unmarshal(reqBody, &updatedArticle)

	for index, article := range articles {
		if article.Id == id {
			article.Title = updatedArticle.Title
			article.Desc = updatedArticle.Desc
			article.Content = updatedArticle.Content
			articles[index] = article
			json.NewEncoder(w).Encode(article)

		}
	}
}

func deleteArticle(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id := vars["id"]

	for index, article := range articles {
		if article.Id == id {
			articles = append(articles[:index], articles[index+1:]...)
		}
	}
}

//routes definition with Mux
func handleRequest() {
	router := mux.NewRouter().StrictSlash(true)
	router.HandleFunc("/articles", allArticles).Methods("GET")
	router.HandleFunc("/article/{id}", getArticle).Methods("GET")
	router.HandleFunc("/article", createArticle).Methods("POST")
	router.HandleFunc("/article/{id}", deleteArticle).Methods("DELETE")
	router.HandleFunc("/article/{id}", updateArticle).Methods("PUT")
	router.HandleFunc("/", homePage)
	log.Fatal(http.ListenAndServe(":8081", router))
}

func main() {

	//We start with this data
	articles = []Article{
		{Id: "1", Title: "Other", Desc: "Other", Content: "Other"},
		{Id: "2", Title: "New", Desc: "New ", Content: "New"},
	}
	handleRequest() //Here we call our request handler
}
