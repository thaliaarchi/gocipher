package gocipher

// mod computes the modulo a mod n in the interval [0, n). The result
// takes the sign of the divisor whereas % takes the sign of the
// dividend.
func mod(a, n int) int {
	return (a%n + n) % n
}

// mod computes the modulo a mod n in the interval [0, n) for runes. The
// result takes the sign of the divisor whereas % takes the sign of the
// dividend.
func modRune(a, n rune) rune {
	return (a%n + n) % n
}

// modInverse computes the multiplicative inverse of a modulo n. An
// inverse only exists when gcd(a, n) == 1.
func modInverse(a, n int) (inv int, ok bool) {
	t, t1 := 0, 1
	r, r1 := n, mod(a, n)
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
