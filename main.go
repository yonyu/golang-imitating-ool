// This project is created inside GoPath
package main

import (
	"fmt"
	"golang-imitating-ool/pets"
	"time"
)

func main() {
	sleepTime := time.Now()
	//sleepTime := time.Now().Add(time.Duration(-5) * time.Hour)
	pet := pets.NewDog("Oreo", "Black and white", "Labrador", sleepTime)

	if pet.IsHungry() {
		fmt.Println(pet.Feed("kibble"))
	} else {
		fmt.Println("Pet is not hungry, waiting")
		time.Sleep(2 * time.Second)
		fmt.Println(pet.Feed("kibble"))
	}
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
