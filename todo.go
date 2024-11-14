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
	Deadline *time.Time
}

type Todos []todo

func (todos *Todos) add(title string, deadline_optional ...time.Time) {
	var deadline *time.Time

	if len(deadline_optional) > 0 {
		deadline = &deadline_optional[0]
	}

	todo := todo{
		Title: title,
		Completed: false,
		CreatedAt: time.Now(),
		CompletedAt: nil,
		Deadline: deadline,
	}

	*todos = append(*todos, todo)
}

func (todos *Todos) validateIndex(index int) error{
	if index < 0 || index >= len(*todos){
		return fmt.Errorf("index out of range")
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
	table.SetHeaders("#", "Title", "Completed", "Created At", "Deadline","Completed At","Reaction")
	table.SetRowLines(false)

	for index, todo := range *todos{
		completed := "❌"
		completedAt := ""
		reaction := "👎"
		deadline := ""
		

		if todo.Completed{
			completed = "✅"
			reaction = "😊👍"
			if todo.CompletedAt != nil{
				completedAt = todo.CompletedAt.Format("2006-01-02 15:04:05")

				if todo.Deadline != nil && todo.CompletedAt.Before(*todo.Deadline) {
					reaction = "😊👍🎉🎊"
				}
			}
		}

		if todo.Deadline != nil{
			deadline = todo.Deadline.Format("2006-01-02 15:04")

		}

		table.AddRow(strconv.Itoa(index), todo.Title, completed, todo.CreatedAt.Format("2006-01-02 15:04:05"), deadline, completedAt, reaction)
	}

	table.Render()
}


func (todos *Todos) refresh(to_delete string) {
    var result Todos

    for _, todo := range *todos {
        if to_delete == "completed" {
            if !todo.Completed {
                result = append(result, todo)
            }
        } else if to_delete == "pending" {
            if todo.Completed {
                result = append(result, todo)
            }
        } else {
            // Clear all elements
            result = result[:0]
        }
    }

    *todos = result
}
