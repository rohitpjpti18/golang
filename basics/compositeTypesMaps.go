package basics

import (
    "fmt"
    "sort"
)

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

    for name, age := range ages {
        fmt.Printf("%s\t%d\n", name, age)
    }

    var names []string
    for name := range ages {
        names = append(names, name)
    }

    sort.Strings(names)

    for _, name := range names {
        fmt.Printf("%s\t%d\n", name, ages[name])
    }
}
