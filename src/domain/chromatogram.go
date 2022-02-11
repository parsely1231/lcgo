package domain

import (
	"fmt"
)

// Chromatogram is a chromatgram output from HPLC
type Chromatogram struct {
	Name     string
	peeksMap map[int]*Peek
}

// AddPeek add peek to chromatogram
func (c *Chromatogram) AddPeek(peek *Peek) {
	id := len(c.peeksMap) + 1
	c.peeksMap[id] = peek
}

// FindPeek return a peek, or error if id does not exist.
func (c *Chromatogram) FindPeek(id int) (*Peek, error) {
	peek, exist := c.peeksMap[id]
	if !exist {
		return nil, fmt.Errorf(`id: %v does not exist`, id)
	}
	return peek, nil
}

// NewChromatogram return a blank chromatogram
func NewChromatogram(name string) *Chromatogram {
	return &Chromatogram{
		Name:     name,
		peeksMap: make(map[int]*Peek),
	}
}
