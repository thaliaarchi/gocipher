package rsa

import (
	"math/big"
	"testing"
)

func TestDecryptInt(t *testing.T) {
	e := big.NewInt(17)
	p, _ := new(big.Int).SetString("78032062222085722974121768604305613921737282772729", 10)
	q, _ := new(big.Int).SetString("95310075487163797179490457039169594160085543772343", 10)
	c, _ := new(big.Int).SetString("3780959489067185927636742897782124526407151622687008317493719218153002463614275880066047374967868051", 10)

	k := NewPrivateKey(p, q, e)
	m := k.DecryptInt(c)

	m0, _ := new(big.Int).SetString("4170008979850082696568008472738312007384078300827371728414", 10)
	if m.Cmp(m0) != 0 {
		t.Errorf("decrypted to %s, want %s", m, m0)
	}
}
