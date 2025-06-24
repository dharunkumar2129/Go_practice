package main

import "fmt"

type Person struct {
    name string
    age  int
}

func (p *Person) Grow() {
    p.age++
}

func (p Person) Show() {
    fmt.Println(p.name, "is", p.age, "years old")
}

func main() {
    p := Person{"Dinesh", 20}
    p.Show()
    p.Grow()
    p.Show()
}
