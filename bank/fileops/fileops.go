package fileops

import (
	"errors"
	"fmt"
	"os"
	"strconv"
)

// in go, only functions starting with capital letters are available in other packages!
func WriteFloatToFile(value float64, fileName string) {
	valueText := fmt.Sprint(value)
	os.WriteFile(fileName, []byte(valueText), 0644)
}

func GetFloatFromFile(fileName string, defaultValue float64) (float64, error) {
	data, err := os.ReadFile(fileName)
	if err != nil {
		return defaultValue, errors.New("failed to find file")
	}
	valueReadText := string(data)
	valueRead, err := strconv.ParseFloat(valueReadText, 64)
	if err != nil {
		return defaultValue, errors.New("failed to parse store value")
	}

	return valueRead, nil
}
