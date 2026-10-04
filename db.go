package main

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var db *pgx.Conn

//  `json:"id` -> when convert in json then send like id not Id

func connectDB() {
	var err error
	connStr := os.Getenv("DB_STRING")
	db, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully..!!")
}
