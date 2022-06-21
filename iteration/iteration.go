package iteration

func Repeat(char string, timesRepeated int) string {
	var repeated string
	for i := 0; i < timesRepeated; i++ {
		repeated += char
	}
	return repeated
}
