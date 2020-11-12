package module01

// Reverse will return the provided word in reverse
// order. Eg:
//
//   Reverse("cat") => "tac"
//   Reverse("alphabet") => "tebahpla"
//
func Reverse(word string) string {
	var result string
	// works with runes hences characters that take more than abyte
	for _,char := range word {
		result = string(char) + result
	}
	return result
}
