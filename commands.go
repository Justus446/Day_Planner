package main

import (
	"flag"
	"fmt"
	"os"
	"strconv"
	"strings"
)

type CmdFlags struct {
    List   bool
    Add    string
    Edit   string
    Delete int
	Toggle int
	Help bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.BoolVar(&cf.List, "list", false, "List all tasks")
	flag.StringVar(&cf.Add, "add", "", "Add a new task")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a task")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task/mark as completed")
	flag.BoolVar(&cf.Help, "help", false, "Show all the available flags and their description")


	flag.Parse()

	return &cf
}


func (cf *CmdFlags) Execute(todos *Todos){
	switch{
		case cf.List:
			todos.print()
		case cf.Add != "":
			todos.add(cf.Add)
		case cf.Edit != "":
			splitter := strings.Split(cf.Edit, ":")
			if len(splitter) != 2 {
				fmt.Println("Invalid edit format. Use index:New Title")
				os.Exit(1)
			}
			
			index,err := strconv.Atoi(splitter[0])

			if err != nil {
				fmt.Println("Error: Invalid index for edit")
				os.Exit(1)
			}

			todos.edit(index, splitter[1])

		case cf.Delete != -1:
			todos.delete(cf.Delete)
		case cf.Toggle != -1:
			todos.toggle(cf.Toggle)
		case cf.Help:
			fmt.Println("Available commands:")
			flag.PrintDefaults()
			os.Exit(0)
		default:
			fmt.Println("Invalid command")

		}

}
