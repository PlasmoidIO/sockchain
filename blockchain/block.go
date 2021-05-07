package blockchain

type Block struct {
	Hash     []byte
	Data     []byte
	PrevHash []byte
	Nonce    int
	Proof    *ProofOfWork
}

type BlockChain struct {
	Blocks []*Block
}

func CreateBlock(data string, prevHash []byte) *Block {
	block := &Block{
		Hash:     []byte{},
		Data:     []byte(data),
		PrevHash: prevHash,
		Nonce:    0,
	}
	pow := NewProof(block)
	nonce, hash := pow.Run()

	block.Proof = pow
	block.Nonce = nonce
	block.Hash = hash

	return block
}

func (chain *BlockChain) AddBlock(data string) {
	prevBlock := chain.Blocks[len(chain.Blocks)-1]
	block := CreateBlock(data, prevBlock.Hash)
	chain.Blocks = append(chain.Blocks, block)
}

func InitBlockChain() *BlockChain {
	genesis := CreateBlock("Genesis", []byte{})
	return &BlockChain{[]*Block{genesis}}
}
