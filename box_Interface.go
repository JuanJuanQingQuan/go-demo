package main

import "fmt"

type Box interface {
	Color() string
	Color2() string
}

type BigBox struct {
	ColorStr string
	Volume   float64
}

// 测试
func (b BigBox) SetColor(c string) BigBox {
	b.ColorStr = c
	return b
}

func (b BigBox) Color2() string {
	return b.ColorStr
}

func (b BigBox) Color() string {
	return b.ColorStr
}

func (b *BigBox) SetColor2(c string) {
	b.ColorStr = c
}

func BoxMain() {
	box := BigBox{}
	boxCopy := box.SetColor("red")
	// BigBox can be Box
	var box2 Box = box
	// *BigBox can be Box, too !!!
	var box3 Box = &box
	box.SetColor2("red")
	// get ""
	fmt.Printf("after SetColor return box color: %v\n", box2.Color())
	// get "red"
	fmt.Printf("after SetColor return boxCopy color: %v\n", boxCopy.Color())
	// "red"
	fmt.Printf("after SetColor2 return box color: %v\n", box3.Color())
}
