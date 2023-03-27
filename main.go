package main

import (
	"fmt"
	"github.com/markossilva/go-study/abstract-factory"
	"github.com/markossilva/go-study/builder"
)

func main() {

	// Abstract Factory
	adidasFactory, _ := abstractfactory.AdidasType.GetFactory()
	adidasShoes := adidasFactory.MakeShoes()
	printDetailsShoes(adidasShoes)

	// Builder
	lBuilder := builder.Normal.GetBuilder()
	lBuilder.SetDoorType("teste")
	lBuilder.SetWindowType("window")
	lBuilder.SetNumFloor(123)

	director := builder.NewDirector(lBuilder)
	printBuilderDetails(director.BuildHouse())

	getBuilder := builder.Igloo.GetBuilder()
	getBuilder.SetDoorType("Snow Door")
	getBuilder.SetWindowType("Snow Window")
	getBuilder.SetNumFloor(321)
	director.SetBuilder(getBuilder)
	printBuilderDetails(director.BuildHouse())
}

func printBuilderDetails(house builder.House) {
	fmt.Printf("DoorType: %s", house.DoorType)
	fmt.Println()
	fmt.Printf("WindowType: %s", house.WindowType)
	fmt.Println()
	fmt.Printf("Floor: %d", house.Floor)
	fmt.Println()
	fmt.Println()
}

func printDetailsShoes(s abstractfactory.IShoe) {
	fmt.Printf("Logo: %s", s.GetLogo())
	fmt.Println()
	fmt.Printf("Size: %d", s.GetSize())
	fmt.Println()
	fmt.Println()
}
