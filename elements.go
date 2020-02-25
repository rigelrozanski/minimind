package main

import (
	"fmt"
	"strings"
)

const (
	MainElementTextWrap = 20
)

type Element struct {
	ElID    uint64
	Content string
	ConnIDs []uint64
}

// NewMapElement creates a new MapElement object
func NewElement(elID uint64, content string, connIDs []uint64) Element {
	return Element{
		ElID:    elID,
		Content: content,
		ConnIDs: connIDs,
	}
}

func (el Element) RenderBordered() string {
	contentWrapped := WordWrapArr(el.Content, MainElementTextWrap)

	// get the biggest element
	maxLen := 0
	for _, line := range contentWrapped {
		if len(line) > maxLen {
			maxLen = len(line)
		}
	}

	// created the base output
	out := fmt.Sprintf("┏%v┓\n", strings.Repeat("━", maxLen+2))
	for _, line := range contentWrapped {

		out += fmt.Sprintf("┃ %v%v ┃\n", line, strings.Repeat(" ", maxLen-len(line)))
	}
	out += fmt.Sprintf("┗%v┛", strings.Repeat("━", maxLen+2))

	return out
}

//__________________________________________________

type Rect struct {
	XMin uint8
	YMin uint8
	XMax uint8
	YMax uint8
}

func ElementOccupies(mindmap, renderedEl string) Rect {
	return Rect{}
}
