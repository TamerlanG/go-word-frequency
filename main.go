package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"
)

type Word struct {
	title string
	count int
}

var words = []Word{}

func main() {

	f, err := os.Open("text.txt")

	reg, err := regexp.Compile("[^a-zA-Z0-9]+")
	if err != nil {
		log.Fatal(err)
	}

	if err != nil {
		fmt.Println(err)
	}

	defer f.Close()

	scanner := bufio.NewScanner(f)
	scanner.Split(bufio.ScanWords)

	for scanner.Scan() {
		processedString := reg.ReplaceAllString(scanner.Text(), "")
		addWord(processedString)
	}

	bubbleSort()

	for j := 0; j <= 20; j++ {
		fmt.Println(words[j].title, words[j].count)
	}

	if err := scanner.Err(); err != nil {
		fmt.Println(err)
	}
}

func addWord(s string) {
	index := isExist(s)
	if index >= 0 {
		words[index] = Word{
			title: s,
			count: words[index].count + 1,
		}
	} else {
		var temp = Word{
			title: s,
			count: 1,
		}
		words = append(words, temp)
	}
}

func isExist(a string) int {
	for i, b := range words {
		if b.title == a {
			return i
		}
	}
	return -1
}

func bubbleSort() {
	for i := 0; i < len(words); i++ {
		for j := 0; j < len(words)-1-i; j++ {
			if words[j].count < words[j+1].count {
				words[j+1], words[j] = words[j], words[j+1]
			}
		}
	}
}
