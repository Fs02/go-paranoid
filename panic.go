package paranoid

import "fmt"

func Panic(err error, message string) {
	if err != nil {
		panic(fmt.Errorf("%s: %v", message, err))
	}
}
