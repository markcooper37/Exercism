package diffiehellman

import (
	"crypto/rand"
	"math/big"
)

// Diffie-Hellman-Merkle key exchange
// Private keys should be generated randomly.

func PrivateKey(p *big.Int) *big.Int {
	var z big.Int
	key, err := rand.Int(rand.Reader, z.Sub(p, big.NewInt(2)))
	if err != nil {
		panic("something went wrong generating private key")
	}
	return z.Add(key, big.NewInt(2))
}

func PublicKey(private, p *big.Int, g int64) *big.Int {
	var z big.Int
	return z.Exp(big.NewInt(g), private, p)
}

func NewPair(p *big.Int, g int64) (*big.Int, *big.Int) {
	private := PrivateKey(p)
	public := PublicKey(private, p, g)
	return private, public
}

func SecretKey(private1, public2, p *big.Int) *big.Int {
	var z big.Int
	return z.Exp(public2, private1, p)
}
