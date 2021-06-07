# output

Inspired (a little) by PowerShell's object printing capabilities, output prints structs in a readable way.

## Install

``` sh
go get github.com/chopnico/output
```

## Examples

``` go
type Cow struct {
    Name        string
    Age         int
    Alive       bool
    Sound       string
    Color       []string
    Weight      float32
    Intervals   []int
    Toys        *[]Toy
    CreatedOn   time.Time
}

type Dog struct {
    Name        string
    Age         int
    Alive       bool
    Sound       string
    Color       []string
    Weight      float32
    Intervals   []int
    Toys        *[]Toy
    CreatedOn   time.Time
}

func main() {
    cow := Cow{
        Name: "Will",
        Age: 3,
        Alive: true,
        Sound: "moooo",
        Color: []string{"white", "black"},
        Weight: 3299.5,
        Intervals: []int{1,20,30,100},
        Toys: &[]Toy{
            {
                Name: "brush",
                Age: 3,
                Best: [3]int{1, 3, 5},
            },
            {
                Name: "ball",
                Age: 4,
                Best: [3]int{1, 3, 5},
                },
            },
        CreatedOn: time.Now(),
    }
    
    dog := Dog{
        Name: "Rain",
        Age: 2,
        Alive: true,
        Sound: "bark, bark",
        Color: []string{"brown", "black", "white"},
        Weight: 34.7,
        Intervals: []int{1,20,10,34},
        Toys: &[]Toy{
        {
                Name: "rope",
                Age: 8,
                Best: [3]int{1, 3, 5},
            },
            {
                Name: "ball",
                Age: 10,
                Best: [3]int{1, 3, 5},
            },
        },
        CreatedOn: time.Now(),
    }
    
    var animals []interface{}
    animals = append(animals, &dog, &cow)
    properties := []string{"Name", "Age", "Weight", "Toys", "Color", "CreatedOn"}
    
    fmt.Printf("%s\n", output.FormatList(&animals, properties))
}
```

Will output:

``` sh
Name      : Rain
Age       : 2
Weight    : 34.7
Toys      : &[{rope 8 [1 3 5]} {ball 10 [1 3 5]}]
Color     : [brown black white]
CreatedOn : 2021-06-06 16:22:42.430851209 -0400 EDT m=+0.000340650
          
Name      : Will
Age       : 3
Weight    : 3299.5
Toys      : &[{brush 3 [1 3 5]} {ball 4 [1 3 5]}]
Color     : [white black]
CreatedOn : 2021-06-06 16:22:42.430824503 -0400 EDT m=+0.000314053
```

