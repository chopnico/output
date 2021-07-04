package main

import (
	"fmt"
	"time"

	"github.com/chopnico/output"
)

type Animal interface {
	Speak()
}

type Dogs []Dog

type Cow struct {
	Name      string
	Age       int
	Alive     bool
	Sound     string
	Color     []string
	Weight    float32
	Intervals []int
	Toys      *[]Toy
	CreatedOn time.Time
}

type Dog struct {
	Name      string
	Age       int
	Alive     bool
	Sound     string
	Color     []string
	Weight    float32
	Intervals []int
	Toys      *[]Toy
	CreatedOn time.Time
}

type Toy struct {
	Name string
	Age  int
	Best [3]int
}

func (d *Dog) Speak() {
	fmt.Println(d.Sound)
}

func (c *Cow) Speak() {
	fmt.Println(c.Sound)
}

func Speak(a Animal) {
	a.Speak()
}

func main() {
	cow := Cow{
		Name:      "Will",
		Age:       3,
		Alive:     true,
		Sound:     "moooo",
		Color:     []string{"white", "black"},
		Weight:    3299.5,
		Intervals: []int{1, 20, 30, 100},
		Toys: &[]Toy{
			{
				Name: "brush",
				Age:  3,
				Best: [3]int{1, 3, 5},
			},
			{
				Name: "ball",
				Age:  4,
				Best: [3]int{1, 3, 5},
			},
		},
		CreatedOn: time.Now(),
	}

	dog := Dog{
		Name:      "Rain",
		Age:       2,
		Alive:     true,
		Sound:     "bark, bark",
		Color:     []string{"brown", "black", "white"},
		Weight:    34.7,
		Intervals: []int{1, 20, 10, 34},
		Toys: &[]Toy{
			{
				Name: "rope",
				Age:  8,
				Best: [3]int{1, 3, 5},
			},
			{
				Name: "ball",
				Age:  10,
				Best: [3]int{1, 3, 5},
			},
		},
		CreatedOn: time.Now(),
	}

	var dogs Dogs
	dogs = append(dogs, dog)

	properties := []string{"Name", "Age", "Weight", "Toys", "Color", "CreatedOn"}

	fmt.Printf("%s\n", output.FormatItemsAsList(dogs, properties))
	fmt.Printf("%s\n", output.FormatItemsAsList(&dogs, properties))
	fmt.Printf("%s\n", output.FormatItemAsList(dog, nil))
	fmt.Printf("%s\n", output.FormatItemAsList(&cow, nil))
	fmt.Printf("%s\n", output.FormatItemsAsJson(dogs))
	fmt.Printf("%s\n", output.FormatItemsAsJson(&dogs))
	fmt.Printf("%s", output.FormatItemsAsPrettyJson(dogs))
	fmt.Printf("%s", output.FormatItemsAsPrettyJson(&dogs))
}
