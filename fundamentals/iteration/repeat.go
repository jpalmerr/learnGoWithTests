package iteration

const repeatCount = 5

// Repeat returns character repeated 5 times.
func Repeat(character string) string {
	var repeated string
	for i := 0; i < repeatCount; i++ {
		repeated += character // repeated = repeated + character
	}
	return repeated
}

func RepeatN(character string, repeats int) string {
	var repeated string
	for i := 0; i < repeats; i++ {
		repeated += character
	}
	return repeated
}
