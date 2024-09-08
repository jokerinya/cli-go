package main

import (
	"flag"
	"fmt"
	"log"
	"strconv"
	"strings"
)

type CmdFlags struct {
	Add    string
	Del    int
	Edit   string
	Toggle int
	List   bool
}

func NewCmdFlags() *CmdFlags {
	cf := CmdFlags{}

	flag.StringVar(&cf.Add, "add", "", "Add a new entry")
	flag.StringVar(&cf.Edit, "edit", "", "Edit an entry")
	flag.IntVar(&cf.Del, "del", -1, "Delete a entry")
	flag.IntVar(&cf.Toggle, "toggle", -1, "Toggle an entry")
	flag.BoolVar(&cf.List, "list", false, "List all entries")

	flag.Parse()
	return &cf
}

func (cf *CmdFlags) Execute(todos *Todos) {
	switch {
	case cf.List:
		todos.print()
	case cf.Add != "":
		todos.add(cf.Add)
	case cf.Edit != "":
		parts := strings.SplitN(cf.Edit, ":", 2)
		if len(parts) != 2 {
			log.Fatal("Invalid edit format, enter like id:new_title")
		}
		index, err := strconv.Atoi(parts[0])
		if err != nil {
			log.Fatal("Invalid index for formating")
		}

		if err := todos.edit(index, parts[1]); err != nil {
			log.Fatal(err)
		}
	case cf.Toggle != -1:
		if err := todos.toggle(cf.Toggle); err != nil {
			log.Fatal(err)
		}
	case cf.Del != -1:
		if err := todos.delete(cf.Del); err != nil {
			log.Fatal(err)
		}
	default:
		fmt.Println("Invalid command")
	}
}
