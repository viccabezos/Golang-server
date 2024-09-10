package main

import (
	"fmt"
	"log"
	"net/http"
)

func formHandler(w http.ResponseWriter, r *http.Request) {
	if err := r.ParseForm(); err != nil {
		fmt.Fprintf(w, "Error parsing form: %v", err)
		return
	}
	fmt.Fprintf(w, "POST request successful:\n")
	name := r.FormValue("name")
	adress := r.FormValue("adress")
	fmt.Fprintf(w, "Name: %s\n", name)
	fmt.Fprintf(w, "Adress: %s\n", adress)
}

func helloHandler(w http.ResponseWriter, r *http.Request) {
	if r.URL.Path != "/Hello" {
		http.NotFound(w, r)
		// http.NotFound(w, "404 page not found", http.StatusNotFound)
		return
	}
	if r.Method != "GET" {
		http.Error(w, "Method not allowed", http.StatusMethodNotAllowed)
		return
	}
	fmt.Fprintf(w, "Hello World")
}

func main() {

	//tell to check the static directly and give the folder
	fileServer := http.FileServer(http.Dir("./static"))

	http.Handle("/", fileServer)
	http.HandleFunc("/form", formHandler)
	http.HandleFunc("/Hello", helloHandler)

	fmt.Printf("server started at port 8080\n")

	if err := http.ListenAndServe(":8080", nil); err != nil {
		log.Fatal(err)
	}

}
