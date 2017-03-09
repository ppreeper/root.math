package main

import (
    "fmt"
    "math"
)

func root(x, p float64) float64 {
    y := x/p
    for i :=0; i < 32; i++ {
        y=(x/math.Pow(y,(p-1))+y*(p-1))/p
    }
    return y
}


func main(){
    y := root(5,2)
    fmt.Println(y)
    y = root(27,3)
    fmt.Println(y)
    y = root(29,3)
    fmt.Println(y)
}
