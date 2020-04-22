package main

import (
    "fmt"
)


// there is only on loop in go
// The "for" loop
func goLoop(){
    var arr1 [3]int // way to declare array
                    // by default array element are 0 for int
    var arr2 [3]int = [3]int{1, 2, 3}
    var arr3 [3]int = [3]int{1, 2}

    for i := range arr1 {
        fmt.Printf("%d\n", i)
    }

    // second varriable tells the value of element and is optional
    for index, value := range arr2{
        fmt.Printf("index: %d, value: %d\n", index, value)
    }

    // third demo of loops
    for _, value := range arr3 {
        fmt.Printf("value: %d\n", value)
    }


}

func main()  {
    goLoop()
}
