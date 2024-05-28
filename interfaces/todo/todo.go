package todo

import (
	"encoding/json"
	"errors"
	"fmt"
	"os"
	"strings"
)

type Todo struct {
	Text string `json:"text"`
}

func (note Todo) Display() {
	fmt.Printf("Your to-do has the following text:\n\n%v\n\n", note.Text)
}

func (note Todo) Save() error {
	fileName := strings.ReplaceAll(note.Text, " ", "_")
	fileName = strings.ToLower(fileName) + ".json"

	json, err := json.Marshal(note)

	if err != nil {
		return err
	}

	return os.WriteFile(fileName, json, 0644)
}

func New(content string) (Todo, error) {
	if content == "" {
		return Todo{}, errors.New("Invalid input.")
	}

	return Todo{
		Text: content,
	}, nil
}
