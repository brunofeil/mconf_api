package main

import (
	"fmt"
	"io/ioutil"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/urfave/negroni"
)

func GetBook(w http.ResponseWriter, r *http.Request) {
	params := mux.Vars(r)
	url := openLibraryAPI + params["book"]
	resp, err := http.Get(url)

	if err != nil {
		fmt.Println("Error http requst", err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		fmt.Println("Error reading resonse", err)
	}

	w.Write(body)
}

func applicationJSON() negroni.Handler {
	return negroni.HandlerFunc(func(w http.ResponseWriter, r *http.Request, next http.HandlerFunc) {

		w.Header().Set("Content-Type", "application/json")
		next(w, r)
	})
}

var (
	openLibraryAPI = "http://openlibrary.org/search.json?q="
)

func main() {
	n := negroni.Classic()
	n.Use(applicationJSON())

	r := mux.NewRouter().StrictSlash(true)
	n.UseHandler(r)

	r.HandleFunc("/{book}", GetBook).Methods("GET")

	fmt.Println("main listen at :8080")
	err := http.ListenAndServe(":8080", n)
	if err != nil {
		fmt.Println(err)
	}
}
