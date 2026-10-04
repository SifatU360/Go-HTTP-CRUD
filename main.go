package main

import (
	"encoding/json"
	"fmt"
	"net/http"
	"slices"
	"strconv"
)

//  `json:"id` -> when convert in json then send like id not Id
type User struct {
	Id    int    `json:"id"`
	Name  string `json:"name"`
	Age   int    `json:"age"`
	Email string `json:"email"`
}

var user = []User{
	{
		Id:    1,
		Name:  "Sifat",
		Age:   20,
		Email: "sifat@gmail.com",
	},
	{
		Id:    2,
		Name:  "Ariyan",
		Age:   21,
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
	mux.HandleFunc("GET /users/{id}", getSingleUserHandler)
	mux.HandleFunc("PUT /users/{id}", updateUserHandler)
	mux.HandleFunc("DELETE /users/{id}", deleteUserHandler)

	fmt.Println("Server is running at port 5000...")

	err := http.ListenAndServe(":5000", mux)

	if err != nil {
		fmt.Println("Server Error .. ", err)
	}

}

func rootHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	fmt.Fprintln(w, "Welcome to the Server..!")
}
func healthHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "Server is healthy UP and Running ...!")
}
func seeUserHandler(w http.ResponseWriter, r *http.Request) {
	fmt.Fprintln(w, "See user is Running in POST method ...!")
}

func createUserHandler(w http.ResponseWriter, r *http.Request) {
	// fmt.Println("Method:", r.Method, "Path:", r.URL.Path)

	if r.Method != "POST" {
		w.WriteHeader(http.StatusMethodNotAllowed)
		fmt.Fprintln(w, "Method is not allowed")
		return
	}

	var newUser User

	err := json.NewDecoder(r.Body).Decode(&newUser)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid Request Body")
		return
	}

	newUser.Id = len(user) + 1
	user = append(user, newUser)

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// user, _ := json.Marshal(user) //json.Mershal -> first it save into memory then write with w.Write
	// w.Write(user)

	encoder := json.NewEncoder(w) // it done by stream and write , memory efficient
	encoder.Encode(user)
}
func getSingleUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	// fmt.Printf("The value of id is %v and the type of value is %T", idParam, idParam)

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	for _, usr := range user{
		if usr.Id == id {
			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(usr)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w,"User Not Found ..!")
}
func updateUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	// fmt.Printf("The value of id is %v and the type of value is %T", idParam, idParam)

	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	var updatedUser User
	err = json.NewDecoder(r.Body).Decode(&updatedUser)
	if err != nil{
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid Request Body")
		return
	}

	for idx, usr := range user{
		if usr.Id == id {
			updatedUser.Id = id
			user[idx] = updatedUser

			w.Header().Set("Content-Type", "application/json")
			json.NewEncoder(w).Encode(updatedUser)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w,"User Not Found ..!")
}
func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	for idx, usr := range user{
		if usr.Id == id {
			user = append(user[:idx], user[idx+1:]... )
			// user = slices.Delete(user, idx, idx+1)
			w.WriteHeader(http.StatusNoContent)
			return
		}
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w,"User Not Found ..!")
}
