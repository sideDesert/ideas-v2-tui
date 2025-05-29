package root

import (
	"encoding/json"
	"os"

	"sideDesert/ideasv2/colors"
	"sideDesert/ideasv2/components"
	"sideDesert/ideasv2/keymap"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/lipgloss"
)

// Change this
type item string

func InitialModel() model {
	tabs := []string{
		"Ideas",
		"Projects",
		"Books",
	}
	inputStyle := lipgloss.NewStyle().Foreground(lipgloss.Color(colors.Pink))
	titleStyle := lipgloss.NewStyle().
		Bold(true).
		Foreground(lipgloss.Color(colors.Black)).
		Background(lipgloss.Color(colors.DarkGreen)).
		PaddingTop(0).
		Padding(0, 4)

	var ideas []Idea
	var books []Book
	var projects []Project

	loadData("ideas.json", &ideas)
	loadData("books.json", &books)
	loadData("projects.json", &projects)

	const defaultWidth = 20

	// Create new list
	items := []list.Item{}
	for _, i := range ideas {
		items = append(items, i)
	}
	ideasList := NewList(items)

	items = []list.Item{}
	for _, i := range books {
		items = append(items, i)
	}
	booksList := NewList(items)

	items = []list.Item{}
	for _, i := range projects {
		items = append(items, i)
	}
	projectsList := NewList(items)

	m := model{
		keys:      keymap.Keys,
		IsTouched: false,
		n_panels:  3,
		help:      help.New(),
		theme: Theme{
			inputStyle: inputStyle,
			titleStyle: titleStyle,
		},
		Tabs: components.NewTabModel(tabs, []string{"", ""}, 0),
		IdeaManager: IdeaManager{
			List:     ideasList,
			Form:     NewIdeasForm(),
			Viewport: viewport.New(0, 0),
		},
		ProjectManager: ProjectManager{
			List:     projectsList,
			Viewport: viewport.New(0, 0),
		},
		BookManager: BookManager{
			List:     booksList,
			Form:     NewBooksForm(),
			Viewport: viewport.New(0, 0),
		},
	}

	return m
}

func (m model) Init() tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	cmd := m.Tabs.Init()
	cmds = append(cmds, cmd)

	cmd = m.IdeaManager.Form.Init()
	cmds = append(cmds, cmd)

	cmd = m.BookManager.Form.Init()
	cmds = append(cmds, cmd)

	cmd = m.IdeaManager.Viewport.Init()
	cmds = append(cmds, cmd)

	return tea.Batch(cmds...)
}

func (m *model) SaveFiles() {
	data, err := json.MarshalIndent(m.IdeaManager.List.Items(), "", " ")
	if err != nil {
		panic("Couldn't marshall the ideas into json")
	}
	writeToFile(data, "ideas.json")

	data, err = json.MarshalIndent(m.BookManager.List.Items(), "", " ")
	if err != nil {
		panic("Couldn't marshall the ideas into json")
	}
	writeToFile(data, "books.json")
}

func writeToFile(data []byte, filepath string) {
	err := os.WriteFile(filepath, data, os.FileMode(os.O_WRONLY))
	if err != nil {
		panic("Couldn't write to ideas.json")
	}
}
