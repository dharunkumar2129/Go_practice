package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
)

func main() {
    reader := bufio.NewReader(os.Stdin)

    fmt.Print("Enter text to save in file: ")
    input, _ := reader.ReadString('\n')
    input = strings.TrimSpace(input)

    file, err := os.Create("output.txt")
    if err != nil {
        fmt.Println("Error creating file:", err)
        return
    }
    defer file.Close()

    writer := bufio.NewWriter(file)
    _, err = writer.WriteString(input + "\n")
    if err != nil {
        fmt.Println("Error writing to file:", err)
        return
    }
    writer.Flush()

    fmt.Println("Data saved to output.txt")
}
