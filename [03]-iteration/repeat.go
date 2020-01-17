package iteration

import "errors"

// Repeat will concatenate a passed in string to itself {x} number of times
func Repeat(character string, repeatCount int) (string, error) {
	if repeatCount < 0 {
		return "", errors.New("you can't repeat a character negative times. Duh")
	}
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character
	}
	return repeated, nil
}
