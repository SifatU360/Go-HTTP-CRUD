package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)

	fmt.Println("Server is running at port 5000...")

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server Error .. ", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Welcome to the Server..!")
}
func healthHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "Server is healthy UP and Running ...!")
}