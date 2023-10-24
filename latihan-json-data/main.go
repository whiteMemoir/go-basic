package main

import (
	"encoding/json"
	"fmt"
	"os"
)

type User struct {
	Name  string
	Class string
}

func saveToFile(users []User, filename string) error {
	file, err := os.Create(filename)
	if err != nil {
		return err
	}

	data, err := json.Marshal(users)
	if err != nil {
		return err
	}

	_, err = file.Write(data)
	if err != nil {
		return err
	}

	fmt.Println("Data pengguna berhasil disimpan dalam file JSON.")
	return err
}


func loadFromFile(filename string) ([]User) {
	file, err := os.Open(filename)
	if err != nil {
		return nil
	}

	var users []User
	fileInfo, err := file.Stat()
	if err != nil {
		return nil
	}

	data := make([]byte, fileInfo.Size())
	_, err = file.Read(data)
	if err != nil {
		return nil
	}

	err = json.Unmarshal(data, &users)
	if err != nil {
		return nil
	}

	return users
}


func main() {
	users := []User{
		{Name: "NooB", Class: "A"},
		{Name: "Java", Class: "B"},
		{Name: "Docker", Class: "C"},
	}

	fileName := "users.json"

	err := saveToFile(users, fileName)
	if err != nil {
		fmt.Println("Error saving to file:", err.Error())
		return
	}

	fmt.Println("Data pengguna yang dibaca dari file JSON:")

	loadedUsers := loadFromFile(fileName)
	if err != nil {
		fmt.Println("Error loading from file:", err.Error())
		return
	}

	for _, user := range loadedUsers {
		fmt.Printf("Nama: %s, Kelas: %s\n", user.Name, user.Class)
	}
}
