package paranoid

type fn func()

func PanicFunc(err error, f fn) {
	if err != nil {
		f()
		panic(err)
	}
}
