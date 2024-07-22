package main

import (
    "fmt"
)

type Printable interface {
    String() string
}

func PrintAll[T Printable](items []T) int {
	var count int
    for _, item := range items {
		str := item.String()
		fmt.Println(str)
		count += len(str)
    }
	return count
}

type Address struct{
	PostCode string
	FirstLine string
}

func (a Address) String() string {
	return fmt.Sprintf("%s at %s", a.FirstLine, a.PostCode)
}

func main() {

	addresses := []Address {
		{PostCode: "SE11 3ER", FirstLine: "Chumbers Grove" },
		{PostCode: "PH2", FirstLine: "Ho Road"},
	}

	PrintAll(addresses)
}

// Can you create a new type that satisfies the Printable interface and use it with the 
// PrintAll function? How would you modify the PrintAll function to also return the 
// total number of characters printed?
