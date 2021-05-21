package main

import (
	"fmt"
	"time"

	"github.com/chopnico/output"
)

type Dog struct {
	Name string
	Breed string
	Age int
	Colors []string
	Created time.Time
}

func main() {
	gershwin := Dog{
		Name: "Gershwin",
		Breed: "Cardigan Welsh Corgi",
		Age: 5,
		Colors: []string{"White", "Gray", "Brown", "Black"},
		Created: time.Now(),
	}

	hildegard := Dog{
		Name: "Hildegard",
		Breed: "Cardigan Welsh Corgi",
		Age: 4,
		Colors: []string{"Black", "White", "Brown"},
		Created: time.Now(),
	}

	mozart := Dog{
		Name: "Mozart",
		Breed: "Cardigan Welsh Corgi",
		Age: 2,
		Colors: []string{"Brown", "White"},
		Created: time.Now(),
	}

	var entries []interface{}
	entries = append(entries, gershwin, hildegard, mozart)

	fmt.Println(output.FormatList(entries, nil))
}
