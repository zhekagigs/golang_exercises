package main

import (
	"fmt"

	"golang.org/x/exp/constraints"
)

func Sum[T constraints.Integer | constraints.Float | string] (numbers []T) T {
	var sum T

	for _, num := range numbers {
		sum += num
	}
	return sum
}

type Numeric interface{
	constraints.Float | constraints.Integer
}

func Avg[T Numeric] (numbers [] T) float32 {
	avg := float32(Sum(numbers)) / float32(len(numbers))
	return float32(avg)
}

func main() {
	ints := []int{1,2,4,8,16,32}
	floats := []float32{0.1, 0.2, 0.16, 0.32}
	runes := []rune{'a', 'b', 'c'}
	strings := []string{"a", "b", "c"}
	// strings2 := []string{"d", "e", "f"}
	// fmt.Println(strings + strings2)
	fmt.Println(Sum(ints))
	fmt.Println(Sum(floats))
	fmt.Println(Sum(runes))
	fmt.Println(Sum(strings))
}