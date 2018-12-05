package core

//区块链表
type BlockChain struct {
	Blocks []*Block
}
//添加一个区块
func (bc *BlockChain) AddBlock(data string) {
	preBlock := bc.Blocks[len(bc.Blocks)-1]
	newBlock := NewBlock(data, preBlock.Hash)
	bc.Blocks = append(bc.Blocks, newBlock)
}

func NewBlockChain() *BlockChain {
	return &BlockChain{Blocks: []*Block{NewGenesisBlock()}}
}
