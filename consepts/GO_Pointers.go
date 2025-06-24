package main

import "fmt"

func doubleValue(num *int) {
    *num = *num * 2
}

func updateArray(arr *[3]int) {
    for i := 0; i < len(arr); i++ {
        arr[i] = arr[i] + 5
    }
}

func main() {
    x := 10
    fmt.Println("Before doubling:", x)
    doubleValue(&x)
    fmt.Println("After doubling:", x)

    a := [3]int{1, 2, 3}
    fmt.Println("Before update:", a)
    updateArray(&a)
    fmt.Println("After update:", a)
}
