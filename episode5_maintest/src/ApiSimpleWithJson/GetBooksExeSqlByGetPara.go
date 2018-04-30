package main

import (
	sqlHelper "TestSqlDatabaseMssql"
	"encoding/json"
	"log"
	"math/rand"
	"net/http"
	"time"
)

type Book struct {
	Name   string
	IsRent bool
	Id     string
}

var books = []Book{}

var letterRunes = []rune("abcdefghijklmnopqrstuvwxyzABCDEFGHIJKLMNOPQRSTUVWXYZ")

func RendStringRunes(n int) string {
	b := make([]rune, n)
	for i := range b {
		b[i] = letterRunes[rand.Intn(len(letterRunes))]
	}
	return string(b)
}

func init() {
	rand.Seed(time.Now().UnixNano())
}

func main() {
	http.HandleFunc("/books", GetBooks)
	http.HandleFunc("/exesql", ExeSql)
	log.Fatal(http.ListenAndServe(":9012", nil))
}

func GetBooks(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	newBook := Book{
		Name:   "book A",
		IsRent: false,
		Id:     RendStringRunes(8),
	}
	books = append(books, newBook)
	var encoder = json.NewEncoder(w)
	encoder.Encode(books)
	w.Write([]byte("test"))
}

func ExeSql(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Access-Control-Allow-Origin", "*")
	var db, err = sqlHelper.GetDB()
	if err != nil {
		w.Write([]byte(err.Error()))
	} else {
		query, ok := r.URL.Query()["query"]
		if ok {
			var exeResult, err = sqlHelper.ExecSql(db, string(query[0]))
			if err != nil {
				w.Write([]byte(err.Error()))
			} else {
				// var encoder = json.NewEncoder(w)
				// encoder.Encode(exeResult)
				w.Write([]byte(exeResult))
			}
		} else {
			w.Write([]byte("query parameter not exist"))
		}
	}
}
