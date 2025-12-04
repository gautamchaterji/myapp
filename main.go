package main

import (
	"fmt"
	"myapp/decrypt"
	"myapp/encrypt"
)

func main() {
	originalText := "Hello, World!"
	encryptedText := encrypt.Nimbus(originalText)
	fmt.Println("Encrypted Text:", encryptedText)
	decryptedText := decrypt.Nimbus(encryptedText)
	fmt.Println("Decrypted Text:", decryptedText)
}
