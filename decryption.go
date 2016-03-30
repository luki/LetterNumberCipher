package main

import "unicode/utf8"
import "strings"
import "fmt"

//	Please note that the decryption algorithm doesn't work all fine yet.

func backTranslation(number string) string {
	switch number {
	case "1":
		return "A"
	case "2":
		return "B"
	case "3":
		return "C"
	case "4":
		return "D"
	case "5":
		return "E"
	case "6":
		return "F"
	case "7":
		return "G"
	case "8":
		return "H"
	case "9":
		return "I"
	case "10":
		return "J"
	case "11":
		return "K"
	case "12":
		return "L"
	case "13":
		return "M"
	case "14":
		return "N"
	case "15":
		return "O"
	case "16":
		return "P"
	case "17":
		return "Q"
	case "18":
		return "R"
	case "19":
		return "S"
	case "20":
		return "T"
	case "21":
		return "U"
	case "22":
		return "V"
	case "23":
		return "W"
	case "24":
		return "X"
	case "25":
		return "Y"
	case "26":
		return "Z"
	default:
		return "?"
	}
}

func decryption(encryptedWord string) string {

	var wordLength int = utf8.RuneCountInString(encryptedWord)

	wordAsArray := strings.Split(encryptedWord, "")
	newWord := ""

	var ignoreNext = false

	for i := 0; i < wordLength; i++ {
		if wordAsArray[i] != " " && wordAsArray[i] != "-" {
			var nextIndex = i + 1
			if wordAsArray[nextIndex] != " " && wordAsArray[nextIndex] != "-" {
				// Do extensional stuff & normal stuff
				var twoNumbers = wordAsArray[i] + wordAsArray[nextIndex]
				newWord = newWord + backTranslation(twoNumbers)
				ignoreNext = true
			} else {
				if ignoreNext != false {
					ignoreNext = false
				} else {
					newWord = newWord + backTranslation(wordAsArray[i])
				}
			}

		} else if wordAsArray[i] == "-" {
			fmt.Println("Score!")
		} else if wordAsArray[i] == " " {
			fmt.Println("Space!")
		}
	}

	fmt.Println(newWord)
	return newWord
}

func main() {
	decryption("8-5-12-12-15 8-5")
}
