package main

import (
	"fmt"
)

func main() {
	todos := Todos{}
	todos.add("Add milk")
	todos.add("A test item")
	todos.print()
	if err := todos.delete(0); err != nil {
		fmt.Println(err)
	}
	if err := todos.toggle(5); err != nil {
		fmt.Println(err)
	}
	todos.print()
}
