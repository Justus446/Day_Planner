package main

import (
	"fmt"
	"time"
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