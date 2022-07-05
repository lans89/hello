package main

import (
	"fmt"

	greet "github.com/lans89/hello/Greet"
)

func main() {
	fmt.Println(greet.Esperanto("Ivan"))
	fmt.Println(greet.Spanish("Ivan"))
	fmt.Println(greet.English("Ivan"))
}
