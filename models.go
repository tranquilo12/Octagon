package main

import "github.com/charmbracelet/bubbles/progress"

const totalTickers = 25

// StateModel For the state of the tui
type StateModel struct{ state int }

// IdxModel for the index page
type IdxModel struct {
	Choices  []string
	Cursor   int
	Selected map[int]interface{}
}

// NewIndexModel function that returns the model for the index page
func NewIndexModel() IdxModel {
	return IdxModel{
		[]string{"Refresh Ticker Names"},
		0,
		map[int]interface{}{},
	}
}

// PbModel for the progress page
type PbModel struct {
	Url       string
	Increment float64
	Percent   float64
	Progress  progress.Model
}

// NewProgressbarModel function that returns the model for the progress page
func NewProgressbarModel(prog progress.Model) PbModel {
	return PbModel{
		"",
		1 / float64(totalTickers),
		0,
		prog,
	}
}
