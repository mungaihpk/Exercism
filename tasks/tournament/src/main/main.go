package main

import (
	"bufio"
	"fmt"
	"os"
	"path/filepath"
)

func main() {
	pwd, _ := os.Getwd()

	file, err := os.Open(filepath.Join(pwd, "input_file/games.txt"))
	if err != nil {
		//panic("Could not open input file.")
		fmt.Println(err)
		fmt.Sprintf("Error: Could not open input file: %v", err)
		os.Exit(1)
	}

	write_file, err := os.Create(filepath.Join(pwd, "output_file/scores.txt"))
	if err != nil {
		//panic("Could not create output file")
		fmt.Println(err)
		fmt.Sprintf("Error: Could not create output file: %v", err)
		os.Exit(1)
	}

	reader := bufio.NewReader(file)
	writer := bufio.NewWriter(write_file)

	err = Tally(reader, writer)
	if err != nil {
		panic("Could not write scores to file")
		fmt.Sprintf("Error: Could not write to output file (scores.txt): %v", err)
		os.Exit(1)
	}
}
