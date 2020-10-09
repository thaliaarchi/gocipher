package mod

// Mod computes a modulo n in the interval [0, n). The result takes the
// sign of the divisor n whereas % takes the sign of the dividend a.
func Mod(a, n int) int {
	return (a%n + n) % n
}

// ModRune computes a modulo n in the interval [0, n) with rune
// operands. The result takes the sign of the divisor n whereas % takes
// the sign of the dividend a.
func ModRune(a, n rune) rune {
	return (a%n + n) % n
}

// Inverse computes the multiplicative inverse of a (mod n). An inverse
// only exists when gcd(a, n) == 1.
func Inverse(a, n int) (inv int, ok bool) {
	t, t1 := 0, 1
	r, r1 := n, Mod(a, n)
	for r1 != 0 {
		quotient := r / r1
		t, t1 = t1, t-quotient*t1
		r, r1 = r1, r-quotient*r1
	}
	if r != 1 {
		return 0, false
	}
	if t < 0 {
		t += n
	}
	return t, true
}
