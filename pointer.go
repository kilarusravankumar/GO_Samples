package main

import (
	"fmt"
)

func makeGreeting(g *string) {
	*g = *g + " I'M Sravan Kumar Kilaru, Nice to meet you."
}
func main() {
	g := "Hello!"
	makeGreeting(&g)
	fmt.Println(g)
}
