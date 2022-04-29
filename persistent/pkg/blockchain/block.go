package blockchain

import (
	"bytes"
	"encoding/gob"
	simple "github.com/mikesupertrampster-corp/blockchain/simple/pkg/blockchain"
)

func Serialize(b *simple.Block) ([]byte, error) {
	var results bytes.Buffer
	encoder := gob.NewEncoder(&results)

	err := encoder.Encode(b)
	if err != nil {
		return nil, err
	}

	return results.Bytes(), nil
}

func DeserializeBlock(d []byte) (*simple.Block, error) {
	var block simple.Block

	decoder := gob.NewDecoder(bytes.NewReader(d))
	err := decoder.Decode(&block)
	if err != nil {
		return nil, err
	}

	return &block, nil
}
