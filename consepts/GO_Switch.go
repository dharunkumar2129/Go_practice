package main

import (
    "fmt"
)

func main() {
    var choice int

    fmt.Println("Enter a number (1-3):")
    fmt.Scanf("%d", &choice)

    switch choice {
    case 1:
        fmt.Println("You chose One")
    case 2:
        fmt.Println("You chose Two")
    case 3:
        fmt.Println("You chose Three")
    default:
        fmt.Println("Invalid choice")
    }
}
