package main

import (
	"fmt"

	"github.com/chopnico/output"
)

type Dog struct {
	Name string
	Breed string
	Age int
	Colors []string
}

func main() {
	gershwin := Dog{
		Name: "Gershwin",
		Breed: "Cardigan Welsh Corgi",
		Age: 5,
		Colors: []string{"White", "Gray", "Brown", "Black"},
	}

	hildegard := Dog{
		Name: "Hildegard",
		Breed: "Cardigan Welsh Corgi",
		Age: 4,
		Colors: []string{"Black", "White", "Brown"},
	}

	mozart := Dog{
		Name: "Mozart",
		Breed: "Cardigan Welsh Corgi",
		Age: 2,
		Colors: []string{"Brown", "White"},
	}

	var entries []interface{}
	entries = append(entries, gershwin, hildegard, mozart)

	fmt.Println(output.FormatList(entries, []string{"Age", "Colors"}))
}
