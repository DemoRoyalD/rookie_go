package main

import "fmt"

func main() {
	person := Person{
		name: "dongchen",
		age:  12,
	}
	fmt.Println(person)
	fmt.Println(person.getAge())
	fmt.Println(person.getName())
}

type Person struct {
	name string
	age  int
}

type Attribute interface {
	getAge() int
	getName() string
}

func (person *Person) getAge() int {
	return person.age
}

func (person *Person) getName() string {
	return person.name
}

// Implement and Composition
type Reader interface {
	Read(p []byte) (n int, err error)
}

type Writer interface {
	Write(p []byte) (n int, err error)
}

type ReadWriter interface {
	Reader
	Writer
}
