package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	//"github.com/golang-migrate/migrate/v4"
	//_ "github.com/golang-migrate/migrate/v4/database/postgres"
	//_ "github.com/golang-migrate/migrate/v4/source/github"
	_ "github.com/jackc/pgx/v5/stdlib"
)

var fileServer = http.FileServer(http.Dir("./assets"))

func main() {

	http.HandleFunc("/", serveIndexPage)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		fmt.Println(err)
	}
}

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	//log.Printf("method %s, path: %s\n", r.Method, r.URL.Path)
	if r.Method == http.MethodGet {
		fileServer.ServeHTTP(w, r)
		return
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		//fmt.Println(r.Body)
		byteValue, err := ioutil.ReadAll(r.Body)
		if err != nil {
			fmt.Println(err)
		}
		w.Write(byteValue)

		var result map[string]interface{}
		err = json.Unmarshal([]byte(byteValue), &result)
		if err != nil {
			fmt.Println(err)
		}

		fmt.Println("User's first name is:", result["fNameText"])
		fmt.Println("User's last name is:", result["lNameText"])

		//os.Getenv("DATABASE_URL")
		dbURL := "postgres://admin:qwert@127.0.0.1:2023/usersInfoDB"
		db, err := sql.Open("pgx", dbURL)
		if err != nil {
			fmt.Println(err)
		} else {
			fmt.Println("Successfully connected to DB")
		}
		err = db.Ping()
		if err != nil {
			fmt.Printf("Ping to db failed, %s", err.Error())
		}
		//initDB(db)
		insertUserIntoDB(db, fmt.Sprint(result["fNameText"]), fmt.Sprint(result["lNameText"]))
		getUsersInfoFromDB(db)
		db.Close()
	}
}

func getUsersInfoFromDB(db *sql.DB) {
	var (
		uid    int
		ufname string
		ulname string
	)

	rows, err := db.Query("SELECT * from usersinfo")
	if err != nil {
		fmt.Println(err)
	}
	for rows.Next() {
		err := rows.Scan(&uid, &ufname, &ulname)
		if err != nil {
			fmt.Println(err)
		}
		fmt.Println(uid, ufname, ulname)
	}
	err = rows.Err()
	if err != nil {
		fmt.Println(err)
	}
	rows.Close()
}

func insertUserIntoDB(db *sql.DB, ufname string, ulname string) {
	insert, insertErr := db.Prepare("INSERT INTO usersinfo (ufname, ulname) VALUES ($1, $2)")
	if insertErr != nil {
		fmt.Println(insertErr)
	}
	_, err := insert.Exec(ufname, ulname)
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You have successfully added new user's info to database!")
	}
	insert.Close()
}

func initDB(db *sql.DB) {
	insert, insertErr := db.Prepare("CREATE TABLE usersinfo ALTER TABLE usersinfo ADD COLUMN Id SERIAL PRIMARY KEY, ufname VARCHAR(50) NOT NULL, ulname VARCHAR(50) NOT NULL")
	if insertErr != nil {
		fmt.Println(insertErr)
	}
	_, err := insert.Exec()
	if err != nil {
		fmt.Println(err)
	} else {
		fmt.Println("You have successfully added new user's info to database!")
	}
	insert.Close()
}
