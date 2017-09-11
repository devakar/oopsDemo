package person

import (
	"fmt"
)

type Person struct {
	Name string
	Age  int
}

func (p *Person) Talk() {
	fmt.Println("Hello World")
}
