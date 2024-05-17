// This project is created inside GoPath
package main

import (
	"fmt"
	"time"

	"golang-imitating-ool/pets"
)

func main() {
	//sleepTime := time.Now()
	sleepTime := time.Now().Add(time.Duration(-5) * time.Hour)
	pet := pets.NewDog("Oreo", "Black and white", "Labrador", sleepTime)

	fmt.Println(pet.Feed("steak"))
	fmt.Println(pet.GiveAttention("play fetch"))
}

func createDog() pets.Dog {
	pet := pets.Dog{
		Name:  "Billy",
		Color: "Black and white",
		Breed: "Labrador",
	}
	return pet
}
