package ascii

import (
	"io/ioutil"
	"strings"
)

// InputASCIIForm is a input form to create experiment data
type InputASCIIForm struct {
	lines []Line
}

// NewInputASCIIText return a InputASCIIForm
func NewInputASCIIText(fileName string) (*InputASCIIForm, error) {
	bytes, err := ioutil.ReadFile(fileName)
	if err != nil {
		return nil, err
	}
	textLines := strings.Split(string(bytes), "\n")

	lines := make([]Line, 0)
	for _, textLine := range textLines {
		lines = append(lines, NewLine(textLine))
	}
	return &InputASCIIForm{lines: lines}, nil
}
