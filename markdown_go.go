package main

import (
	"bufio"
	"fmt"
	"log"
	"os"

	. "github.com/loktacar/markdown_go/lib"
)

func main() {
	fmt.Println("Start..")

	file, err := os.Open("notes.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var currBlock Block
	var parseErr error
	cont := false

	scanner := bufio.NewScanner(file)
	for scanner.Scan() {

		if currBlock == nil {
			currBlock, parseErr = ParseParagraphBlock(scanner.Text())
			if parseErr != nil {
				continue
			}

		} else if b, ok := currBlock.(ParagraphBlock); ok {
			currBlock, cont, parseErr = b.ParseNext(scanner.Text())
			if parseErr != nil {
				continue
			}
			if !cont {
				fmt.Println(b.Render())
				currBlock = nil
			}
		}

		//fmt.Println(scanner.Text())
	}

	fmt.Println(currBlock.Render())

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
