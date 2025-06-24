package main

import "fmt"

func countEvenOdd(nums []int) (int, int) {
    even := 0
    odd := 0
    for _, n := range nums {
        if n%2 == 0 {
            even++
        } else {
            odd++
        }
    }
    return even, odd
}

func main() {
    arr := [5]int{10, 15, 20, 25, 30}
    slice := arr[:]

    even, odd := countEvenOdd(slice)

    fmt.Println("Array:", arr)
    fmt.Println("Even count:", even)
    fmt.Println("Odd count:", odd)
}
