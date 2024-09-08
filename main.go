package main

import (
	"log"
)

func main() {
	file := NewStorage[Todos]("data.json")
	todos := Todos{}
	if err := file.Load(&todos); err != nil {
		log.Fatal("error loading the data from file", err)
	}
	cmdFlags := NewCmdFlags()
	cmdFlags.Execute(&todos)
	if err := file.Save(todos); err != nil {
		log.Fatal("error saving the file", err)
	}
}
