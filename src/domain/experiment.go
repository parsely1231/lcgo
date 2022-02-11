package domain

import (
	"fmt"
)

// Experiment is a experiment.
// One experiment has multiple chromatgram.
type Experiment struct {
	chromatogramMap map[string]*Chromatogram
}

// AddChromatogram adds to a chromatgram to inner map, or return an error if the chromatogram name is duplicated.
func (exp Experiment) AddChromatogram(chromatgram *Chromatogram) error {
	_, exist := exp.chromatogramMap[chromatgram.Name]
	if exist {
		return fmt.Errorf(`%v exist. chromatogram name must be unique`, chromatgram.Name)
	}
	exp.chromatogramMap[chromatgram.Name] = chromatgram
	return nil
}

// NewExperiment return a blank experiment
func NewExperiment() *Experiment {
	return &Experiment{
		chromatogramMap: make(map[string]*Chromatogram),
	}
}
