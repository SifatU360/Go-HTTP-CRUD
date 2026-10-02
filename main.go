package main

import (
	"encoding/json"
	"fmt"
	"net/http"
)

type User struct{
	Id int
	Name string
	Age int
	Email string
}

var user = []User{
	{
		Id: 1,
		Name: "Sifat",
		Age: 20,
		Email: "sifat@gmail.com",
	},
	{
		Id: 2,
		Name: "Ariyan",
		Age: 21,
		Email: "ariyan@gmail.com",
	},
}

func main() {
	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/createUser", createUserHandler)
	mux.HandleFunc("POST /seeUser", seeUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)

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
func seeUserHandler(w http.ResponseWriter, r *http.Request){
	fmt.Fprintln(w, "See user is Running in POST method ...!")
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

func getUserHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	user, _ := json.Marshal(user)
	w.Write(user)
}