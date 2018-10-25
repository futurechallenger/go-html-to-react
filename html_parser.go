package parser

import (
	"fmt"
	"go-html-to-react/lib"
)

// NewHTMLParser generate a new parser for use.
func NewHTMLParser() {
	fmt.Println("Hello Parser")
	parser := &lib.Parser{}

	parser.Parse("<div>Hello World!</div>")

	fmt.Println("Parse done")
}
