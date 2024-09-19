package main

import "fmt"

type Person struct {
	Name string
	Age  int
}

func main() {
	person := Person{Name: "John", Age: 30}
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
