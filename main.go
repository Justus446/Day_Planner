package main

import "fmt"

func main() {

	fmt.Println(("Hello, World!"))
	todos := Todos{}
	todos.add("Learn Go")
	todos.add("Create a CLI App")
	todos.add("Learn Go backend")
	fmt.Printf("%+v\n", todos)
	todos.delete(1)
	fmt.Printf("%+v\n", todos)
}