package main

import (
	"bufio"
	"fmt"
	"log"
	"os"
	"regexp"

	. "github.com/loktacar/markdown_go/lib"
)

func main() {
	fmt.Println("Start..")

	file, err := os.Open("notes.md")
	if err != nil {
		log.Fatal(err)
	}
	defer file.Close()

	var blocks []Block
	//var parseErr error
	//cont := false

	scanner := bufio.NewScanner(file)
	var textBlock []string
	blockBreak := regexp.MustCompile("^([ \t]*|[ \t]{0,3}=+[ \t]*|[ \t]{0,3}-+[ \t]*)$")
	for scanner.Scan() {
		line := scanner.Text()

		if blockBreak.MatchString(line) {
			currBlock, parseErr := ParseParagraphBlock(textBlock)
			if parseErr == nil {
				blocks = append(blocks, currBlock)
			}
			textBlock = []string{}
		} else {
			textBlock = append(textBlock, line)
		}

		//fmt.Println(scanner.Text())
	}

	for _, block := range blocks {
		fmt.Println(block.Render())
	}

	if err := scanner.Err(); err != nil {
		log.Fatal(err)
	}

	fmt.Println("Done!")
}
