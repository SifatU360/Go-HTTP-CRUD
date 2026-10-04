package main

import (
	"context"
	"encoding/json"
	"fmt"
	"go_http_crud/db"
	"net/http"
	"strconv"

	"github.com/jackc/pgx/v5"
)

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

	// Using Slice -> static DB
	// newUser.Id = len(user) + 1
	// user = append(user, newUser)

	query := `
		insert into users(username, age, email)
		values ($1, $2, $3)
		returning id
	`

	err = db.Db.QueryRow(context.Background(), query, newUser.Name, newUser.Age, newUser.Email).Scan(&newUser.Id)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not create user")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(newUser)
}

func getUserHandler(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	// user, _ := json.Marshal(user) //json.Mershal -> first it save into memory then write with w.Write
	// w.Write(user)

	query := `select id, username, age, email from users`

	rows, err := db.Db.Query(context.Background(), query)
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not get users")
		return
	}

	defer rows.Close()

	var users []User

	for rows.Next() {
		var user User

		err := rows.Scan(&user.Id, &user.Name, &user.Age, &user.Email)
		if err != nil {
			w.WriteHeader(http.StatusInternalServerError)
			fmt.Fprintln(w, "Could not scan user")
			return
		}
		users = append(users, user)
	}

	err = rows.Err()
	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not read user")
		return
	}

	encoder := json.NewEncoder(w) // it done by stream and write , memory efficient
	encoder.Encode(users)
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

	// for _, usr := range user{
	// 	if usr.Id == id {
	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(usr)
	// 		return
	// 	}
	// }

	var user User
	query := `SELECT id, username, age, email FROM users WHERE id = $1`
	err = db.Db.QueryRow(context.Background(), query, id).Scan(&user.Id, &user.Name, &user.Age, &user.Email)
	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User not found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not get User")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNotFound)
	fmt.Fprintln(w, "User Not Found ..!")
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
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid Request Body")
		return
	}

	// for idx, usr := range user{
	// 	if usr.Id == id {
	// 		updatedUser.Id = id
	// 		user[idx] = updatedUser

	// 		w.Header().Set("Content-Type", "application/json")
	// 		json.NewEncoder(w).Encode(updatedUser)
	// 		return
	// 	}
	// }

	// w.Header().Set("Content-Type", "application/json")
	// w.WriteHeader(http.StatusNotFound)
	// fmt.Fprintln(w,"User Not Found ..!")

	query := `
		update users
		set username = $1, age = $2, email = $3
		where id = $4
		returning id, username, age, email
	`

	err = db.Db.QueryRow(context.Background(), query, updatedUser.Name, updatedUser.Age, updatedUser.Email, id).Scan(&updatedUser.Id, &updatedUser.Name, &updatedUser.Age, &updatedUser.Email)

	if err == pgx.ErrNoRows {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User Not Found")
		return
	}

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could Not Update User")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	json.NewEncoder(w).Encode(updatedUser)
}

func deleteUserHandler(w http.ResponseWriter, r *http.Request) {
	idParam := r.PathValue("id")
	id, err := strconv.Atoi(idParam)

	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		fmt.Fprintln(w, "Invalid user id")
		return
	}

	// for idx, usr := range user{
	// 	if usr.Id == id {
	// 		// user = append(user[:idx], user[idx+1:]... )
	// 		user = slices.Delete(user, idx, idx+1)
	// 		w.WriteHeader(http.StatusNoContent)
	// 		return
	// 	}
	// }

	query := `
		delete from users where id = $1
	`

	cmdTag, err := db.Db.Exec(context.Background(), query, id)

	if err != nil {
		w.WriteHeader(http.StatusInternalServerError)
		fmt.Fprintln(w, "Could not Delete User")
		return
	}

	if cmdTag.RowsAffected() != 1 {
		w.WriteHeader(http.StatusNotFound)
		fmt.Fprintln(w, "User Not Found")
		return
	}

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusNoContent)
	fmt.Fprintln(w, "User deleted Successfully ..!")
}
