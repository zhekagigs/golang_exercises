package main

import (
	"errors"
	"fmt"
	"reflect"
)

type Stack[T any] struct {
    items []T
}

func (s *Stack[T]) Push(item T) {
    s.items = append(s.items, item)
}

func (s *Stack[T]) Pop() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, errors.New("stack is empty")
    }
    item := s.items[len(s.items)-1]
    s.items = s.items[:len(s.items)-1]
    return item, nil
}

func (s *Stack[T]) Peek() (T, error) {
    var zero T
    if len(s.items) == 0 {
        return zero, errors.New("stack is empty")
    }
    return s.items[len(s.items)-1], nil
}

func (s *Stack[T]) IsEmpty() bool {
    return len(s.items) == 0
}

func (s *Stack[T]) Find(search T) (T, error) {
	var zero T
	if len(s.items) == 0 {
		return zero, errors.New("stack is empty") 
	}
	for _, v := range s.items {
		if reflect.DeepEqual(v, search) {
			return v, nil	
		} 
	}
	return zero, errors.New("nothing found")
	
}

func main() {
    intStack := &Stack[int]{}
    intStack.Push(1)
    intStack.Push(2)
    intStack.Push(3)

	val, _ := intStack.Find(2)
	fmt.Println(val)

    for !intStack.IsEmpty() {
        item, _ := intStack.Pop()
        fmt.Println(item)
    }

    stringStack := &Stack[string]{}
    stringStack.Push("hello")
    stringStack.Push("world")

    for !stringStack.IsEmpty() {
        item, _ := stringStack.Pop()
        fmt.Println(item)
    }

}