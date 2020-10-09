package mod

import "testing"

func TestInverse(t *testing.T) {
	for i := -50; i < 50; i++ {
		inv1, ok1 := Inverse(i, 26)
		inv2, ok2 := inverseNaive(i, 26)
		if ok1 != ok2 || inv1 != inv2 {
			t.Errorf("%d: inverse is %d %t, but naive is %d %t", i, inv1, ok1, inv2, ok2)
		}
	}
}

func inverseNaive(a, n int) (int, bool) {
	for i := 1; i < n; i++ {
		if Mod(a*i, n) == 1 {
			return i, true
		}
	}
	return 0, false
}

func TestExp(t *testing.T) {
	for _, test := range []struct {
		x, y, n, e int
	}{
		{3, 75, 103, 24},
	} {
		e := Exp(test.x, test.y, test.n)
		if e != test.e {
			t.Errorf("Exp(%d, %d, %d) = %d, want %d", test.x, test.y, test.n, e, test.e)
		}
	}
}
