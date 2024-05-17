package pets

import (
	"fmt"
	"time"
)

type Cat struct {
	Name  string
	Color string
	Breed string
	Animal
}

func (c Cat) GiveAttention(activity string) string {
	return fmt.Sprintf("The cat is ignoring your attempts to %s", activity)
}

func NewCat(name, color, breed string) Pet {
	// how to fix the problem the function cannot use Cat{…} (value of type Cat) as Pet value in return statement
	// because Cat does not implement Pet (missing GiveAttention method)
	// return Cat{Name: name, Color: color, Breed: breed}

	return Cat{Name: name, Color: color, Breed: breed, Animal: Animal{lastAte: time.Now()}}
}
