package main

import (
	"fmt"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Welcome to ThreatX")
}

func main() {
	mux := http.NewServeMux()
	mux.HandleFunc("/", home)
	fmt.Println("Server running at port http://localhost:8000")
	err := http.ListenAndServe(":8000", mux)
	if err != nil {
		fmt.Println("error starting the server: ", err)
	}

}
