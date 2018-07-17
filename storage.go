package main

import "fmt"
import "os"

type storage struct {
	location string
}

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func (s *storage) save_data(text string) {

	// Open file for writing
	file, err := os.Create(s.location)
	check(err)

	defer file.Close()

	bytes, err := file.WriteString(text)
	check(err)

	fmt.Println("Bytes written: %d", bytes)
}
