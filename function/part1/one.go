package part1

import (
	"fmt"
	"log"
	"math"
	"strings"
)

type StrFunc func(string) string

func componse(f StrFunc, g StrFunc) StrFunc {
	return func(s string) string {
		return g(f(s))
	}
}

func One() {
	var recognize = func(name string) string {
		return fmt.Sprintf("Hey %s", name)
	}

	var emphasize = func(statment string) string {
		return fmt.Sprintf(strings.ToUpper(statment) + "!")
	}

	var greetFog = componse(recognize, emphasize)
	log.Println(greetFog("Veloss"))
}

type square struct {
	X float64
}

func (s square) Area() float64 {
	return s.X * s.X
}

func (s square) Perimeter() float64 {
	return 4 * s.X
}

type circle struct {
	R float64
}

func (c circle) Area() float64 {
	return c.R * c.R * math.Pi
}

func (c circle) Perimeter() float64 {
	return 2 * c.R * math.Pi
}

type Shape interface {
	Area() float64
	Perimeter() float64
}

func C() Shape {
	c := circle{}
	return c
}

func S() Shape {
	s := square{}
	return s
}

func One1() {
	cvalue := C()
	svalue := S()
}
