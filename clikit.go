package clikit

import (
	"fmt"
)

type ParagraphSet struct {
	items []Paragraph
}

type Paragraph struct {
	Title string
	lines string
}

func (c *Paragraph) print_() {

	fmt.Println()
	printBold(c.Title)
	fmt.Println(c.lines)
}

func printBold(s string) {
	fmt.Printf("\x1b[1m%s\x1b[0m\n", s)
}

func NewParagraphSet() *ParagraphSet {
	return &ParagraphSet{}
}

func (ps *ParagraphSet) Add(title string, line string) {
	ps.items = append(ps.items, Paragraph{title, line})
}

func (ps *ParagraphSet) Flush() {

	for _, item := range ps.items {
		item.print_()
	}
}
