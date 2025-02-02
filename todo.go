package main

import (
	"fmt"
	"os"
	"strconv"
	"time"

	"github.com/aquasecurity/table"
)

type todo struct {
	Title string
	Completed bool
	CreatedAt time.Time
	CompletedAt *time.Time
}

type Todos []todo

func (todos *Todos) add(title string){

	todo := todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error{
	if index < 0 || index >= len(*todos){
		return fmt.Errorf("Invalid index")
	}

	return nil
}

func (todos *Todos) delete(index int) error{
	if err := todos.validateIndex(index); err != nil {
		return err
	}

	*todos = append((*todos)[:index], (*todos)[index+1:]...)

	return nil

}

func (todos *Todos) toggle(index int) error{
	if err:= todos.validateIndex(index); err != nil{
		return err
	}

	isCompleted := (*todos)[index].Completed

	if !isCompleted {
		now := time.Now()
		(*todos)[index].CompletedAt = &now

	}
	(*todos)[index].Completed = !isCompleted

	return nil

}


func (todos *Todos) edit(index int, title string) error{
	if err:= todos.validateIndex(index); err != nil{
		return err
	}


	(*todos)[index].Title = title

	return nil

}


func (todos *Todos) print(){

	table := table.New(os.Stdout)
	table.SetHeaders("#", "Title", "Completed", "Created At", "Completed At","Reaction")
	table.SetRowLines(false)

	for index, todo := range *todos{
		completed := "❌"
		completedAt := ""
		reaction := "😞"

		if todo.Completed{
			completed = "✅"
			reaction = "😊🎊"
			if todo.CompletedAt != nil{
				completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")
			}
		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format("2006-01-02 15:04:05"), completedAt, reaction)


	}

	table.Render()
}