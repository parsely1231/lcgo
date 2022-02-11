package ascii

import (
	"io/ioutil"
	"strings"

	"github.com/parsely1231/lcgo/domain"
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

// ToChromatogram return a chromatogram
func (form *InputASCIIForm) ToChromatogram() (*domain.Chromatogram, error) {

	var chromato *domain.Chromatogram
	for _, line := range form.lines {
		if !line.IsChromatoNameLine() {
			continue
		}
		name, err := line.FindChromatoName()
		if err != nil {
			return nil, err
		}
		chromato = domain.NewChromatogram(name)
	}

	for _, line := range form.lines {
		if line.IsPeekLine() {
			peek, err := line.ToPeek()
			if err != nil {
				return nil, err
			}
			chromato.AddPeek(peek)
		}

		if line.IsEndLine() {
			break
		}
	}
	return chromato, nil
}
