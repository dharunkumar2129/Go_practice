package main

import (
    "bufio"
    "fmt"
    "os"
    "strconv"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter a float number: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    num, err := strconv.ParseFloat(input, 64)
    if err != nil {
        fmt.Println("Invalid float:", err)
        return
    }

    fmt.Println("Parsed float:", num)
}
