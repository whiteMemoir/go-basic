package main

import "fmt"

type Calculate interface {
	Divide()		float32
	Multiply()	float32
	Sum()				float32
	Substract()		float32
}

type Number struct {
	Num				float32
	Divider		float32
}

func NewNum(num float32, divider float32) Calculate {
	return &Number{
		Num: num,
		Divider: divider,
	}
}

func (n Number) Divide() float32  {
	return n.Num / n.Divider
}
func (n Number) Multiply() float32  {
	return n.Num * n.Num
}
func (n Number) Sum() float32  {
	return n.Num + n.Num
}
func (n Number) Substract() float32  {
	return n.Num - n.Divider
}

func main() {
	num := NewNum(4.0,2.0)
	fmt.Println(num.Substract())
}