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
		// "longs",
		// "lucky",
		// "dimka",
		// "scarface",
		// "manoah",
		// "victor",
	}

	merkleTree := generateMerkleTree(transactions)

	for layer, hashes := range merkleTree {
		println(layer, hashes)
		println("====================================================>>>>>>>>>>>>>>>>>")
	} 

	root := getMerkleRootHash(merkleTree)
	println("Merkle Root: ", root)

	leaf := "0xa247681e6cdde53e207de4a7a39c2e4b36bcc8e794c74bcb8f841f8051639b24"
	proofs := []string{
		"0x575644babd48c4fa383ac90e01bef43631ccc45514a762e7394d56fc87c1f663",
		"0x773d2b57449bafd0cf7ac96800a91ea65f25097a6d8b7ef5cbdfb307c76b1be0",
	}

	println(merkleProof(leaf, proofs, root, 0))

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

func generateMerkleTree(transactions []string) []string {
	hashedTransactions := convertTransactionsToHash(transactions)
	transactionLength := len(hashedTransactions)
	merkleTree := hashedTransactions;

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
		println(nextLayer);
		hashedTransactions = nextLayer
		merkleTree = append(merkleTree, nextLayer...)
		transactionLength = len(hashedTransactions)
	}

	return merkleTree
}

func merkleProof(
	leafToVerify string,
	proofs []string,
	root string,
	index uint64,
) bool {
	hashedVal := leafToVerify;
	for _, proof := range proofs {
		if (index % 2 == 0) {
			hashedVal = hash(hashedVal + proof)
		} else {
			hashedVal = hash(proof + hashedVal)
		}

		index /= 2
	}
	println("output: ",hashedVal)
	
	return hashedVal == root
}

func getMerkleRootHash(merkleTree []string) string {
	return merkleTree[len(merkleTree) - 1]
}