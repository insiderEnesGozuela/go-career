package iteration

import "testing"

func TestRepeat(t *testing.T) {
	got := Repeat("a")
	expected := "aaaa"

	if got != expected {
		t.Errorf("got %q expected %q", got, expected)
	}
}

func BenchmarkRepeat(b *testing.B) {
	// loop içindeki bloğu art arda çağırır ve ortalama çalışma hızını hesaplar
	for b.Loop() {
		Repeat("a")
	}
}