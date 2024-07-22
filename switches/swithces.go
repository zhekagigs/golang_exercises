package main

import (
    "fmt"
	"strings"
)


// Min returns the smaller of two values
func Min[T Ordered ](values[]T) T {
    if len(values) == 0 {
        panic("Cannot find minimum of an empty slice")
    }
    min := values[0]
    for _, v := range values[1:] {
        if v.Less(min) {
            min = v
        }
    }
    return min
}

type Person struct{
	name string
	age int
}

type Ordered interface{
	Less(other Ordered) bool
}

func (p Person) Less(other Ordered) bool {
	otherPerson, ok := other.(Person)
    if !ok {
        panic("Invalid comparison")
    }
    // First compare by age
    if p.age != otherPerson.age {
        return p.age < otherPerson.age
    }
    // If ages are equal, compare by name
    return strings.Compare(p.name, otherPerson.name) < 0
}

func main() {
	people := []Person{
        {name: "Bob", age: 43},
        {name: "Alice", age: 22},
        {name: "Charlie", age: 35},
        {name: "David", age: 22},
    }

    minPerson := Min(people)
    fmt.Printf("Minimum person: %s, age %d\n", minPerson.name, minPerson.age)
}

