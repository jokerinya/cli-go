package main

import (
	"fmt"
	"log"
)

func main() {
	file := NewStorage[Todos]("data.json")
	todos := Todos{}
	err := file.Load(&todos)
	if err != nil {
		log.Fatal("error loading the data from file", err)
	}
	todos.add("Add milk 2")
	todos.add("Add milk 4")
	if err := todos.toggle(1); err != nil {
		fmt.Println(err)
	}
	todos.print()
	if err := todos.delete(0); err != nil {
		fmt.Println(err)
	}
	todos.print()
	if err := file.Save(todos); err != nil {
		fmt.Println(err)
	}
}
