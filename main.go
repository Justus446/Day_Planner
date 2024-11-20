package main

import (
	"fmt"
	"sync"
)

func main() {
	fmt.Println("Failing to plan is planning to fail!")
	todos := Todos{}
	storage := NewStorage[Todos]("todos.json")
	storage.Load(&todos)

	var wg sync.WaitGroup
	
	wg.Add(1)
	go func() {
		defer wg.Done()
		RunNotifications(&todos)
	}()

	//  command execution
	CmdFlags := NewCmdFlags()
	CmdFlags.Execute(&todos)

	storage.Save(todos)

	wg.Wait()
}