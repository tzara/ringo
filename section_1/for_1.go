package main

import "fmt"

type Rectangle struct {
	width, height int
}

func rectangleArea(rect *Rectangle) int {
	return rect.width * rect.height
}

func (rect *Rectangle) area() int {
	return rect.width * rect.height
}

func main() {
	sum := 0

	for i := 0; i < 10; i++ {
		sum += i
	}

	fmt.Println(sum)

	rect := Rectangle{30, 20}
	area := rectangleArea(&rect)
	fmt.Println(area)
	fmt.Println(rect.area())
}
