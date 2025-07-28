package views

import (
	"github.com/charmbracelet/lipgloss"
)

func RenderHeader(s string) string {
	var styleHeader lipgloss.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#ffffff")).
		Bold(true).
		Width(50).
		Height(50)
	return styleHeader.Render(s)
}
