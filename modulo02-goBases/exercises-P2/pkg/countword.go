package main
import "fmt"

func countWords(word string) {
	fmt.Println(len(word))
	for index := 0; index < len(word); index++ {
		fmt.Println(string(word[index]))
	}

}