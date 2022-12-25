package main

import (
	"fmt"
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

		fmt.Println("POST catched successfully!")
		firstName := r.FormValue("fNameText")
		lastName := r.FormValue("lNameText")
		fmt.Println("User's first name is ", firstName)
		fmt.Println("User's last name is ", lastName)
		fileServer.ServeHTTP(w, r)
	}
}
