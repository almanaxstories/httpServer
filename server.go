package main

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
)

var fileServer = http.FileServer(http.Dir("./assets"))

func main() {

	http.HandleFunc("/", serveIndexPage)

	if err := http.ListenAndServe(":8090", nil); err != nil {
		log.Fatal(err)
	}
}

func serveIndexPage(w http.ResponseWriter, r *http.Request) {
	log.Printf("method %s, path: %s\n", r.Method, r.URL.Path)
	if r.Method == http.MethodGet {
		fileServer.ServeHTTP(w, r)
		return
	} else if r.Method == http.MethodPost {
		if err := r.ParseForm(); err != nil {
			fmt.Fprintf(w, "ParseForm() err: %v", err)
			return
		}
		fmt.Println(r.Body)
		byteValue, _ := ioutil.ReadAll(r.Body)

		var result map[string]interface{}
		json.Unmarshal([]byte(byteValue), &result)

		fmt.Println("User's first name is:", result["fNameText"])
		fmt.Println("User's last name is:", result["lNameText"])
		//fileServer.ServeHTTP(w, r)

	}
}
