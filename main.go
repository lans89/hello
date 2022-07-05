package main

import (
	"fmt"

	"github.com/lans89/hello/Greet"
)

func main() {
	fmt.Println(Greet.Esperanto("Ivan"))
	fmt.Println(Greet.Spanish("Ivan"))
	fmt.Println(Greet.English("Ivan"))

	fmt.Println("")
}
