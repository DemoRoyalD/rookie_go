package main

import (
	"encoding/json"
	"fmt"
	"reflect"
)

func main() {
	// reflect 获取类型
	variable := 1
	vr := reflect.ValueOf(variable)
	vt := reflect.TypeOf(variable)
	fmt.Println(vr, vt)
	person := Person{
		Name: "dongchen",
		Age:  12,
	}
	fmt.Println(person)

	// struct -> json
	jsonB, err := json.Marshal(person)
	if err == nil {
		fmt.Println(string(jsonB))
	}
	fmt.Println("==============")
	// json to struct
	respJson := "{\"name\":\"李四\",\"age\":40}"
	json.Unmarshal([]byte(respJson), &person)
	fmt.Println(person)
}

type Person struct {
	Name string
	Age  int
}

func NewPerson(name string, age int) *Person {
	return &Person{Name: name, Age: age}
}
