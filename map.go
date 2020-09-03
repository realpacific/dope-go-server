package main

import (
	"fmt"
	"math"
	"math/rand"
)

func main() {

	// Maps are unordered.
	mappings := map[string]int{
		"Rust":   90,
		"Python": 110,
		"Java":   0,
		"Kotlin": 1000000000000000,
	}

	mappings["JS"] = 105

	for key, value := range mappings {
		fmt.Println(value, key)
	}

	var points []Point
	for i := 0; i < 10; i++ {
		points = append(points, Point{rand.Intn(10), rand.Intn(20)})
	}

	for _, i := range points {
		fmt.Println(i)
	}

	fmt.Println(DistanceBetween(Point{5, 5}, Point{0, 0}))
}

type Point struct {
	X int
	Y int
}

func DistanceBetween(p1, p2 Point) float64 {
	return math.Sqrt(math.Pow(float64(p2.X-p1.X), 2) + math.Pow(float64(p2.Y-p1.Y), 2))
}
