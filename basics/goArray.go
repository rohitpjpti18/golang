package basics

import "fmt"

func goArray() {
    // various types of declaration of array types.
    var arr1 [3]int                         // array of 3 integer
    var r [3]int = [3]int{1, 2}             // by default elements are initialized by "0"
    q := [...]int{1, 2, 3}

    fmt.Println(r[2])                       // prints "0"
    fmt.Printf("%T\n", q)                  // prints "[3]int"

    fmt.Println(arr1[0])                    // print the first element
    fmt.Println(arr1[len(arr1) - 1])        // print the last element, a[2]

    // print the indices and element.
    for index, value := range arr1 {
        fmt.Printf("%d %d\n", index, value)
    }

    // print the elements only.
    for _, value := range arr1 {
        fmt.Printf("%d\n", value)
    }
}
