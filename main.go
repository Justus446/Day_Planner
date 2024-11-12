package main

import "fmt"

func main() {

	fmt.Println(("Hello, World!"))
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	fmt.Println("storage",storage)
	storage.Load(&todos)
	todos.toggle(1)
	todos.print()
	storage.Save(todos)
}