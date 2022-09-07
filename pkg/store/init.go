package store

import (
	"fmt"
	"log"

	badger "github.com/dgraph-io/badger/v3"
)

func InitDB() (*badger.DB, error) {
	fmt.Println("INIT DB")
	// Open the Badger database located in the /tmp/badger directory.
	// It will be created if it doesn't exist.
	bdb, err := badger.Open(badger.DefaultOptions("/tmp/badger"))
	if err != nil {
		log.Fatal(err)
	}
	return bdb, err
}
