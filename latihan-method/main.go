package main

import "fmt"

type Player struct {
	Name		string
	Score		int
}

func (p Player) AddScore(input int) Player {
	p.Score += input
	return p
}

func (p Player) DisplayInfo()  {
	fmt.Printf("Nama pemain: %s, Skor:%d", p.Name, p.Score)
}

func main() {
	player := Player{
		Name: "John",
		Score: 0,
	}

	player = player.AddScore(10)
	player = player.AddScore(5)
	player.DisplayInfo()
}