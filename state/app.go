package state

import (
	// "fmt"
	// view "github.com/DumbNoxx/actio/views"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

type Model struct {
	Choices  []string
	Cursor   int
	Selected map[int]struct{}
	Width    int
	Height   int
}

func InitialModel() Model {
	return Model{
		Choices:  []string{"Buy carrots", "Buy celery", "Buy kohlrabi"},
		Selected: make(map[int]struct{}),
		Width:    0,
		Height:   0,
		Cursor:   0,
	}
}

func (m Model) Init() tea.Cmd {
	return nil
}

func (m Model) Update(msg tea.Msg) (tea.Model, tea.Cmd) {
	switch msg := msg.(type) {
	case tea.WindowSizeMsg:
		m.Width = msg.Width
		m.Height = msg.Height
		return m, nil
	case tea.KeyMsg:
		switch msg.String() {
		case "ctrl+c", "q":
			return m, tea.Quit
		case "up", "k":
			if m.Cursor > 0 {
				m.Cursor--
			}
		case "down", "j":
			if m.Cursor < len(m.Choices)-1 {
				m.Cursor++
			}
		case "enter":
			_, ok := m.Selected[m.Cursor]
			if ok {
				delete(m.Selected, m.Cursor)
			} else {
				m.Selected[m.Cursor] = struct{}{}
			}

		}
	}
	return m, nil
}

func (m Model) View() string {
	contentStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color("#FAFAFA"))

	message := "Dentro del contenedor"
	message2 := "Hola"

	renderendContent := contentStyle.Render(message, message2)

	ui := lipgloss.Place(
		m.Width,
		m.Height,
		lipgloss.Center,
		lipgloss.Center,
		renderendContent,
	)

	backgroundStyle := lipgloss.NewStyle().
		Background(lipgloss.Color("#5D5D5D"))

	return backgroundStyle.Render(ui)
}
