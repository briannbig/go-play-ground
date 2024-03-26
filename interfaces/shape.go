package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math"
)

type Shape interface {
	Area() float64
	Perimeter() float64
}

type Rectangle struct {
	Width, Length float64
}

func (r Rectangle) String() string {
	out, err := json.MarshalIndent(r, "", "   ")

	if err != nil {
		log.Println(err)
	}
	return string(out)
}
func (c Circle) String() string {
	out, err := json.MarshalIndent(c, "", "   ")
	if err != nil {
		log.Println(err)
	}
	return string(out)
}

type Circle struct {
	Radius float64
}

func (r Rectangle) Area() float64 {
	return float64(r.Length) * float64(r.Width)
}

func (r Rectangle) Perimeter() float64 {
	return 2*r.Length + 2*r.Width
}

func (c Circle) Area() float64 {
	return math.Pi * c.Radius * c.Radius
}

func (c Circle) Perimeter() float64 {
	return 2 * math.Pi * c.Radius
}

func measure(s Shape) {
	fmt.Println("Shape:", s)
	fmt.Println(s.Area())
	fmt.Println(s.Perimeter())
}
