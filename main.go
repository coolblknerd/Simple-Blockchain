/* Blockchain: a digital ledger of records in smaller sets called blocks. These blocks are linked with on another through a cryptographic hashing. Each block contains a hash pointing to a previous block. 

The blockchain is useful for cyptocurrencies because of its decentralize nature, meaning the data stored isn't in a single location, but is accessible by everyone and at same time, immutable by anyone.
*/

package main

import "fmt"

type Block struct {
	Pos int
	Data BookCheckout
	Timestamp string
	Hash string
	PrevHash string
}

type BookCheckout struct {
	BookID string `json:"book_id"`
	User string `json:"user"`
	CheckoutDate string `json:"checkout_date"`
	IsGenesis bool `json:"is_genesis"`
}

type Book struct {
	ID string `json:"id"`
	Title string `json:"title"`
	Author string `json:"author"`
	PublishDate string `json:"publsh_date"`
	ISBN string `json:"isbn"`
}

func main {

}