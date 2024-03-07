package main

import (
	"fmt"
	"net/http"
)

func handleHomePage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w,"This is the homepage")
}

func handleContactsPage(w http.ResponseWriter, r *http.Request){
	fmt.Fprintf(w, "This is the contacts page")
}

func main(){

	http.HandleFunc("/", handleHomePage)
	http.HandleFunc("/contacts",handleContactsPage)
	http.ListenAndServe(":3000",nil)

}
