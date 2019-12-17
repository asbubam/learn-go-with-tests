package iteration

func Repeat(character string, repeat_count int) string {
	var repeated string
	for i := 0; i < repeat_count; i++ {
		repeated += character
	}
	return repeated
}