package main

import (
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
	myRouter.HandleFunc("/words", insertWord).Methods("POST")
	myRouter.HandleFunc("/words/{id}", updateWord).Methods("PUT")
	myRouter.HandleFunc("/words/{id}", deleteWord).Methods("DELETE")
	myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("../webapp/dist/webapp/")))
	//myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("../webapp/dist/translate-id/")))
	//myRouter.PathPrefix("/").Handler(http.FileServer(http.Dir("../webapp/dist/translate-ko/")))
	//175.125.246.138
	log.Fatal(http.ListenAndServe(":8000", myRouter))
}

func homePage(w http.ResponseWriter, r *http.Request) {
	t, err := template.ParseFiles("index.html")
	checkError(err)
	t.Execute(w, nil)
}

func returnWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	json.NewEncoder(w).Encode(db_getWord(key))
}

func returnWords(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	var words []Word = db_getWords()
	json.NewEncoder(w).Encode(words)
}

func insertWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var w1 Word
	json.Unmarshal(reqBody, &w1)
	w1.Id = db_getLastId() + 1
	db_insertWord(w1)
	fmt.Fprintf(w, "id:%d, name:%s, mean:%s", w1.Id, w1.Name, w1.Mean)
}

func updateWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	reqBody, _ := ioutil.ReadAll(r.Body)
	var w1 Word
	json.Unmarshal(reqBody, &w1)
	db_updateWord(w1)
	fmt.Fprintf(w, "id:%d, name:%s, mean:%s", w1.Id, w1.Name, w1.Mean)
}

func deleteWord(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	key := vars["id"]
	db_deleteWord(key)
	fmt.Fprintf(w, "id:%s", key)
}











//DB
func db_deleteWord(searchId string) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	result, err := db.Exec("DELETE FROM word_note WHERE word_id=$1", searchId)
	checkError(err)

	cnt, err := result.RowsAffected()
	checkError(err)

	fmt.Println("Delete Rows: ", cnt)
}

func db_updateWord(w1 Word) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	result, err := db.Exec("UPDATE word_note SET word_name=$1, word_mean=$2 WHERE word_id=$3", w1.Name, w1.Mean, w1.Id)
	checkError(err)

	cnt, err := result.RowsAffected()
	checkError(err)

	fmt.Println("Update Rows: ", cnt)
}

func db_insertWord(w1 Word) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	result, err := db.Exec("INSERT INTO word_note (word_id, word_name, word_mean) VALUES($1, $2, $3)", w1.Id, w1.Name, w1.Mean)
	checkError(err)

	cnt, err := result.RowsAffected()
	checkError(err)

	fmt.Println("Insert Rows: ", cnt)
}

func db_getWord(searchId string) Word {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	var id int
	var name string
	var mean string
	err = db.QueryRow("SELECT word_id, word_name, word_mean FROM word_note WHERE word_id=$1", searchId).Scan(&id, &name, &mean)
	checkError(err)

	var w1 Word = Word{Id: id, Name: name, Mean: mean}

	return w1
}

func db_getWords() []Word {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	rows, err := db.Query("SELECT word_id, word_name, word_mean FROM word_note ORDER BY word_id")
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

func db_getLastId() int {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	checkError(err)

	defer db.Close()

	var id int
	err = db.QueryRow("SELECT MAX(word_id) FROM word_note").Scan(&id)
	checkError(err)

	return id
}
