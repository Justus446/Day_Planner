package main

import "fmt"

func main() {

	fmt.Println(("Failing to plan is planning to fail!"))
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)
	storage.Save(todos)
}