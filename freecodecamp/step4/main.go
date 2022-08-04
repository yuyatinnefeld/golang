package main

import (
	"fmt"
)

type Car struct {
	motor       string
	manifacture string
	speed       int
}

func main() {
	statePopulations := map[string]int{
		"Californina": 897988,
		"New York":    428932,
		"Ohio":        98098,
		"Hawaii":      8797,
	}
	fmt.Println(statePopulations)

	fmt.Println(statePopulations["Hawaii"])
	statePopulations["LA"] = 13213
	fmt.Println(statePopulations)
	delete(statePopulations, "LA")
	fmt.Println(statePopulations)

	pop, ok := statePopulations["Ohio"]
	fmt.Println(pop, ok)
	pop1, ok1 := statePopulations["Ohaiiii"]
	fmt.Println(pop1, ok1)

	aCar := Car{
		motor:       "greAudi Q7A4",
		manifacture: "Audi",
		speed:       7987,
	}

	fmt.Println(aCar)
	fmt.Println(aCar.manifacture)

	aDoctor := struct{ name string }{name: "John"}
	fmt.Println(aDoctor)

}
