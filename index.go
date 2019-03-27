package main

import "fmt"
import m "math"

func main() {
	i := 0
	for i < 5 {
		fmt.Println("Hello world")
		i = i + 1
	}

	fmt.Println("----------------")
	fmt.Println("Rectangle")
	r := rect{ length: 25, breadth: 10 }
	fmt.Println(r.area())
	fmt.Println(r.perimeter())

	fmt.Println("----------------")
	fmt.Println("Circle")
	c := circle{ radius: 5 }
	fmt.Println(c.area())
	fmt.Println(c.perimeter())

	// Using generics eqvt
	measure(r)
}

type shape interface {
	area() float64
	perimeter() float64
}

type rect struct {
	length float64
	breadth float64
}

func (r rect) area() float64 {
	return r.length * r.breadth
}


func (r rect) perimeter() float64 {
	return 2 * (r.length + r.breadth)
}

type circle struct {
	radius float64
}

func (c circle) area() float64 {
	return m.Pi * c.radius * c.radius
}


func (c circle) perimeter() float64 {
	return m.Pi * 2 * c.radius
}

func measure(s shape) {
	fmt.Println(s.area())
	fmt.Println(s.perimeter())
}