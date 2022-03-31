package main

import (
	"Octagon/utils"
	"github.com/charmbracelet/bubbles/progress"
	tea "github.com/charmbracelet/bubbletea"
)

func (m Model) IndexUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case ContinueMsg:
		m.state.state = msg.state
		cmd := m.SetState(progressState)
		return m, cmd

	case StopMsg:
		return m, tea.Quit

	// Is it a key press?
	case tea.KeyMsg:

		// Cool, what was the actual key pressed?
		switch msg.String() {

		// These keys should exit the program.
		case "ctrl+c", "q":
			return m, tea.Quit

		// The "up" and "k" keys move the cursor up
		case "up", "k":
			if m.index.Cursor > 0 {
				m.index.Cursor--
			}

		// The "down" and "j" keys move the cursor down
		case "down", "j":
			if m.index.Cursor < len(m.index.Choices)-1 {
				m.index.Cursor++
			}

		// The "enter" key and the space-bar (a literal space) toggle
		// the selected state for the item that the cursor is pointing at.
		case "enter", " ":
			_, ok := m.index.Selected[m.index.Cursor]
			if ok {
				delete(m.index.Selected, m.index.Cursor)
				return m, nil
			} else {
				m.index.Selected[m.index.Cursor] = m.index.Choices[m.index.Cursor]
				cmd := m.SetState(progressState)
				return m, cmd
			}
		}

	default:
		return m, nil
	}

	return m, nil
}

func (m Model) progressbarUpdate(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {

	case ContinueMsg:
		// TODO: Make an IF condition where the function will depend on the selected choice
		cmd := m.GetAllTickerNames(msg.url, utils.GetPolygonioKey(), &tickerNamesResults)
		if msg.url != "" {
			m.progress.Percent += m.progress.Increment
		}
		if m.progress.Percent > 1.0 {
			cmd := m.SetState(indexState)
			return m, cmd
		}
		return m, cmd

	case StopMsg:
		// TODO: Make an IF condition where the save function will depend on the selected choice
		WriteTickerNamesToJsonFile(tickerNamesResults)
		m.state.state = indexState
		m.progress.Percent = 0.0
		return m, nil

	// FrameMsg is sent when the progressbar bar wants to animate itself
	case progress.FrameMsg:
		progressModel, cmd := m.progress.Progress.Update(msg)
		m.progress.Progress = progressModel.(progress.Model)
		return m, cmd

	// What if the terminal window is resized?
	case tea.WindowSizeMsg:
		m.progress.Progress.Width = msg.Width - padding*2 - 4
		if m.progress.Progress.Width > maxWidth {
			m.progress.Progress.Width = maxWidth
		}
		return m, nil

	default:
		return m, nil
	}
}
