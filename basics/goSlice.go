package basics

import (
    "fmt"
)

func BasicSlice() {
    fmt.Println("\n>>>>> Running BasicSlice()")
    months := [...]string{
        1: "January",
        2: "February",
        3: "March",
        4: "April",
        5: "May",
        6: "June",
        7: "July",
        8: "August",
        9: "September",
        10: "October",
        11: "Novemeber",
        12: "Decemeber",
    }

    Q2 := months[4:7]
    summer := months[6:9]
    fmt.Println(Q2)         // ["April" "May" "June"]
    fmt.Println(summer)     // ["June" "July" "August"]
}


// Slices cannot be compared using ==
// Only legal slice comparison is against nil
// if summer == nil { /* ... */ }


func AppendSlice() {
    fmt.Println("\n>>>>> Running AppendSlice()")
    var runes []rune
    for _, r := range "Hello, Rohit" {
        runes = append(runes, r)
    }
    fmt.Printf("%q\n", runes)           //  "['H' 'e' 'l' 'l' 'o' ',' ' ' 'R' 'o' 'h' 'i' 't']"
}


func appendInt(x []int, y int) []int {
    fmt.Println("\n>>>>> Running AppendInt()")
    var z []int
    zlen := len(x) + 1
    if zlen <= cap(x) {
        z = x[:zlen]                    // There is room to grow. Extend the Slices
    } else {
        // There is insufficient space. Allocate a new array.
        // Grow by doubling, for amortized linear complexity.
        zcap := zlen
        if zcap < 2*len(x) {
            zcap = 2 * len(x)
        }
        z = make([]int, zlen, zcap)
        copy(z, x)
    }
    z[len(x)] = y
    return z
}


func ReverseUsingPointer() {
    fmt.Println("\n>>>>> Running ReverseUsingPointer()")
    s := [5]int{1, 2, 3, 4, 5}
    var temp int

    for i, j := 0, len(s)-1; i<j; i, j = i+1, j-1 {
        temp = s[i]
        s[i] = *&s[j]
        s[j] = temp
    }

    fmt.Println(s)
}


func MapBasics(){
    ages := map[string]int {                        // "declaration"
        "alice": 31,
        "charlie": 34,
    }

    /*
    // Another way of declaration

    ages := make(map[string]int)
    ages["alice"] = 31
    ages["charlie"] = 34
    */

    fmt.Println(ages["alice"])                      // "32"

    delete(ages, "alice")                           // remove element ages["alice"]
}
