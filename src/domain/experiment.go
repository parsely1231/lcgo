package domain

import (
	"errors"
	"fmt"
)

// Experiment is a experiment.
// One experiment has multiple chromatgram.
type Experiment struct {
	chromatogramMap map[string]Chromatogram
}

// AddChromatogram adds to a chromatgram to inner map, or return an error if the chromatogram name is duplicated.
func (exp Experiment) AddChromatogram(chromatgram Chromatogram) error {
	_, exist := exp.chromatogramMap[chromatgram.Name]
	if exist {
		message := fmt.Sprintf(`%v exist. chromatogram name must be unique`, chromatgram.Name)
		return errors.New(message)
	}
	exp.chromatogramMap[chromatgram.Name] = chromatgram
	return nil
}
