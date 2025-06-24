package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter your full name: ")
    nameInput, _ := reader.ReadString('\n')
    name := strings.TrimSpace(nameInput)

    var age int
    fmt.Print("Enter your age: ")
    fmt.Scanf("%d", &age)

    fmt.Println("Name:", name)
    fmt.Println("Age:", age)
}
