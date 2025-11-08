# Word Counter
firs mini project for learning `GoLang`.

A simple Go program that counts word occurrences in text. This program offers two modes of operation:
1. Count words from standard input
2. Count words from a text file

## Features

- Interactive command-line interface
- Support for both stdin and file input modes
- Real-time word counting
- Case-sensitive word counting
- Easy-to-read output format

## Usage

### Running the Program

```bash
go run main.go
```

### Mode 1: Standard Input
1. Select option 1 when prompted
2. Type or paste your text
3. Press Ctrl+D (EOF) to finish input
4. View the word count results

### Mode 2: File Input
1. Select option 2 when prompted
2. Enter the path to your text file
3. View the word count results

## Project Structure

```
count_words/
├── main.go          # Main program entry point
├── handlers/        # Package containing text processing logic
│   └── handlers.go
└── text1.txt       # Sample text file
```

## Dependencies

This project uses only Go standard library packages:
- `fmt` for I/O operations
- `bufio` for buffered I/O
- `strings` for string manipulation
- `os` for file operations