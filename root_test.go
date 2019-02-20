package root

import "testing"

var mulRootTests = []struct {
	x        float64
	p        float64
	expected float64
}{
	{5, 2, 2.23606797749979},
	{29, 3, 3.072316825685847},
	{27, 3, 3.0},
}

func TestRoot(t *testing.T) {
	for _, mt := range mulRootTests {
		r := root(mt.x, mt.p)
		if r != mt.expected {
			t.Errorf("\nexpected %f got %f", mt.expected, r)
		}
	}
}
func BenchmarkRoot(b *testing.B) {
	for n := 0; n < b.N; n++ {
		root(29, 3)
	}
}
