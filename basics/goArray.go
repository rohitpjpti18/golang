package basics

import (
    "fmt"
    "crypto/sha256"
    )

func BasicArray() {
    fmt.Println("\n>>>>> Running BasicArray()")
    // various types of declaration of array types.
    var arr1 [3]int                         // array of 3 integer
    var r [3]int = [3]int{1, 2}             // by default elements are initialized by "0"
    q := [...]int{1, 2, 3}

    // printing the results.
    fmt.Println(r[2])                       // prints "0"
    fmt.Printf("%T\n", q)                   // prints "[3]int"
    fmt.Println(arr1[0])                    // print the first element
    fmt.Println(arr1[len(arr1) - 1])        // print the last element, a[2]

    for index, value := range arr1 {
        fmt.Printf("%d %d\n", index, value) // print the indices and element.
    }

    for _, value := range arr1 {
        fmt.Printf("%d\n", value)           // print the elements only.
    }
}


func CompareArrayBasic() {
    a := [2]int{1, 2}
    b := [...]int{1, 2}
    c := [2]int{1, 3}
    fmt.Println(a == b, a == c, b == c)     // "true false false"
    /*
    d := [3]int{1, 2}
    fmt.Println(a == d)                     // compile error: cannot compare [2]int == [3]int
    */
}


func CompareArrayExample() {
    fmt.Println("\n>>>>> Running CompareArrayExample()")
    // example of comparing two arrays
    c1 := sha256.Sum256([]byte("x"))
    c2 := sha256.Sum256([]byte("X"))
    fmt.Printf("%x\n%x\n%t\n%T\n", c1, c2, c1 == c2, c1)
}


func PassArrayByPointer(ptr *[3]int) {
    fmt.Println("\n>>>>> Running PassArrayByPointer()")
    for i := range ptr {
        fmt.Println(ptr[i])
    }
}
