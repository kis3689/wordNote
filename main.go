package main

import (
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"

	"github.com/kis3689/WordNote/util"

	"github.com/gorilla/mux"

	"github.com/kis3689/WordNote/api"
)

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	myRouter.HandleFunc("/", homePage)
	myRouter.HandleFunc("/word", insertWord).Methods("POST")
	myRouter.HandleFunc("/word/{id}", returnWord)
	myRouter.HandleFunc("/words", returnWords)
	//175.125.246.138
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("views/index.html")
	util.CheckError(err)
	t.Execute(w, nil)
}

func insertWord(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var w1 api.Word
	json.Unmarshal(reqBody, &w1)
	fmt.Fprintf(w, "id:%d, name:%s, mean:%s", w1.Id, w1.Name, w1.Mean)

	// var w1 api.Word
	// json.Unmarshal(reqBody, &w1)
	// w1.Id = api.GetLastId()
	// api.AddWord(w1)
}

func returnWords(w http.ResponseWriter, r *http.Request) {
	var words []api.Word = api.GetWords()
	json.NewEncoder(w).Encode(words)
}

func returnWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	json.NewEncoder(w).Encode(api.GetWord(key))
}
