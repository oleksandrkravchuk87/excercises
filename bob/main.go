package main

import (
	"fmt"
	"log"
)

func main() {
	//reader := bufio.NewReader(os.Stdin)
	//fmt.Println("Enter text")
	//text, _ := reader.ReadString('\n')
	//fmt.Println(text)
	//Hey(text)

	defer fmt.Println("deddefer")

	i := [...]int{3, 5}

	fmt.Printf("content %v, typr %T \n", i, i)
	log.Fatal("log fastal")
}

//func Hey(word string) string {
//	word = strings.TrimSpace(word)
//	switch {
//	case strings.HasSuffix(word, "?"):
//		fmt.Println("question")
//		if isUpperCase(word[:len(word)-1]) {
//			return "Calm down, I know what I'm doing!"
//		}
//		return "Sure."
//	case strings.HasSuffix(word, "!") || isUpperCase(word):
//
//		if !isUpperCase(word[:len(word)-1]) {
//			return "Whatever."
//		}
//
//		return "Whoa, chill out!"
//
//	case isWhiteSpace(word):
//		return "Fine. Be that way!"
//
//	default:
//		return "Whatever."
//
//	}
//}
//
//func isWhiteSpace(s string) bool {
//	for _, r := range []rune(s) {
//		if !unicode.IsSpace(r) {
//			return false
//		}
//	}
//	return true
//}
//
//func isUpperCase(s string) bool {
//	var count int
//	for _, r := range []rune(s) {
//		if unicode.IsLetter(r) {
//			count++
//
//			if !unicode.IsUpper(r) {
//				return false
//			}
//		}
//	}
//	if count == 0 {
//		return false
//	}
//	return true
//}
