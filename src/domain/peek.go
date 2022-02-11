package domain

// Peek is a peek in chromatgram
type Peek struct {
	Name          string
	RetentionTime float64
	Area          int
}

// Rename change peek name
func (p *Peek) Rename(newName string) {
	p.Name = newName
}

// AreaRatio return the area ratio of total area
func (p *Peek) AreaRatio(totalArea int) float64 {
	return float64(p.Area) / float64(totalArea)
}

// RelativeRT return the relative retention time (RRT)
func (p *Peek) RelativeRT(baseRT float64) float64 {
	return p.RetentionTime / baseRT
}

// NewUnnamedPeek return a no name peek.
func NewUnnamedPeek(retentionTime float64, area int) *Peek {
	return &Peek{
		Name:          "",
		RetentionTime: retentionTime,
		Area:          area,
	}
}
