package main

import "fmt"

// Mendefinisikan struktur Orang
type Person struct {
	Name string
	Age  int
}

func main() {
	// Membuat instance dari struktur Orang
	person := Person{Name: "John", Age: 30}
	// Menampilkan data orang
	fmt.Printf("Name: %s, Age: %d\n", person.Name, person.Age)
}
