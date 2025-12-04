package encrypt

func Nimbus(str string) string {
	encryptedText := ""
	for _, char := range str {
		encryptedText += string(char + 3)
	}
	return encryptedText
}
