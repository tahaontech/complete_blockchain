package core

import (
	"testing"
	"time"

	"github.com/stretchr/testify/assert"
	"github.com/tahaontech/complete_blockchain/crypto"
	"github.com/tahaontech/complete_blockchain/types"
)

func TestHashBlock(t *testing.T) {
	b := randomBlock(0, types.RandomHash())
	h := b.Hash(BlockHasher{})

	assert.False(t, h.IsZero())
}

func TestSignBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.RandomHash())

	assert.Nil(t, b.Sign(privKey))
	assert.NotNil(t, b.Signature)
}

func TestVerifyBlock(t *testing.T) {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(0, types.RandomHash())

	assert.Nil(t, b.Sign(privKey))
	assert.Nil(t, b.Verify())

	otherPrivKey := crypto.GeneratePrivateKey()
	b.Validator = otherPrivKey.PublicKey()
	assert.NotNil(t, b.Verify())

	b.Height = 20
	assert.NotNil(t, b.Verify())
}

func randomBlock(height uint32, prevBlockhash types.Hash) *Block {
	header := &Header{
		Version:       1,
		PrevBlockHash: prevBlockhash,
		Height:        height,
		TimeStamp:     time.Now().UnixNano(),
	}

	return NewBlock(header, []Transaction{})
}

func randomBlockWithSignature(t *testing.T, height uint32, prevBlockHash types.Hash) *Block {
	privKey := crypto.GeneratePrivateKey()
	b := randomBlock(height, prevBlockHash)
	tx := randomTxWithSignature(t)
	b.AddTransaction(tx)
	assert.Nil(t, b.Sign(privKey))

	return b
}
