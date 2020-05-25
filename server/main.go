package main

import (
	"strconv"
	"database/sql"
	"encoding/json"
	"fmt"
	"html/template"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

type Word struct {
	Id   int
	Name string
	Mean string
}

type Result struct {
	Rst string
}

const (
	DB_USER = "postgres"
	DB_PASS = "dlstjq3689"
	DB_NAME = "test1"
)

func checkError(err error) {
	if err != nil {
		panic(err)
	}
}

func main() {
	myRouter := mux.NewRouter().StrictSlash(true)
	//myRouter.HandleFunc("/", homePage)
	//myRouter.HandleFunc("/word", insertWord).Methods("POST")
	//myRouter.HandleFunc("/word/{id}", returnWord)
	myRouter.HandleFunc("/words", returnWords).Methods("GET")
	myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("../webapp/dist/webapp/")))
	//175.125.246.138
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	checkError(err)
	t.Execute(w, nil)
}

func returnWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var words []Word = getWords()
	json.NewEncoder(w).Encode(words)
}

func returnWord(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	key := vars["id"]
	json.NewEncoder(w).Encode(getWord(key))
}

func insertWord(w http.ResponseWriter, r *http.Request) {
	reqBody, _ := ioutil.ReadAll(r.Body)
	var w1 Word
	json.Unmarshal(reqBody, &w1)
	w1.Id = getLastId() + 1
	var rst Result
	rst.Rst = addWord(w1)
	json.NewEncoder(w).Encode(rst)
	fmt.Fprintf(w, "id:%d, name:%s, mean:%s", w1.Id, w1.Name, w1.Mean)
}

func addWord(w1 Word) string {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	result, err := db.Exec("INSERT INTO word_note (word_id, word_name, word_mean) VALUES(" + strconv.Itoa(w1.Id) + ", '" + w1.Name + "', '" + w1.Mean + "')")
	checkError(err)

	cntAffected, err := result.RowsAffected()
	checkError(err)

	fmt.Println("Affected Rows: ", cntAffected)

	return "success"
}

func getWord(searchId string) Word {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	rows, err := db.Query("SELECT word_id, word_name, word_mean FROM word_note WHERE word_id=" + searchId)
	checkError(err)

	defer rows.Close()

	var w1 Word

	var id int
	var name string
	var mean string
	for rows.Next() {
		err := rows.Scan(&id, &name, &mean)
		checkError(err)
		w1 = Word{Id: id, Name: name, Mean: mean}
	}
	return w1
}

func getWords() []Word {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	rows, err := db.Query("SELECT word_id, word_name, word_mean FROM word_note")
	checkError(err)

	defer rows.Close()

	var words []Word

	var id int
	var name string
	var mean string
	for rows.Next() {
		err := rows.Scan(&id, &name, &mean)
		checkError(err)
		w1 := Word{Id: id, Name: name, Mean: mean}
		words = append(words, w1)
	}
	return words
}

func getLastId() int {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	rows, err := db.Query("SELECT MAX(word_id) FROM word_note")
	checkError(err)

	defer rows.Close()

	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		checkError(err)
	}

	return id
}
