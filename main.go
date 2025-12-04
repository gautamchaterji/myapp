package main

import (
	"fmt"

	"github.com/gautamchaterji/myapp/decrypt"
	"github.com/gautamchaterji/myapp/encrypt"
)

func main() {
	originalText := "Hello, World!"
	encryptedText := encrypt.Nimbus(originalText)
	fmt.Println("Encrypted Text:", encryptedText)
	decryptedText := decrypt.Nimbus(encryptedText)
	fmt.Println("Decrypted Text:", decryptedText)
}
