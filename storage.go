package main

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
)

type Storage[T any] struct {
	FileName string
}

func NewStorage[T any](fileName string) *Storage[T] {
	return &Storage[T]{fileName}
}

func (s *Storage[T]) Save(data T) error {
	byteData, err := json.MarshalIndent(data, "", "  ")
	if err != nil {
		fmt.Println("error marshalling data:", err)
		return err
	}
	return os.WriteFile(s.FileName, byteData, 0644)
}

func (s *Storage[T]) Load(data *T) error {
	if err := s.createFileIfNotExist(); err != nil {
		fmt.Println("error creating file:", err)
		return err
	}
	byteData, err := os.ReadFile(s.FileName)
	if err != nil {
		fmt.Println("error reading file:", err)
		return err
	}
	err = json.Unmarshal(byteData, data)
	if err != nil {
		fmt.Println("error unmarshalling data:", err)
		return err
	}
	return nil
}

func (s *Storage[T]) doesFileExist() bool {
	if _, err := os.Stat(s.FileName); errors.Is(err, os.ErrNotExist) {
		return false
	}
	return true
}

func (s *Storage[T]) createFileIfNotExist() error {
	if res := s.doesFileExist(); res {
		return nil
	}
	return os.WriteFile(s.FileName, []byte("[]"), 0644)
}
