package custom_type

import "fmt"

type CustomString string

func (text CustomString) Log() {
	fmt.Println(text)
}
