package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
)

func main() {
	title, content := getNoteData()
	userNote, err := note.New(title, content)
	if err != nil {
		fmt.Println(err)
		return
	}

	userNote.Display()
	err = userNote.Save()
	if err != nil {
		fmt.Println("Error saving the note")
		return
	}
	fmt.Println("Note saved successfully")
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")
	content := getUserInput("Note content: ")
	return title, content
}

func getUserInput(prompt string) string {
	fmt.Print(prompt)
	var input string
	// Scanln allow user to input whitespaces such as \n
	// Scan does not take multiple words!
	// fmt.Scanln(&input)

	// to read longer text input:
	reader := bufio.NewReader(os.Stdin)
	// reader reads until user inputs the byte in delim parameter
	input, err := reader.ReadString('\n')
	if err != nil {
		return ""
	}
	// remove last whitespace from input
	input = strings.TrimSuffix(input, "\n")

	return input
}
