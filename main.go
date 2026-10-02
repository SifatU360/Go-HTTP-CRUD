package main

import (
	"fmt"
	"net/http"
)

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/createUser", createUserHandler)

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
func createUserHandler(w http.ResponseWriter, r *http.Request){
	// fmt.Println("Method:", r.Method, "Path:", r.URL.Path)

	if r.Method != "POST"{
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method is not allowed")
		return
	}

	fmt.Fprintln(w, "User Created..!")
}