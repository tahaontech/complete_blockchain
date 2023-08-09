package crypto

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestKeyPairs(t *testing.T) {
	privKey := GeneratePrivateKey()
	pubKey := privKey.PublicKey()
	address := pubKey.Address()

	assert.False(t, address.IsZero())

	msg := "Hello Worlds"
	sig, err := privKey.Sign([]byte(msg))
	assert.Nil(t, err)

	assert.True(t, sig.Verify(pubKey, []byte(msg)))
}
