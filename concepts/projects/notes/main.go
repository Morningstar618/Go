package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"

	"example.com/notes/note"
	"example.com/notes/todo"
)

type saver interface {
	Save() error
}

type outputable interface {
	saver
	Display()
}

func main() {
	// Any type testing - these should work
	printSomething(1)
	printSomething(1.5)
	printSomething("One")

	title, content := getNoteData()
	todoText := getUserInput("Todo Text: ")

	todo, err := todo.New(todoText)

	printSomething(todo) // should not work

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
		return
	}

	outputData(userNote)
}

func printSomething(value interface{}) {

	intVal, ok := value.(int)

	if ok {
		fmt.Println(intVal)
		return
	}

	floatVal, ok := value.(float64)

	if ok {
		fmt.Println(floatVal)
		return
	}

	stringVal, ok := value.(string)

	if ok {
		fmt.Println(stringVal)
		return
	}

	// switch value.(type) {
	// case int:
	// 	fmt.Println("Integer: ", value)
	// case float64:
	// 	fmt.Println("Float64: ", value)
	// case string:
	// 	fmt.Println("String: ", value)
	// }
}

func outputData(data outputable) error {
	data.Display()
	err := saveData(data)

	if err != nil {
		return err
	}

	return nil
}

func saveData(data outputable) error {
	err := data.Save()

	if err != nil {
		fmt.Println("Saving the todo failed!")
		return err
	}

	fmt.Println("Saving the todo succeeded!")
	return nil
}

func getNoteData() (string, string) {
	title := getUserInput("Note title: ")

	content := getUserInput("Note content: ")

	return title, content
}

func getUserInput(prompt string) string {
	fmt.Printf("%v ", prompt)
	var text string

	reader := bufio.NewReader(os.Stdin)
	text, err := reader.ReadString('\n')

	if err != nil {
		return ""
	}

	text = strings.TrimSuffix(text, "\n")
	text = strings.TrimSuffix(text, "\r")

	return text
}
