package main

import "fmt"

type Student struct {
	Name		string
	Class		string
}

func (s *Student) SetMyName(name string)  {
	s.Name = name
}

func (s Student) CallMyName()  {
	fmt.Printf("Hello, My Name is %s.", s.Name)
}

func main() {
	student := Student{
		Name: "Noobee",
		Class: "12E",
	}

	fmt.Printf("Nama awal student:%s\n", student.Name)

	student.SetMyName("Ubit")
	student.CallMyName()
}