package main

import "fmt"

type Speaker interface {
    Speak()
}

type Person struct {
    name string
}

type Robot struct {
    id int
}

func (p Person) Speak() {
    fmt.Println("Hello, my name is", p.name)
}

func (r Robot) Speak() {
    fmt.Println("Beep! I am robot number", r.id)
}

func saySomething(s Speaker) {
    s.Speak()
}

func main() {
    p := Person{name: "Dinesh"}
    r := Robot{id: 101}

    saySomething(p)
    saySomething(r)
}
