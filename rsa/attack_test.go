package rsa

import (
	"math/big"
	"testing"
)

func TestAttackShortPlaintext(t *testing.T) {
	// m := big.NewInt(12345678901) // product of primes 857 and 14405693
	// m := big.NewInt(0x65822107fcfd52) // 56 bits like DES key
	m := big.NewInt(1234567890) // quick
	n, _ := new(big.Int).SetString("7437241740806052646950317400323393037944894112922390037623344174206297959788872391199205745384834047", 10)
	e := big.NewInt(17)
	k := PublicKey{n, e}
	c := k.EncryptInt(m)
	m1 := k.AttackShortPlaintext(c, 32)
	if m.Cmp(m1) != 0 {
		t.Errorf("got %s, want %s", m1, m)
	}
}

func TestFactorClosePrimes(t *testing.T) {
	n := big.NewInt(295927)
	p1, q1 := big.NewInt(541), big.NewInt(547)
	p, q, ok := FactorClosePrimes(n, big.NewInt(100))
	if !ok {
		t.Fatal("factoring failed")
	}
	if p.Cmp(p1) != 0 || q.Cmp(q1) != 0 {
		t.Errorf("got factors (%s, %s), want (%s, %s)", p, q, p1, q1)
	}
}
