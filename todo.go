package main

import (
	"errors"
	"github.com/aquasecurity/table"
	"os"
	"reflect"
	"strconv"
	"time"
)

type Todo struct {
	Title       string     `json:"title"`
	Completed   bool       `json:"completed"`
	CreatedAt   time.Time  `json:"created_at"`
	CompletedAt *time.Time `json:"completed_at,omitempty"`
}

type Todos []Todo

func (t *Todos) add(title string) {
	todo := Todo{
		Title:       title,
		Completed:   false,
		CompletedAt: nil,
		CreatedAt:   time.Now(),
	}
	*t = append(*t, todo)
}

func (t *Todos) delete(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	*t = append((*t)[:index], (*t)[index+1:]...)
	return nil
}

func (t *Todos) toggle(index int) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	todo := (*t)[index]
	todo.Completed = !todo.Completed
	if todo.Completed {
		compAt := time.Now()
		todo.CompletedAt = &compAt
	}
	(*t)[index] = todo
	return nil
}

func (t *Todos) edit(index int, title string) error {
	if err := t.validateIndex(index); err != nil {
		return err
	}
	(*t)[index].Title = title
	return nil
}

func (t *Todos) validateIndex(index int) error {
	if index < 0 || index >= len(*t) {
		return errors.New("index out of range")
	}
	return nil
}

func (t *Todos) print() {
	tbl := table.New(os.Stdout)
	headers := []string{"ID"}
	headers = append(headers, (*t).getKeys()...)
	tbl.SetHeaders(headers...)
	for i, todo := range *t {
		row := []string{strconv.Itoa(i)}
		row = append(row, todo.getValues()...)
		tbl.AddRow(row...)
	}
	tbl.Render()
}

func (todo *Todo) getValues() []string {
	var values []string
	values = append(values, (*todo).Title)
	if todo.Completed {
		values = append(values, "DONE!")
	} else {
		values = append(values, "NOT YET")
	}
	values = append(values, (*todo).CreatedAt.Format("2006-01-02 15:04:05"))
	if todo.CompletedAt != nil {
		values = append(values, (*todo).CompletedAt.Format("2006-01-02 15:04:05"))
	} else {
		values = append(values, "")
	}
	return values
}

func (t *Todos) getKeys() []string {
	var keys []string
	typ := reflect.TypeOf(Todo{})
	for i := 0; i < typ.NumField(); i++ {
		keys = append(keys, typ.Field(i).Name)
	}
	return keys
}
