package main

import (
	"fmt"
	"os"
	"strings"

	"golang.org/x/text/cases"
	"golang.org/x/text/language"
)

func main() {
	caser := cases.Title(language.English)
	dir := os.Args[1]

	splitString := strings.Split(dir, "-")
	title := caser.String(strings.Join(splitString[:], " "))

	if err := os.Mkdir(dir, os.ModePerm); err != nil {
		fmt.Print(err)
	}

	// fmt.Sprintf("%[1]v/%[1]v.md", dir)
	if err := os.WriteFile(dir+"/"+dir+".md", []byte("# "+title), os.ModePerm); err != nil {
		fmt.Print(err)
	}

	if err := os.WriteFile(dir+"/main.go", nil, os.ModePerm); err != nil {
		fmt.Print(err)
	}
}
