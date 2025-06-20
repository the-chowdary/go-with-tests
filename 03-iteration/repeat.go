package iteration

func Repeat(c string, count int) string {
	if count <= 0 {
		return c
	}
	var repeated string
	for i := 0; i < count; i++ {
		repeated += c
	}
	return repeated
}

func main() {
	Repeat("b", 5)
}
