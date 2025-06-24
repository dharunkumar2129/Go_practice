package main

import "fmt"

func main() {
    arr := [5]int{10, 25, 30, 45, 50}

    fmt.Println("Array elements and check if even or odd:")
    for i := 0; i < len(arr); i++ {
        isEven := arr[i]%2 == 0
        if isEven {
            fmt.Printf("arr[%d] = %d is Even\n", i, arr[i])
        } else {
            fmt.Printf("arr[%d] = %d is Odd\n", i, arr[i])
        }
    }

    slice := arr[2:]

    fmt.Println("\nSlice elements and check if greater than 40:")
    for i, val := range slice {
        if val > 40 {
            fmt.Printf("slice[%d] = %d is Greater than 40\n", i, val)
        } else if val == 40 {
            fmt.Printf("slice[%d] = %d is Equal to 40\n", i, val)
        } else {
            fmt.Printf("slice[%d] = %d is Less than 40\n", i, val)
        }
    }

    allAbove10 := true
    for _, val := range arr {
        if val <= 10 {
            allAbove10 = false
            break
        }
    }

    fmt.Println("\nAll elements greater than 10:", allAbove10)
}
