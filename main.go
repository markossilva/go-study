package main

import (
	"fmt"
	. "github.com/markossilva/go-study/abstract-factory"
)

func main() {
	adidasFactory, _ := AdidasType.GetFactory()
	adidasShoes := adidasFactory.MakeShoes()
	printDetailsShoes(adidasShoes)
}

func printDetailsShoes(s IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
}
