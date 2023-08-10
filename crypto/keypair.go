package crypto

import (
	"crypto/ecdsa"
	"crypto/elliptic"
	"crypto/rand"
	"crypto/sha256"
	"math/big"

	"github.com/tahaontech/complete_blockchain/types"
)

type PrivateKey struct {
	key *ecdsa.PrivateKey
}

func (p PrivateKey) Sign(b []byte) (*Signature, error) {
	r, s, err := ecdsa.Sign(rand.Reader, p.key, b)
	if err != nil {
		return nil, err
	}

	return &Signature{r, s}, nil
}

func GeneratePrivateKey() PrivateKey {
	key, err := ecdsa.GenerateKey(elliptic.P256(), rand.Reader)
	if err != nil {
		panic(err)
	}

	return PrivateKey{
		key: key,
	}
}

func (pk *PrivateKey) PublicKey() PublicKey {
	return PublicKey{
		Key: &pk.key.PublicKey,
	}
}

type PublicKey struct {
	Key *ecdsa.PublicKey
}

func (k PublicKey) ToSlice() []byte {
	return elliptic.MarshalCompressed(k.Key, k.Key.X, k.Key.Y)
}

func (k *PublicKey) Address() types.Address {
	h := sha256.Sum256(k.ToSlice())

	return types.NewAddressFromBytes(h[len(h)-20:])
}

type Signature struct {
	R, S *big.Int
}

func (sig *Signature) Verify(pubKey PublicKey, data []byte) bool {
	return ecdsa.Verify(pubKey.Key, data, sig.R, sig.S)
}
