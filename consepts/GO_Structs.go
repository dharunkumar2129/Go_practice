package main

import "fmt"

type Student struct {
    name  string
    marks int
}

type Students []Student

func printPassed(students Students, passMark int) {
    for _, s := range students {
        if s.marks >= passMark {
            fmt.Println(s.name, "passed with", s.marks)
        } else {
            fmt.Println(s.name, "failed with", s.marks)
        }
    }
}

func increaseMarks(s *Student, bonus int) {
    s.marks += bonus
}

func main() {
    s1 := Student{"Alice", 45}
    s2 := Student{"Bob", 80}
    s3 := Student{"Charlie", 60}

    group := Students{s1, s2, s3}

    printPassed(group, 50)

    increaseMarks(&group[0], 10)
    increaseMarks(&group[2], 5)

    printPassed(group, 50)
}
