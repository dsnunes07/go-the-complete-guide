package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/note/note"
	"example.com/note/todo"
)

// declare Saver interface
type saver interface {
	Save() error
}

// type displayer interface {
// 	Display()
// }

// embedded interface
type outputtable interface {
	saver
	Display()
}

func main() {
	title, content := getNoteData()
	todoText := getUserInput("Todo text: ")

	todo, err := todo.New(todoText)
	if err != nil {
		fmt.Println(err)
		return
	}
	err = outputData(todo)
	if err != nil {
		return
	}

	userNote, err := note.New(title, content)

	if err != nil {
		fmt.Println(err)
		return
	}
	outputData(userNote)
	printSomething("abc")
	printSomething(todo)
	printSomething(1.5)
	result := add(1, 1)
	printSomething(result)
	resultString := add("Text", " other text")
	printSomething(resultString)

}

func outputData(data outputtable) error {
	data.Display()
	return saveData(data)

}

func add[T int | float64 | string](a, b T) T {
	return a + b
}

// function accept any value type
func printSomething(value interface{}) {
	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float: ", value)
	// case string:
	// 	fmt.Println("string: ", value)
	// default:
	// 	fmt.Println(value)
	// }
	typedVal, isInt := value.(int)
	if isInt {
		fmt.Printf("Integer: %v\n", typedVal)
	}
	floatVal, isFloat := value.(float64)
	if isFloat {
		fmt.Printf("float64: %v\n", floatVal)
	}
	stringVal, isString := value.(string)
	if isString {
		fmt.Printf("stringVal: %v\n", stringVal)
	}
}

func saveData(data saver) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving data failed.")
		return err
	}
	fmt.Println("Saving data succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title:")
	content := getUserInput("Note content:")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)

	reader := bufio.NewReader(os.Stdin)

	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
