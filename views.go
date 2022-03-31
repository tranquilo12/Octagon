package main

import (
	"fmt"
	"strings"
)

func (m Model) IndexView() string {
	// The header
	s := "\nWhat do you want to download?\n\n"
	for i, choice := range m.index.Choices {

		// Is the cursor pointing at this choice?
		cursor := " " // no cursor
		if m.index.Cursor == i {
			cursor = ">" // cursor!
		}

		// Is this choice selected?
		checked := " " // not selected
		if _, ok := m.index.Selected[i]; ok {
			checked = "x" // selected!
		}

		// Render the row
		s += fmt.Sprintf("%s [%s] %s\n", cursor, checked, choice)
	}

	// Send the UI for rendering
	return s
}

func (m Model) progressbarView() string {
	pad := strings.Repeat(" ", padding)
	s := "\n" + pad + "Downloading ticker names..." + pad
	s += "\n" + pad + m.progress.Progress.ViewAs(m.progress.Percent) + "\n\n" + pad
	return s
}
