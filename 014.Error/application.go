package main

import "fmt"

func main() {

	person := Person{
		name: "dongchen",
		age:  12,
	}
	person.Run()
	person.Walk()
}

type WalkRun interface {
	Walk()

	Run()
}

type Person struct {
	name string
	age  int
}

func (person *Person) Walk() {
	fmt.Printf("name: %v can walk \n", person.name)
}

func (person *Person) Run() {
	fmt.Printf("name: %v can run \n", person.name)
}

type CommonError struct {
	errorCode int
	errorMsg  string
}

func (commonError *CommonError) Error() string {
	return commonError.errorMsg
}
