package iteration

// Repeat a character for n number of times and return the repeated characters.
func Repeat(character string,  repeatCount int) string {
	var word string

	for i := 0; i < repeatCount; i++ {
		word += character
	}
	return word
}
