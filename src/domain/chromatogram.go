package domain

// Chromatogram is a chromatgram output from HPLC
type Chromatogram struct {
	Name  string
	peeks []Peek
}

// AddPeek add peek to chromatogram
func (c *Chromatogram) AddPeek(peek Peek) {
	c.peeks = append(c.peeks, peek)
}
