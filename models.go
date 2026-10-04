package main

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
