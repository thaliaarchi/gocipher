package mod

// Mod computes the modulo x mod n in the interval [0, n). The result
// takes the sign of the divisor n whereas % takes the sign of the
// dividend x.
func Mod(x, n int) int {
	return (x%n + n) % n
}

// ModInt8 computes the modulo x mod n in the interval [0, n). The
// result takes the sign of the divisor n whereas % takes the sign of
// the dividend x.
func ModInt8(x, n int8) int8 {
	return (x%n + n) % n
}

// ModRune computes the modulo x mod n in the interval [0, n) with rune
// operands. The result takes the sign of the divisor n whereas % takes
// the sign of the dividend x.
func ModRune(x, n rune) rune {
	return (x%n + n) % n
}

// Inverse computes the multiplicative inverse of x (mod n). An inverse
// only exists when gcd(x, n) == 1.
func Inverse(x, n int) (inv int, ok bool) {
	t, t1 := 0, 1
	r, r1 := n, Mod(x, n)
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

// Exp computes the modular exponation x ** y (mod n).
func Exp(x, y, n int) int {
	if y < 1 {
		return 1
	}
	x = Mod(x, n)
	e := 1
	for {
		if y&1 == 1 {
			e = (e * x) % n
		}
		if y == 0 {
			return e
		}
		x = (x * x) % n
		y >>= 1
	}
}
