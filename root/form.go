package root

import (
	"github.com/charmbracelet/huh"
)

func NewForm(title string) *huh.Form {
	return huh.NewForm(
		huh.NewGroup(
			huh.NewInput().Key("title").Title(title).Prompt("> "),
			huh.NewText().Key("desc").Title("Description"),
			huh.NewConfirm().Key("confirm").Title("Confirm"),
		),
	)
}

func NewIdeasForm() *huh.Form {
	return NewForm("Idea Title")
}

func NewBooksForm() *huh.Form {
	return NewForm("Book Title")
}

func NewProjectsForm() *huh.Form {
	return NewForm("Project Title")
}

func (m *model) NewTabForm() *huh.Form {
	activeTab := m.Tabs.ActiveTab

	if activeTab == ideasTab {
		return NewIdeasForm()
	}
	if activeTab == booksTab {
		return NewBooksForm()
	}
	if activeTab == projectsTab {
		return NewProjectsForm()
	}
	return nil
}
