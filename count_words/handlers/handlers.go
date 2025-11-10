package handlers

import (
	"bufio"
	"fmt"
	"os"
	"strings"
	"time"
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
// It return the word object that contain counts of words.
func HandleStdin() map[string]int {
	words := make(map[string]int)
	fmt.Println("\nStart write your text here(Press Ctrl+D to end input):")

	sc := bufio.NewReader(os.Stdin)
	for {
		input_line, err := sc.ReadString('\n')
		if err != nil {
			fmt.Printf("[Error] %s\n", err)
			break
		}
		countWordsInLine(input_line, &words)
	}

	return words
}

// HandleFile read contents of <filename> and counts word occurrences,
// 
// It return the word object that contain counts of words.
func HandleFile(filename string) map[string]int {
	words := make(map[string]int)
	f, err := os.Open(filename)

	if err != nil {
		fmt.Printf("[Error] %s\n", err)
		return nil
	}
	defer f.Close()

	sc := bufio.NewScanner(f)

	start := time.Now()
	for sc.Scan() {
		line := sc.Text()
		countWordsInLine(line, &words)
	}
	elapsed := time.Since(start)
	fmt.Println("Total time:", elapsed)

	return words
}
