package main

import (
	"encoding/hex"
	"fmt"
)

// Minimal demo: shows original, XOR-encrypted bytes (hex), and decrypted text
func main() {
	secret := "Hidden message: The eagle lands at dawn."
	key := []byte("mykey")

	fmt.Println("Original:", secret)

	enc := cipher(secret, key)
	fmt.Println("Encrypted (hex):", hex.EncodeToString(enc))

	// XORing again with the same key returns the original
	dec := cipher(string(enc), key)
	fmt.Println("Decrypted:", string(dec))
}
