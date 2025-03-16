package main

import (
	"fmt"
	"math"
)

func printText() {
	fmt.Println("Hello inside")
}

func circleArea(rad float32) float32 {
	return rad * 2 * math.Pi
}

func main() {
	printText()
	fmt.Printf("Rad of circle: %d \nArea: %f \n", 2, circleArea(2))
	// Anonymous func

	num := 0
	fmt.Println(num)
	increment := func() int {
		num++
		return num
	}
	increment()
	fmt.Println(num)
}
