package iteration

const repeatedCount = 5

func Repeat(c string) string {
	var repeated string
	for i := 0; i < repeatedCount; i++ {
		repeated += c
	}
	return repeated
}
