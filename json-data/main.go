package main

import (
	"encoding/json"
	"fmt"
)

type User struct {
	Name		string
	Class		string
}


func main() {
	users := []User{
		{Name: "Haha", Class: "A"},
		{Name: "Hehe", Class: "B"},
		{Name: "Hihi", Class: "C"},
	}

	fmt.Println("Bentuk dasar", users)

	// Encode struct object ke JSON
	userByteJson, err := json.Marshal(users)
	if err != nil {
		fmt.Println(err.Error())
		return
	}

	userJson := string(userByteJson)
	fmt.Println("Bentuk json:", userJson)


	/*DECODE JSON*/
	var users2 []User
	fmt.Println(users2)

	dataFrontend := `{"Name": "Hoho", "Class": "D"}`

	data := []byte(dataFrontend)
	var user User

	err = json.Unmarshal(data, &user)

	if err != nil {
		fmt.Println(err.Error())
		return
	}

	users2 = append(users2, user)

	fmt.Println("Setelah tambah data json:", users2)
	fmt.Println(len(users2))
}