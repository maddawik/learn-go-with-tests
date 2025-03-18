package iteration

func Repeat(character string) string {
	var repeated string
	for range 5 {
		repeated += character
	}
	return repeated
}
