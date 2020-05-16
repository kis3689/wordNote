package api

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"

	"github.com/kis3689/WordNote/util"
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

func AddWord(w1 Word) {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	util.CheckError(err)

	defer db.Close()

	result, err := db.Exec("INSERT INTO word_note (word_id, word_name, word_mean) VALUES(" + string(w1.Id) + ", '" + w1.Name + "', '" + w1.Mean + "')")
	util.CheckError(err)

	cntAffected, err := result.RowsAffected()
	util.CheckError(err)

	fmt.Println("Affected Rows: ", cntAffected)
}

func GetWord(searchId string) Word {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	util.CheckError(err)

	defer db.Close()

	rows, err := db.Query("SELECT word_id, word_name, word_mean FROM word_note WHERE word_id=" + searchId)
	util.CheckError(err)

	defer rows.Close()

	var w1 Word

	var id int
	var name string
	var mean string
	for rows.Next() {
		err := rows.Scan(&id, &name, &mean)
		util.CheckError(err)
		w1 = Word{Id: id, Name: name, Mean: mean}
	}
	return w1
}

func GetWords() []Word {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	util.CheckError(err)

	defer db.Close()

	rows, err := db.Query("SELECT word_id, word_name, word_mean FROM word_note")
	util.CheckError(err)

	defer rows.Close()

	var words []Word

	var id int
	var name string
	var mean string
	for rows.Next() {
		err := rows.Scan(&id, &name, &mean)
		util.CheckError(err)
		w1 := Word{Id: id, Name: name, Mean: mean}
		words = append(words, w1)
	}
	return words
}

func GetLastId() int {
	dbInfo := fmt.Sprintf("user=%s password=%s dbname=%s sslmode=disable", DB_USER, DB_PASS, DB_NAME)
	db, err := sql.Open("postgres", dbInfo)

	util.CheckError(err)

	defer db.Close()

	rows, err := db.Query("SELECT MAX(word_id) FROM word_note")
	util.CheckError(err)

	defer rows.Close()

	var id int
	for rows.Next() {
		err := rows.Scan(&id)
		util.CheckError(err)
	}

	return id
}
