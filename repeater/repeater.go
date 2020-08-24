package repeater

func Repeater(value string) (repeated string) {
	for i := 0; i < 5; i++ {
		repeated += value
	}
	return
}
