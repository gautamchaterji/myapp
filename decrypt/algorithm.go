package decrypt

func Nimbus(str string) string {
	decryptedText := ""
	for _, char := range str {
		decryptedText += string(char - 3)
	}
	return decryptedText
}
