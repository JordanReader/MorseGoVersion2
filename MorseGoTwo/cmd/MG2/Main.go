package main

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)
var morseAlphabet = [26] string {".-", "-...", "-.-.", "-..", ".", "..-.", "--.", "....", "..", ".---", "-.-", ".-..", "--", "-.", "---", ".--.", "--.-", ".-.", "...", "-", "..-", "...-", ".--", "-..-", "-.--", "--.."}
var alphabet = [26] string {"a", "b", "c", "d", "e", "f", "g", "h", "i", "j", "k", "l", "m", "n", "o", "p", "q", "r", "s", "t", "u", "v", "w", "x", "y", "z"}

func main() {

	fmt.Println("Hi there! Please type a letter, and I'll translate it into morse code!")
	reader := bufio.NewReader(os.Stdin)
	response, _ := reader.ReadString('\n')
	var letter string = strings.ToLower(response)
	letter = TrimSuffix(letter, "\n")

	var i int
	for i = 0; i<len(alphabet); i++ {

		if alphabet[i] == letter {
			fmt.Println(morseAlphabet[i])
			break
		} else if i ==25 {
			fmt.Println("Not found.")
		}

		}

	}

func TrimSuffix(s, suffix string) string {
	if strings.HasSuffix(s, suffix) {
		s = s[:len(s)-len(suffix)]
	}
	return s
}




