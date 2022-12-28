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
			log.Fatal(err)
		}

		fmt.Println("User's first name is:", result["fNameText"])
		fmt.Println("User's last name is:", result["lNameText"])
		//fileServer.ServeHTTP(w, r)

	}
}
