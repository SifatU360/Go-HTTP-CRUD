package main

import (
	"context"
	"fmt"
	"go_http_crud/db"
	"net/http"

	// "slices"

	"github.com/joho/godotenv"
)

func main() {
	var err error

	err = godotenv.Load()
	if err != nil {
		panic("Env not found")
	}

	db.ConnectDB()
	defer db.Db.Close(context.Background())

	mux := http.NewServeMux()

	mux.HandleFunc("/", rootHandler)
	mux.HandleFunc("/health", healthHandler)
	mux.HandleFunc("/createUser", createUserHandler)
	mux.HandleFunc("POST /seeUser", seeUserHandler)
	mux.HandleFunc("GET /users", getUserHandler)
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running at port 5000...")

	err = http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server Error .. ", err)
	}

}
