package core

import (
	"bytes"
	"math/big"
)

const targetbits = 24

type ProofOfWork struct {
	block  *Block
	target *big.Int
}

func NewProofOfWork(b *Block) *ProofOfWork {
	target := big.NewInt(1)
	target.Lsh(target, uint(256-targetbits))
	return &ProofOfWork{
		block:  b,
		target: target,
	}
}

func (pow *ProofOfWork) prepareData(nonce int) []byte {
	return bytes.Join([][]byte{
		pow.block.PreBlockHash,
		pow.block.Hash,
		IntToHex(pow.block.Timestamp),
		IntToHex(int64(targetbits)),
		IntToHex(int64(nonce)),
	}, []byte{})
}

func (pow *ProofOfWork)Run()  {
	
}
