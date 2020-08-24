package repeater

import "testing"

func TestRepeater(t *testing.T) {
	repeated := Repeater("a")
	expected := "aaaaa"
	if repeated != expected {
		t.Errorf("Expected '%s', but got '%s'", expected, repeated)
	}
}

func BenchmarkRepeater(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeater("a")
	}
}
