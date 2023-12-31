package iteration

const repeatCount = 5

func Repeat(character string)  (word string) {
	// var word string

	for i := 0; i < repeatCount; i++ {
		word += character
	}
	return
}