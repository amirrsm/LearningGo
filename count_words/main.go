package main

import "fmt"
import "count_words/handlers"

func main() {
	fmt.Println("Please select mode:")
	fmt.Println("1. Count words from standard input (stdin)")
	fmt.Println("2. Count words from a text file")

	var choice int
	for {
		fmt.Print("Enter 1 or 2: ")
		_, err := fmt.Scan(&choice)
		if err != nil {
			fmt.Println("Invalid input. Please enter 1 or 2.")
			continue
		}

		switch choice {
		case 1:
			words := handlers.HandleStdin()

			fmt.Println("\nWord Count Results:")
			for word, count := range words {
				fmt.Printf("%s: %d\n", word, count)
			}
		case 2:
			var filename string
			fmt.Print("Enter the filename: ")
			_, err := fmt.Scan(&filename)
			if err != nil {
				fmt.Println("Invalid filename.")
				continue
			}

			words := handlers.HandleFile(filename)
			fmt.Println("\nWord Count Results:")
			for word, count := range words {
				fmt.Printf("%s: %d\n", word, count)
			}
		default:
			fmt.Println("Invalid choice. Please enter 1 or 2.")
		}
	}

}
