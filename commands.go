package main

import (
	"flag"
	"fmt"
	"os"
	"regexp"
	"strconv"
	"strings"
	"time"
)

type CmdFlags struct {
    List   bool
    Add    string
    Edit   string
    Delete int
	Toggle int
	Help bool
	Refresh bool
}

func validateDateTime(date string) error {
	regex_date_hour :=  `^(?:(\d{4})-(0[1-9]|1[0-2])-(0[1-9]|[12][0-9]|3[01])(?: (?:[01][0-9]|2[0-3]):[0-5][0-9])?|(?:[01][0-9]|2[0-3]):[0-5][0-9])$`

	re := regexp.MustCompile(regex_date_hour)

	if re.MatchString(date) {
		return nil
	} else {
		return fmt.Errorf("error: invalid date format: use YYYY-MM-DD HH:MM, or YYYY-MM-DD or HH:MM")
	}
}


func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.BoolVar(&cf.List, "list", false, "List all tasks ")
	flag.StringVar(&cf.Add, "add", "", "Add a new task: -add \"Deploy code changes\" \"2024-11-12 22:57\"")
	flag.StringVar(&cf.Edit, "edit", "", "Edit a task: -edit \"5:Deploy code changes\"")
	flag.IntVar(&cf.Delete, "delete", -1, "Delete a task: -delete index")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle a task/mark as completed : -toggle index")
	flag.BoolVar(&cf.Help, "help", false, "Show all the available flags and their description")	
	flag.BoolVar(&cf.Refresh, "refresh", false, "Clear all all tasks: -refresh(completed, pending, all tasks)")
	
	flag.Parse()

	return &cf
}


func (cf *CmdFlags) Execute(todos *Todos){
	switch{
		case cf.List:
			todos.print()
		case cf.Add != "":
			if len(flag.Args()) > 0 {
				args := flag.Args()
				dateTime := args[0]

				if err := validateDateTime(dateTime); err != nil {
					fmt.Println(err)
					os.Exit(1)
				}

				layout := "2006-01-02 15:04"
			
				parsedTime, err := time.Parse(layout, dateTime)
				if err != nil {

					fmt.Println("Error parsing date:", err)
					os.Exit(1)

				}else if parsedTime.Before(time.Now()) {

					fmt.Println("Error: Date must be in the future")
					os.Exit(1)

				}else {
					todos.add(cf.Add, parsedTime)

				}

			}else{
			todos.add(cf.Add)}


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
			if err := todos.delete(cf.Delete); err != nil {
				fmt.Println("error : index out of range")
				os.Exit(1)
			}
		case cf.Toggle != -1:
			if err := todos.toggle(cf.Toggle); err != nil {
				fmt.Println("error : index out of range")
				os.Exit(1)
			}
		case cf.Help:
			fmt.Println("Available commands:")
			flag.PrintDefaults()
			os.Exit(0)
		case cf.Refresh:
			if len(flag.Args()) > 0 {
				switch{
					case flag.Args()[0] == "completed":
						todos.refresh("completed")
					case flag.Args()[0] == "pending":
						todos.refresh("pending")
					case flag.Args()[0] == "all":
						todos.refresh("all")
					default:
						fmt.Printf("invalid arguement %v: options(completed, pending ,all)",flag.Args()[0])

				}
			}
		default:
			fmt.Println("Invalid command")

		}

}
