package rsa

import (
	"math/big"
)

// AttackShortPlaintext exploits short plaintexts to compute the key.
func (k *PublicKey) AttackShortPlaintext(c *big.Int, bits int) (m *big.Int) {
	bm := make(bigMap)
	tmp := new(big.Int)

	// Find x and y such that c*(x^-e) = y^e (mod n)
	i := big.NewInt(1)
	var x, y *big.Int
	for i.BitLen() <= bits {
		ye := new(big.Int).Exp(i, k.E, k.N) // i^e      (mod n)
		tmp.ModInverse(ye, k.N)             // i^-e     (mod n)
		xe := new(big.Int).Mul(c, tmp)      // c*(i^-e) (mod n)
		xe.Mod(xe, k.N)

		// Check for match
		if v, ok := bm.GetOrPut(xe, i); ok {
			x, y = i, v
			break
		}
		if v, ok := bm.GetOrPut(ye, i); ok {
			x, y = v, i
			break
		}
		i = new(big.Int).Add(i, one)
	}

	// We now have c = (x*y)^e (mod n), thus m = x*y (mod n)
	xy := new(big.Int).Mul(x, y)
	xy.Mod(xy, k.N)
	return xy
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
