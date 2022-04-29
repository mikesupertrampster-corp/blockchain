package blockchain

import (
	"bytes"
	"crypto/sha256"
	"fmt"
	"math"
	"math/big"
)

const targetBits = 12

type Proof struct {
	block  *Block
	target *big.Int
}

func NewProof(b *Block) *Proof {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetBits))

	proof := &Proof{b, target}
	return proof
}

func (pow *Proof) InitNonce(nonce int) []byte {
	data := bytes.Join(
		[][]byte{
			pow.block.PreviousHash,
			pow.block.Data,
			ToHex(pow.block.Timestamp),
			ToHex(int64(targetBits)),
			ToHex(int64(nonce)),
		},
		[]byte{},
	)

	return data
}

func (pow *Proof) Run() (int, []byte) {
	var hashInt big.Int
	var hash [32]byte
	nonce := 0

	fmt.Printf("Mining the block containing \"%s\"\n", pow.block.Data)
	for nonce < math.MaxInt64 {
		data := pow.InitNonce(nonce)
		hash = sha256.Sum256(data)
		fmt.Printf("\r%x", hash)
		hashInt.SetBytes(hash[:])

		if hashInt.Cmp(pow.target) == -1 {
			break
		} else {
			nonce++
		}
	}
	fmt.Println()

	return nonce, hash[:]
}

func (pow Proof) Validate() bool {
	var initHash big.Int
	data := pow.InitNonce(pow.block.Nonce)

	hash := sha256.Sum256(data)
	initHash.SetBytes(hash[:])

	return initHash.Cmp(pow.target) == -1
}
