package rsa

import "math/big"

// A PublicKey represents the public part of an RSA key.
type PublicKey struct {
	N *big.Int // modulus
	E *big.Int // public exponent
}

// A PrivateKey represents an RSA key
type PrivateKey struct {
	PublicKey
	D    *big.Int // private exponent
	P, Q *big.Int // prime factors of N
}

var one = big.NewInt(1)

// NewPrivateKey creates a private key given two prime factors p and q
// and public exponent e.
func NewPrivateKey(p, q, e *big.Int) *PrivateKey {
	t1, t2 := new(big.Int), new(big.Int)
	p1q1 := new(big.Int).Mul(t1.Sub(p, one), t2.Sub(q, one)) // (p-1)*(q-1)
	n := t1.Mul(p, q)
	d := t2.ModInverse(e, p1q1) // where d*e = 1 (mod (p-1)*(q-1))
	return &PrivateKey{PublicKey{n, e}, d, p, q}
}

func (k *PublicKey) Encrypt(msg []byte) []byte {
	m := new(big.Int).SetBytes(msg)
	return k.EncryptInt(m).Bytes()
}

func (k *PublicKey) EncryptInt(m *big.Int) *big.Int {
	return new(big.Int).Exp(m, k.E, k.N)
}

func (k *PrivateKey) Decrypt(ciphertext []byte) []byte {
	c := new(big.Int).SetBytes(ciphertext)
	return k.DecryptInt(c).Bytes()
}

func (k *PrivateKey) DecryptInt(c *big.Int) *big.Int {
	return new(big.Int).Exp(c, k.D, k.N)
}
