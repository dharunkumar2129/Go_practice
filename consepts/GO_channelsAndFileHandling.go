package main

import (
    "bufio"
    "fmt"
    "os"
    "strings"
    "sync"
)

func ProcessLogs(inputFiles []string, outputFile string) error {
    errorChan := make(chan string)
    var wg sync.WaitGroup

    for _, filePath := range inputFiles {
        wg.Add(1)
        go func(path string) {
            defer wg.Done()
            file, err := os.Open(path)
            if err != nil {
                fmt.Println("Error opening file:", path, err)
                return
            }
            defer file.Close()

            scanner := bufio.NewScanner(file)
            for scanner.Scan() {
                line := scanner.Text()
                if strings.Contains(line, "ERROR") {
                    errorChan <- line
                }
            }

            if err := scanner.Err(); err != nil {
                fmt.Println("Error reading file:", path, err)
            }
        }(filePath)
    }

    go func() {
        wg.Wait()
        close(errorChan)
    }()

    outFile, err := os.Create(outputFile)
    if err != nil {
        return fmt.Errorf("failed to create output file: %w", err)
    }
    defer outFile.Close()

    writer := bufio.NewWriter(outFile)
    for line := range errorChan {
        _, err := writer.WriteString(line + "\n")
        if err != nil {
            return fmt.Errorf("error writing to output file: %w", err)
        }
    }

    if err := writer.Flush(); err != nil {
        return fmt.Errorf("error flushing to output file: %w", err)
    }

    return nil
}

func main() {
    inputFiles := []string{"server1.log", "server2.log", "server3.log"}
    err := ProcessLogs(inputFiles, "errors.log")
    if err != nil {
        fmt.Println("Error:", err)
    } else {
        fmt.Println("Error messages extracted successfully.")
    }
}
