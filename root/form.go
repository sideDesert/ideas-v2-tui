package root

import (
	"fmt"
	"os"

	tea "github.com/charmbracelet/bubbletea"
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

func (m *model) HandleUpdateIdeaForm(msg tea.Msg) tea.Cmd {
	newForm, cmd := m.IdeaManager.Form.Update(msg)
	if form, ok := newForm.(*huh.Form); ok {
		m.IdeaManager.Form = form
	} else {
		fmt.Println("NO form updates")
	}

	// Check if form is completed
	if m.IdeaManager.Form.State == huh.StateCompleted {
		title := m.IdeaManager.Form.GetString("title")
		desc := m.IdeaManager.Form.GetString("desc")
		confirm := m.IdeaManager.Form.GetBool("confirm")

		newIdea := Idea{
			TitleText:       title,
			DescriptionText: desc,
		}

		DEBUG := false
		if DEBUG {
			f, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			fmt.Fprintf(f, "DEBUG: title='%s', desc='%s', confirm=%v\n", title, desc, confirm)
			f.Close()
		}

		if confirm {
			m.IdeaManager.List.InsertItem(len(m.IdeaManager.List.Items()), newIdea)
			m.SaveFiles()
		}
		m.mode = Read
	}

	return cmd
}

func (m *model) HandleUpdateBookForm(msg tea.Msg) tea.Cmd {
	newForm, cmd := m.BookManager.Form.Update(msg)
	if form, ok := newForm.(*huh.Form); ok {
		m.BookManager.Form = form
	} else {
		fmt.Println("NO form updates")
	}

	// Check if form is completed
	if m.BookManager.Form.State == huh.StateCompleted {
		title := m.BookManager.Form.GetString("title")
		desc := m.BookManager.Form.GetString("desc")
		confirm := m.BookManager.Form.GetBool("confirm")

		newIdea := Book{
			TitleText:       title,
			DescriptionText: desc,
		}

		DEBUG := false
		if DEBUG {
			f, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			fmt.Fprintf(f, "DEBUG: title='%s', desc='%s', confirm=%v\n", title, desc, confirm)
			f.Close()
		}

		if confirm {
			m.BookManager.List.InsertItem(len(m.BookManager.List.Items()), newIdea)
			m.SaveFiles()
		}
		m.mode = Read
	}

	return cmd
}
