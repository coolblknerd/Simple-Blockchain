/* Blockchain: a digital ledger of records in smaller sets called blocks. These blocks are linked with on another through a cryptographic hashing. Each block contains a hash pointing to a previous block.

The blockchain is useful for cyptocurrencies because of its decentralize nature, meaning the data stored isn't in a single location, but is accessible by everyone and at same time, immutable by anyone.
*/

package main

import (
	"crypto/sha256"
	"encoding/hex"
	"encoding/json"
	"log"
	"net/http"
	"time"
)

/*
The hashes are used to identify and keep the blocks in the right order.
*/
type Block struct {
	Pos       int
	Data      BookCheckout
	Timestamp string
	Hash      string
	PrevHash  string
}

type BookCheckout struct {
	BookID       string `json:"book_id"`
	User         string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis    bool   `json:"is_genesis"`
}

type Book struct {
	ID          string `json:"id"`
	Title       string `json:"title"`
	Author      string `json:"author"`
	PublishDate string `json:"publsh_date"`
	ISBN        string `json:"isbn"`
}

type Blockchain struct {
	blocks []*Block
}

// This is a global variable that'll return the mutated Blockchain struct
var Blockchain *Blockchain

func (bc *Blockchain) AddBlock(data BookCheckout) {
	prevBlock := bc.blocks[len(bc.blocks)-1]
	block := CreateBlock(prevBlock, data)

	// validate integrity of blocks
	if validBlock(block, prevBlock) {
		bc.blocks = append(bc.blocks, block)
	}
}

func (b *Block) generateHash() {
	bytes, _ := json.Marshal(b.Data)
	data := string(b.Pos) + b.Timestamp + string(bytes) + b.PrevHash
	hash := sha256.New()
	hash.Write([]byte(data))
	b.Hash = hex.EncodeToString(hash.Sum(nil))
}

/*
This method will take a previous block and create a new block based on the info from the previous block.
*/

func CreateBlock(prevBlock *Block, checkoutItem BookCheckout) *Block {
	block := &Block{}
	block.Pos = prevBlock.Pos + 1
	block.Timestamp = time.Now().String()
	block.Data = checkoutItem
	block.PrevHash = prevBlock.Hash
	block.generateHash()

	return block
}

func GenesisBlock() *Block {
	return CreateBlock(&Block{}, BookCheckout{IsGenesis: true})
}

func NewBlockchain() *Blockchain {
	return &Blockchain{[]*Block{GenesisBlock()}}
}

func validBlock(block, prevBlock *Block) bool {
	if prevBlock.Hash != block.PrevHash {
		return false
	}

	if !block.validateHash(block.Hash) {
		return false
	}

	if prevBlock.Pos+1 != block.Pos {
		return false
	}

	return true
}

func (b *Block) validateHash(hash string) bool {
	b.generateHash()
	if b.Hash != hash {
		return false
	}

	return true
}

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/", getBlockchain).Methods("GET")
	r.HandleFunc("/", writeBlock).Methods("POST")
	r.HandleFunc("/new", newBook).Methods("POST")

	log.Println("Listening on port 3000")

	log.Fatal(http.ListenAndServe(":3000", r))
}
