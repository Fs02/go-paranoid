package paranoid

import "fmt"

type fn func()

func PanicFunc(err error, message string, f fn) {
	if err != nil {
		f()
		panic(fmt.Errorf("%s: %v", message, err))
	}
}
