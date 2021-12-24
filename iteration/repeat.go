package iteration

const repeatCount = 5

func Repeat(c string) (repeated string) {
	for i := 0; i < repeatCount; i++ {
		repeated += c
	}
	return
}
