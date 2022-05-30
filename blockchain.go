package main

import (
        "crypto/sha256"
        "encoding/json"
        "fmt"
        "strconv"
        "strings"
        "time"
)

type Block struct {
	data         map[string]interface{}
	hash         string
	previousHash string
	timestamp    time.Time
	pow          int
}

type Blockchain struct {
	genesisBlock Block // Primeiro bloco adicionado ao blockchain
	chain        []Block
	difficulty   int // Esforço mínimo que os mineiros devem realizar para extrair um bloco
}

func (b Block) calculateHash() string {
	data, _ := json.Marshal(b.data) // Converter os dados do bloco para o formato JSON
	blockData := b.previousHash + string(data) + b.timestamp.String() + strconv.Itoa(b.pow) // Concatenar o hash anterior do bloco, dados, carimbo de data / hora e prova de trabalho (PoW)
	blockHash := sha256.Sum256([]byte(blockData)) // Hashed a concatenação anterior com o algoritmo SHA256
	return fmt.Sprintf("%x", blockHash) // Retornar o resultado de hash na base 16, com letras minúsculas para AF
}

func (b *Block) mine(difficulty int) {
	for !strings.HasPrefix(b.hash, strings.Repeat("0", difficulty)) {
					b.pow++
					b.hash = b.calculateHash()
	}
}

func CreateBlockchain(difficulty int) Blockchain {
	genesisBlock := Block{
					hash:      "0",
					timestamp: time.Now(),
	}
	return Blockchain{
					genesisBlock,
					[]Block{genesisBlock},
					difficulty,
	}
}

// Incluir novos blocos em um blockchain.
func (b *Blockchain) addBlock(from, to string, amount float64) {
	blockData := map[string]interface{}{
					"from":   from,
					"to":     to,
					"amount": amount,
	}
	lastBlock := b.chain[len(b.chain)-1]
	newBlock := Block{
					data:         blockData,
					previousHash: lastBlock.hash,
					timestamp:    time.Now(),
	}
	newBlock.mine(b.difficulty)
	b.chain = append(b.chain, newBlock)
}

func (b Blockchain) isValid() bool {
	for i := range b.chain[1:] {
					previousBlock := b.chain[i]
					currentBlock := b.chain[i+1]
					if currentBlock.hash != currentBlock.calculateHash() || currentBlock.previousHash != previousBlock.hash {
									return false
					}
	}
	return true
}

func main() {
	// Create a new blockchain instance with a mining difficulty of 2
	blockchain := CreateBlockchain(2)

	// Record transactions on the blockchain for Alice, Bob, and John
	blockchain.addBlock("Gustavo", "Vivian", 5)
	blockchain.addBlock("Amanda", "Vivian", 2)
	blockchain.addBlock("Gustavo", "Amanda", 5)

	// Check if the blockchain is valid; expecting true
	fmt.Println(blockchain)
	fmt.Println(blockchain.isValid())
}