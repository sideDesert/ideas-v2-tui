package root

import (
	"sideDesert/ideasv2/components"
	"sideDesert/ideasv2/keymap"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/textarea"
	"github.com/charmbracelet/huh"
	"github.com/charmbracelet/lipgloss"
)

type Idea struct {
	TitleText       string `json:"title"`
	DescriptionText string `json:"description"`
}

func (i Idea) Title() string {
	return i.TitleText
}

func (i Idea) Description() string {
	return i.DescriptionText
}

func (i Idea) FilterValue() string {
	return i.Title()
}

type Book struct {
	TitleText       string `json:"title"`
	DescriptionText string `json:"description"`
}

func (i Book) Title() string {
	return i.TitleText
}
func (i Book) Description() string {
	return i.DescriptionText
}

func (i Book) FilterValue() string {
	return i.Title()
}

type Project struct {
	TitleText       string `json:"title"`
	DescriptionText string `json:"pomodoro"`
}

func (i Project) Title() string {
	return i.TitleText
}

func (i Project) Description() string {
	return i.DescriptionText
}

func (i Project) FilterValue() string {
	return i.Title()
}

type Panel = string

type Tab struct {
	panels []Panel
}

type Theme struct {
	inputStyle lipgloss.Style
	titleStyle lipgloss.Style
}

type IdeaManager struct {
	List     list.Model
	TextArea textarea.Model
	Form     *huh.Form
}

type ProjectManager struct {
	List     list.Model
	TextArea textarea.Model
	Form     *huh.Form
}

type BookManager struct {
	List     list.Model
	TextArea textarea.Model
	Form     *huh.Form
}

type model struct {
	IsTouched      bool
	Tabs           components.TabModel
	IdeaManager    IdeaManager
	ProjectManager ProjectManager
	BookManager    BookManager
	mode           int
	activePanel    int
	cursor         int
	keys           keymap.KeyMap
	theme          Theme
	help           help.Model
	quitting       bool
	width          int
	height         int
}
