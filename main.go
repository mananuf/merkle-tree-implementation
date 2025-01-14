package main

import (
    "crypto/sha256"
	"encoding/hex"
)

func main()  {
	transactions := []string{
		"suleiman",
		"kingsley",
		"emmanuel",
		"jethro",
		"longs",
		"lucky",
		"dimka",
		"scarface",
		"manoah",
		"victor",
	}

	merkleRoot := generateMerkleRootHash(transactions)

	println("Merkle Root: ", merkleRoot)
}

func hash(tx string) string {
	hasher := sha256.New() // returns sha256 checksum
	hasher.Write([]byte(tx))
	_sha256 := hex.EncodeToString(hasher.Sum(nil))

	return "0x" + _sha256
}

func convertTransactionsToHash(transactions []string) []string {
	var hashedTrasactions []string

	for _, transaction := range transactions {
		hashedTrasactions = append(hashedTrasactions, hash(transaction))
	}

	return hashedTrasactions
}

func generateMerkleRootHash(transactions []string) string {
	hashedTransactions := convertTransactionsToHash(transactions)
	transactionLength := len(hashedTransactions)

	for transactionLength > 1 {
		// when the length of hashed Array is odd, duplicate the last hash to balance the hash pairing
		if transactionLength % 2 != 0 {
			hashedTransactions = append(hashedTransactions, hashedTransactions[transactionLength - 1])
			transactionLength++
		}
		
		var nextLayer []string
		for index := 0; index < transactionLength; index += 2 {
			combinedHash := hash(hashedTransactions[index] + hashedTransactions[index + 1])
			nextLayer = append(nextLayer, combinedHash)
		}

		hashedTransactions = nextLayer
		transactionLength = len(hashedTransactions)
	}

	return hashedTransactions[0]
}