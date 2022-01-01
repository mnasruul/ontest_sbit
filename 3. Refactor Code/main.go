package main

import (
	"fmt"
	"strings"
)

const exampleWord = "the(crow)bar"

func main() {
	fmt.Println("Result Before Refactor: " + findFirstStringInBracket(exampleWord))
	fmt.Println("Result After Refactor: " + refactorFindFirstStringInBracket(exampleWord))
}
func findFirstStringInBracket(str string) string {
	if len(str) > 0 {
		indexFirstBracketFound := strings.Index(str, "(")
		if indexFirstBracketFound >= 0 {
			runes := []rune(str)
			wordsAfterFirstBracket := string(runes[indexFirstBracketFound:len(str)])
			indexClosingBracketFound := strings.Index(wordsAfterFirstBracket, ")")
			if indexClosingBracketFound >= 0 {
				runes := []rune(wordsAfterFirstBracket)
				return string(runes[1:indexClosingBracketFound])
			} else {
				return ""
			}
		} else {
			return ""
		}
	} else {
		return ""
	}

}

func refactorFindFirstStringInBracket(str string) string {
	indexFirstBracketFound := strings.IndexByte(str, '(')
	if indexFirstBracketFound < 0 {
		return ""
	}
	indexFirstBracketFound++
	wordsAfterFirstBracket := str[indexFirstBracketFound:]
	indexClosingBracketFound := strings.IndexByte(wordsAfterFirstBracket, ')')
	if indexClosingBracketFound < 0 {
		return ""
	}
	return str[indexFirstBracketFound : indexFirstBracketFound+indexClosingBracketFound]
}
