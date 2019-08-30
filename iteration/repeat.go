package iteration

const repeatCount = 5

// Repeat will concatenate a passed in string to itself {x} number of times
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated
}
