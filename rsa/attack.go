package rsa

import "math/big"

// AttackShortPlaintext exploits short plaintexts to compute the key.
func (k *PublicKey) AttackShortPlaintext(c *big.Int, bits int) (m *big.Int) {
	// Described in Introduction to Cryptography with Coding Theory by
	// Trappe and Washington, section 6.2.2.

	bm := make(bigMap)
	tmp := new(big.Int)

	// Find x and y such that c*(x^-e) = y^e (mod n)
	var x, y *big.Int
	for i := one; i.BitLen() <= bits; i = new(big.Int).Add(i, one) {
		yi := new(big.Int).Exp(i, k.E, k.N) // i^e      (mod n)
		tmp.ModInverse(yi, k.N)             // i^-e     (mod n)
		xi := new(big.Int).Mul(c, tmp)      // c*(i^-e) (mod n)
		xi.Mod(xi, k.N)

		// Check for match
		if v, ok := bm.GetOrPut(xi, i); ok {
			x, y = i, v
			break
		}
		if v, ok := bm.GetOrPut(yi, i); ok {
			x, y = v, i
			break
		}
	}

	// c*(x^-e) = y^e (mod n)
	// c = (x*y)^e    (mod n)
	// m = x*y        (mod n)
	m = tmp.Mod(tmp.Mul(x, y), k.N)
	return m
}

func FactorClosePrimes(n, diff *big.Int) (p, q *big.Int, ok bool) {
	ni2, sqrt, t := new(big.Int), new(big.Int), new(big.Int)
	for i := new(big.Int); i.Cmp(diff) < 1; i.Add(i, one) {
		ni2.Add(n, t.Mul(i, i)) // n + i^2
		sqrt.Sqrt(ni2)
		if t.Mul(sqrt, sqrt).Cmp(ni2) == 0 { // n + i^2 is a perfect square
			p = new(big.Int).Sub(sqrt, i)
			q = new(big.Int).Add(sqrt, i)
			return p, q, true
		}
	}
	return nil, nil, false
}

// bigMap is a hash table for *big.Int keys. Keys are not copied and
// must not be modified after insertion.
type bigMap map[uint64][]mapPair
type mapPair struct{ K, V *big.Int }

// GetOrPut gets the value at the key, if it exists, or otherwise puts a
// value at the key.
func (m bigMap) GetOrPut(key, val *big.Int) (*big.Int, bool) {
	hash := key.Uint64()
	bucket := m[hash]
	for _, pair := range bucket {
		if pair.K.Cmp(key) == 0 {
			return pair.V, true
		}
	}
	m[hash] = append(bucket, mapPair{key, val})
	return nil, false
}
