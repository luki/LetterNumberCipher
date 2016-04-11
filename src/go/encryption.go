package main

import "strings"
import "fmt"
import "unicode/utf8"

func letterToNumber(letter string) string {
	lowercaseString := strings.ToUpper(letter)
	switch lowercaseString {
	case "A":
		return "1"
	case "B":
		return "2"
	case "C":
		return "3"
	case "D":
		return "4"
	case "E":
		return "5"
	case "F":
		return "6"
	case "G":
		return "7"
	case "H":
		return "8"
	case "I":
		return "9"
	case "J":
		return "10"
	case "K":
		return "11"
	case "L":
		return "12"
	case "M":
		return "13"
	case "N":
		return "14"
	case "O":
		return "15"
	case "P":
		return "16"
	case "Q":
		return "17"
	case "R":
		return "18"
	case "S":
		return "19"
	case "T":
		return "20"
	case "U":
		return "21"
	case "V":
		return "22"
	case "W":
		return "23"
	case "X":
		return "24"
	case "Y":
		return "25"
	case "Z":
		return "26"
	case "Ä":
		return "27"
	case "Ö":
		return "28"
	case "Ü":
		return "29"
	default:
		return "0"
	}
}

func encryption(word string) {
	currentWord := strings.ToUpper(word)
	var wordLength int = utf8.RuneCountInString(currentWord)
	fmt.Println("WordLength", wordLength)
	currentWordAsString := strings.Split(currentWord, "")

	var newWord string = ""

	for i := 0; i < wordLength; i++ {
		if currentWordAsString[i] != " " {
			newWord = newWord + letterToNumber(currentWordAsString[i])
			fmt.Println(i)
			if i != wordLength-1 {
				newWord = newWord + "-"
			}
		} else {
			newWord = newWord + " " // Space
		}

	}

	fmt.Println(newWord)

}

func main() {
	encryption("Hello")
}
