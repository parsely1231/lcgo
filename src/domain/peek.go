package domain

// Peek is a peek in chromatgram
type Peek struct {
	Name          string
	RetentionTime float64
	Area          int
}

// Rename return the renamed peek
func (p Peek) Rename(newName string) Peek {
	return Peek{
		Name:          newName,
		RetentionTime: p.RetentionTime,
		Area:          p.Area,
	}
}

// AreaRatio return the area ratio of total area
func (p Peek) AreaRatio(totalArea int) float64 {
	return float64(p.Area) / float64(totalArea)
}

// RelativeRT return the relative retention time (RRT)
func (p Peek) RelativeRT(baseRT float64) float64 {
	return p.RetentionTime / baseRT
}
