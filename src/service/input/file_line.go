package ascii

import (
	"fmt"
	"strconv"
	"strings"
	"unicode"

	"github.com/parsely1231/lcgo/domain"
)

// Line is a line of input ascii file
type Line struct {
	words []string
}

// NewLine create a Line from raw file line
func NewLine(fileLine string) Line {
	words := strings.Split(fileLine, " ")
	return Line{
		words: words,
	}
}

// IsChromatoNameLine return true if line is a type of chromatogram name
func (l Line) IsChromatoNameLine() bool {
	if len(l.words) < 3 {
		return false
	}
	return l.words[0] == "Sample" && l.words[1] == "Name"
}

// FindChromatoName return a name
func (l Line) FindChromatoName() (string, error) {
	if !l.IsChromatoNameLine() {
		return "", fmt.Errorf("line is invalid")
	}
	return l.words[2], nil
}

// IsPeekLine return true if line is a type of peek
func (l Line) IsPeekLine() bool {
	headWord := l.words[0]
	for _, c := range headWord {
		if !unicode.IsDigit(c) {
			return false
		}
	}
	return true
}

// IsEndLine return true if line is a type of end line
func (l Line) IsEndLine() bool {
	return l.words[0] == "[Compound Results(PDA)]"
}

// ToPeek create a Peek
func (l Line) ToPeek() (*domain.Peek, error) {
	if len(l.words) < 5 {
		return nil, l.errorLength(5)
	}
	rt, err := l.findRT()
	if err != nil {
		return nil, fmt.Errorf("line is invalid")
	}

	area, err := l.findArea()
	if err != nil {
		return nil, fmt.Errorf("line is invalid")
	}

	return domain.NewUnnamedPeek(rt, area), nil

}

func (l Line) findRT() (float64, error) {
	rt, err := strconv.ParseFloat(l.words[1], 64)
	if err != nil {
		return 0, fmt.Errorf("line is invalid")
	}
	return rt, nil
}

func (l Line) findArea() (int, error) {
	area, err := strconv.Atoi(l.words[4])
	if err != nil {
		return 0, fmt.Errorf("line is invalid")
	}
	return area, nil
}

func (l Line) errorLength(requiredLength int) error {
	return fmt.Errorf(
		`%v is invalid. the length of line.words needs %v. but %v`, l, requiredLength, len(l.words),
	)
}
