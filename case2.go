package main

import "fmt"

// Define the interface
type Animal interface {
	Speak() string
}

// Define a struct for Dog
type Dog struct {
	Name string
}

// Define a struct for Cat
type Cat struct {
	Name string
}

// Implement the Speak method for Dog
func (d Dog) Speak() string {
	return d.Name + " says Woof!"
}

// Implement the Speak method for Cat
func (c Cat) Speak() string {
	return c.Name + " says Meow!"
}

// Function that accepts any type that satisfies the Animal interface
func makeAnimalSpeak(a Animal) {
	fmt.Println(a.Speak())
}

func main() {
	dog := Dog{Name: "Buddy"}
	cat := Cat{Name: "Whiskers"}

	// Both Dog and Cat satisfy the Animal interface
	makeAnimalSpeak(dog)
	makeAnimalSpeak(cat)
}

/**
Penjelasan:
1.	type Animal interface { Speak() string }: Ini adalah deklarasi antarmuka bernama Animal. Antarmuka ini mendefinisikan satu metode Speak(), yang mengembalikan sebuah string. Setiap tipe yang mengimplementasikan metode ini akan otomatis mengadopsi antarmuka Animal.
2.	Struct Dog dan Cat: Dua struktur ini masing-masing memiliki bidang Nama dan mereka akan mengimplementasikan metode Speak().
3.	func (m Dog) Speak() string: Ini adalah implementasi metode Speak() untuk tipe Dog. Hal yang sama berlaku untuk Cat. Mereka memiliki cara yang berbeda untuk mengimplementasikan metode tersebut.
4.	func makeAnimalSpeak(p Animal): Fungsi ini menerima parameter tipe Animal, yang berarti bisa menerima tipe apapun yang mengimplementasikan metode Speak(). Dalam hal ini, baik Dog maupun Cat bisa diterima oleh fungsi ini.
**/
