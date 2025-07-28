package views

import (
	"fmt"
	"github.com/charmbracelet/lipgloss"
	"golang.org/x/term"
	"os"
)

func Background() {
	width, height, err := term.GetSize(int(os.Stdout.Fd()))
	if err != nil {
		fmt.Println("Error obteniendo el tamaño de la terminal:", err)
		return
	}
	var style lipgloss.Style = lipgloss.NewStyle().
		Foreground(lipgloss.Color("#FAFAFA")).
		Background(lipgloss.Color("#7D56F4")).
		Width(width).
		Height(height)
	style.Render(" ")
}
