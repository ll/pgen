package main

import (
	"fmt"
	"os"
	"strings"
)

const id = "ID"

func snake2Camel(word string) string {
	//ID should be always uppercas
	words := strings.Split(word, "_")
	for i, w := range words {
		if strings.ToLower(w) == strings.ToLower(id) {
			words[i] = strings.ToUpper(id)
		}
	}

	return strings.Replace(strings.Title(strings.Join(words, " ")), " ", "", -1)
}

func getTableName(word string) string {
	//ID should be always uppercas
	words := strings.Split(word, ".")
	if len(words) != 2 {
		fmt.Println("table name is wrong")
	}

	return snake2Camel(words[1])
}

func saveToFile(name string, b []byte) error {
	words := strings.Split(name, ".")
	f, err := os.OpenFile(words[len(words)-1]+".go", os.O_CREATE|os.O_WRONLY|os.O_TRUNC, 0644)
	if err != nil {
		return err
	}

	if _, err := f.Write(b); err != nil {
		return err
	}

	if err := f.Close(); err != nil {
		return err
	}
	return nil
}
