package node

import (
	"d3z41k/blockchain-bar/database"
	"encoding/hex"
	"testing"
)

func TestValidBlockHash(t *testing.T) {
	// Create a random hex string starting with 6 zeroes
	hexHash := "000000fa04f816039...a4db586086168edfa"
	var hash = database.Hash{}

	// Convert it to raw bytes
	hex.Decode(hash[:], []byte(hexHash))
	// Validate the hash
	isValid := database.IsBlockHashValid(hash)
	if !isValid {
		t.Fatalf("hash '%s' with 6 zeroes should be valid", hexHash)
	}
}
