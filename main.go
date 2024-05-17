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
	var animals []pets.Pet
	animals = append(animals, pets.NewDog("Oreo", "Black and white", "Labrador", sleepTime))
	animals = append(animals, pets.NewCat("Milo", "Black", "Siamese"))
	//pet := pets.NewDog("Oreo", "Black and white", "Labrador", sleepTime)
	for _, pet := range animals {
		if pet.IsHungry() {
			fmt.Println(pet.Feed("kibble"))
		} else {
			fmt.Println("Pet is not hungry, waiting")
			time.Sleep(2 * time.Second)
			fmt.Println(pet.Feed("kibble"))
		}
		fmt.Println(pet.GiveAttention("play fetch"))
	}
}
