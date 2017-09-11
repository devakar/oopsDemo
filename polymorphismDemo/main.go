package main

import (
	"fmt"

	"github.com/oopsDemo/polymorphismDemo/shape"
)

func main() {
	var shapes = []shape.Shape{shape.Circle{1.0}, shape.Square{5.0}, shape.Rectangle{1.5, 2.0}}
	for _, shape := range shapes {
		fmt.Println(shape.Area())
	}
}
