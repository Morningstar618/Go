package str

import "fmt"

type Str string

func (text Str) Log() {
	fmt.Println(text)
}
