package root

import (
	"encoding/json"
	"os"

	"sideDesert/ideasv2/components"
	"sideDesert/ideasv2/keymap"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	tea "github.com/charmbracelet/bubbletea"
)

// Change this
type item string

func InitialModel() model {
	tabs := []string{
		"Ideas",
		"Projects",
		"Books",
	}

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
	ideaDelegate := list.NewDefaultDelegate()
	ideasList := NewList(items, ideaDelegate)

	items = []list.Item{}
	for _, i := range books {
		items = append(items, i)
	}
	bookDelegate := list.NewDefaultDelegate()
	booksList := NewList(items, bookDelegate)

	items = []list.Item{}
	for _, i := range projects {
		items = append(items, i)
	}
	projectDelegate := list.NewDefaultDelegate()
	projectsList := NewList(items, projectDelegate)

	m := model{
		keys:      keymap.Keys,
		IsTouched: false,
		n_panels:  3,
		help:      help.New(),
		theme:     DefaultTheme(),
		Tabs:      components.NewTabModel(tabs, []string{"", ""}, 0),
		IdeaManager: Manager{
			tabIndex:     ideasTab,
			List:         ideasList,
			ListDelegate: ideaDelegate,
			Form:         NewIdeasForm(),
			Viewport:     viewport.New(0, 0),
		},
		ProjectManager: Manager{
			tabIndex:     projectsTab,
			List:         projectsList,
			ListDelegate: projectDelegate,
			Form:         NewForm("New Project"),
			Viewport:     viewport.New(0, 0),
		},
		BookManager: Manager{
			tabIndex:     booksTab,
			List:         booksList,
			ListDelegate: bookDelegate,
			Form:         NewBooksForm(),
			Viewport:     viewport.New(0, 0),
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
