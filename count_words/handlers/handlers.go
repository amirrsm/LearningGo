package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
)

// countWordsInLine processes a single line of text and updates the word count map.
func countWordsInLine(line string, words *map[string]int) {
	tokens := strings.SplitSeq(line, " ")
	for token := range tokens {
		cleanedToken := strings.TrimSpace(token)
		if cleanedToken != "" {
			(*words)[cleanedToken]++
		}
	}
}

// HandleStdin reads from standard input and counts word occurrences.
//
// Stops reading on EOF (Ctrl+D).
//
// It prints the word count results to the console.
func HandleStdin() (words map[string]int) {
	fmt.Println("\nStart write your text here(Press Ctrl+D to end input):")

	words = make(map[string]int)
	sc := bufio.NewReader(os.Stdin)
	for {
		input_line, err := sc.ReadString('\n')
		if err != nil {
			break
		}
		countWordsInLine(input_line, &words)
	}

	return
}

func HandleFile(filename string) {
	// Implementation for handling word count from a file
}
