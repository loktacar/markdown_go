package markdown

import (
	"fmt"
	"regexp"
)

type ParagraphBlock struct {
	InlineText string
}

func ParseParagraphBlock(lines []string) (ParagraphBlock, error) {
	if len(lines) == 0 {
		return ParagraphBlock{}, BlockNotApplicableError("No content")
	}

	incorrectFirstLineIndent := regexp.MustCompile("^\\ {4,}")
	anyIndent := regexp.MustCompile("^\\s*")

	inlineText := ""

	// First line cannot be indented more than three spaces
	if incorrectFirstLineIndent.MatchString(lines[0]) {
		return ParagraphBlock{}, BlockNotApplicableError("Incorrect indent")
	}

	inlineText += anyIndent.ReplaceAllString(lines[0], "")

	for _, line := range lines[1:len(lines)] {
		inlineText += " " + anyIndent.ReplaceAllString(line, "")
	}

	return ParagraphBlock{
		inlineText,
	}, nil
}

func (pb ParagraphBlock) Render() string {
	return fmt.Sprintf("<p>%s</p>", pb.InlineText)
}
