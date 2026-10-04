package db

import (
	"context"
	"fmt"
	"os"

	"github.com/jackc/pgx/v5"
)

var Db *pgx.Conn

//  `json:"id` -> when convert in json then send like id not Id

func ConnectDB() {
	var err error
	connStr := os.Getenv("DB_STRING")
	Db, err = pgx.Connect(context.Background(), connStr)
	if err != nil {
		panic(err)
	}

	fmt.Println("Database connected successfully..!!")
}
