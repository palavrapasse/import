package parser

func panicOnErrors(err []error) {
	if len(err) != 0 {
		panic(err)
	}
}

func panicOnError(err error) {
	if err != nil {
		panicOnErrors([]error{err})
	}
}
