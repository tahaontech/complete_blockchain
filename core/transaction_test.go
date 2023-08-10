package core

import (
	"testing"

	"github.com/stretchr/testify/assert"
	"github.com/tahaontech/complete_blockchain/crypto"
)

func TestSignTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.Nil(t, tx.Sign(privKey))
	assert.NotNil(t, tx.Signature)
}

func TestVerifyTransaction(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}

	assert.NotNil(t, tx.Verify())

	assert.Nil(t, tx.Sign(privKey))
	assert.Nil(t, tx.Verify())

	anotherKey := crypto.GeneratePrivateKey()
	tx.From = anotherKey.PublicKey()

	assert.NotNil(t, tx.Verify())
}

func randomTxWithSignature(t *testing.T) *Transaction {
	privKey := crypto.GeneratePrivateKey()
	tx := &Transaction{
		Data: []byte("foo"),
	}
	assert.Nil(t, tx.Sign(privKey))
	return tx
}
