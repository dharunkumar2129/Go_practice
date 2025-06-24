package main

import "fmt"

func main() {
    var arr = [5]int{10, 20, 30, 40, 50}

    fmt.Println("Using for loop with array:")
    for i := 0; i < len(arr); i++ {
        fmt.Printf("arr[%d] = %d\n", i, arr[i])
    }

    fmt.Println("\nUsing for-range loop with array:")
    for i, val := range arr {
        fmt.Printf("arr[%d] = %d\n", i, val)
    }

    slice := arr[1:4]

    fmt.Println("\nUsing for loop with slice:")
    for i := 0; i < len(slice); i++ {
        fmt.Printf("slice[%d] = %d\n", i, slice[i])
    }

    fmt.Println("\nUsing for-range loop with slice:")
    for i, val := range slice {
        fmt.Printf("slice[%d] = %d\n", i, val)
    }
}
