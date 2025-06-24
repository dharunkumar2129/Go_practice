package main

import "fmt"

func main() {
    var arr [5]int
    arr[0] = 10
    arr[1] = 20
    arr[2] = 30
    arr[3] = 40
    arr[4] = 50

    fmt.Println("Array elements:")
    for i, val := range arr {
        fmt.Printf("arr[%d] = %d\n", i, val)
    }

    slice := arr[1:4]

    fmt.Println("\nSlice elements (from arr[1:4]):")
    for i, val := range slice {
        fmt.Printf("slice[%d] = %d\n", i, val)
    }

    slice = append(slice, 60, 70)
    fmt.Println("\nAfter appending to slice:")
    fmt.Println("Slice:", slice)

    slice[0] = 99
    fmt.Println("\nAfter modifying slice[0] = 99:")
    fmt.Println("Original array:", arr)
    fmt.Println("Slice:", slice)

    newSlice := make([]int, 3)
    newSlice[0] = 1
    newSlice[1] = 2
    newSlice[2] = 3
    fmt.Println("\nNew slice using make():", newSlice)
}


