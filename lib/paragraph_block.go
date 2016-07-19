package markdown

import (
	"fmt"
	"regexp"
)

type ParagraphBlock struct {
	InlineText string
}

func ParseParagraphBlock(line string) (ParagraphBlock, error) {
	firstLineIndent := regexp.MustCompile("^\\ {0,3}")

	// First line cannot be indented more than three spaces
	if !firstLineIndent.MatchString(line) {
		return ParagraphBlock{}, BlockNotApplicableError("Incorrect indent")
	}

	return ParagraphBlock{
		firstLineIndent.ReplaceAllString(line, ""),
	}, nil
}

func (pb ParagraphBlock) ParseNext(line string) (Block, bool, error) {
	anyIndent := regexp.MustCompile("^\\s*")

	// Lines after the first may be indented any amount
	// http://spec.commonmark.org/0.26/#example-184
	line = anyIndent.ReplaceAllString(line, "")

	pb.InlineText += fmt.Sprintf(" %s", line)

	if line == "" {
		return pb, false, nil
	} else {
		return pb, true, nil
	}
}

func (pb ParagraphBlock) Render() string {
	return fmt.Sprintf("<p>%s</p>", pb.InlineText)
}
