package main

import "fmt"

func countAboveThreshold(m map[string]int, threshold int) (int, int) {
    above := 0
    belowOrEqual := 0
    for _, value := range m {
        if value > threshold {
            above++
        } else {
            belowOrEqual++
        }
    }
    return above, belowOrEqual
}

func main() {
    scores := map[string]int{
        "Alice":  75,
        "Bob":    60,
        "Charlie": 85,
        "David":  40,
    }

    fmt.Println("Student Scores:")
    for name, score := range scores {
        fmt.Printf("%s: %d\n", name, score)
    }

    above, below := countAboveThreshold(scores, 65)

    fmt.Println("Scores above 65:", above)
    fmt.Println("Scores 65 or below:", below)
}
