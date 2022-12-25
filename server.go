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
		/*fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
		firstName := r.FormValue("fNameText")
		lastName := r.FormValue("lNameText")
		//test := r.FormValue()
		fmt.Fprintf(w, "Name = %s\n", firstName)
		fmt.Fprintf(w, "Address = %s\n", lastName)*/
	}

	//case "POST":
	//  if err := r.ParseForm(); err != nil {
	//    fmt.Fprintf(w, "ParseForm() err: %v", err)
	//    return
	//  }
	//  fmt.Fprintf(w, "Post from website! r.PostFrom = %v\n", r.PostForm)
	//  name := r.FormValue("name")
	//  address := r.FormValue("address")
	//  fmt.Fprintf(w, "Name = %s\n", name)
	//  fmt.Fprintf(w, "Address = %s\n", address)
	//default:
	//  fmt.Fprintf(w, "Sorry, only GET and POST methods are supported.")
	//}

}
