package main

import (
	"fmt"
)

func plusMinus(arr []int32) {
    pos := 0
    neg := 0
    zero := 0
    for _, num := range arr{
        if num == 0 {
            zero += 1            
        }
        if num > 0 {
            pos += 1            
        }
        if num < 0 {
            neg += 1                 
        }
    }
    fmt.Println(float32(pos) / float32(len(arr)))
    fmt.Println(float32(neg) / float32(len(arr)))
    fmt.Println(float32(zero) / float32(len(arr)))
    
}


func main() {
	arr := []int32{1,1,0,-1,-1, -1}
	plusMinus(arr)
}