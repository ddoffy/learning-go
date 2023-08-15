// File: main.go
package main

import (
	"encoding/json"
	"fmt"
)

// main is the entry point for the application.
func main() {
	// try to convert string to []byte

	str := "Hello, World!"
	bytes := []byte(str)

	fmt.Println("Bytes:", string(bytes))

	// revert bytes to string
	str2 := string(bytes)
	fmt.Println("String:", str2)

	toFile := Person{
		Name: "Fred",
		Age:  40,
	}

	data, err := json.Marshal(toFile)
	if err != nil {
		panic(err)
	}

	fmt.Println("JSON:", string(data))

	// unmarshal data from file

	var p Person
	err = json.Unmarshal(data, &p)
	if err != nil {
		panic(err)
	}
	fmt.Println("Person:", p)

}

// Person the struct of a person
type Person struct {
	Name string `json:"name"`
	Age  int    `json:"age"`
}
