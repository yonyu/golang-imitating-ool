package pets

import (
	"fmt"
	"strings"
	"time"
)

// Dog struct allows you to store data and provide methods on that data
type Dog struct {
	Name  string
	Color string
	Breed string
	//Age       int
	lastSlept time.Time // not accessible outside the package
}

func (d Dog) needSleep() bool { // inaccessible outside the package
	return time.Now().Sub(d.lastSlept) > time.Hour*4
}

func (d Dog) Sleep() {
	d.lastSlept = time.Now()
}

func (d Dog) Feed(food string) string {
	return fmt.Sprintf("%s is eating %s", d.Name, food)
}

func (d Dog) GiveAttention(activity string) string {
	if d.needSleep() {
		d.Sleep()
		return fmt.Sprintf("Your dog %s is tired and needs to rest", d.Name)
	}
	response := ""
	switch strings.ToUpper(activity) {
	case "PET":
		response = "wags tail"
	case "PLAYING FETCH":
		response = "return the ball and jump waiting for you to throw it again"
	default:
		response = "barks"
	}
	return fmt.Sprintf("%s loves attention, %s will cause him to %s", d.Name, activity, response)
}

func NewDog(name, color, breed string, lastSlept time.Time) Dog {
	return Dog{
		Name:      name,
		Color:     color,
		lastSlept: lastSlept,
		Breed:     breed,
	}
}
