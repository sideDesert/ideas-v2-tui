package root

import (
	"sideDesert/ideasv2/components"
	"sideDesert/ideasv2/keymap"

	"github.com/charmbracelet/bubbles/help"
	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/bubbles/viewport"
	"github.com/charmbracelet/huh"
)

type ListItem struct {
	TitleText       string `json:"title"`
	FilePath        string `json:"file_path"`
	DescriptionText string `json:"description"`
}

type Idea = ListItem
type Book = ListItem
type Project = ListItem

func (i ListItem) Title() string {
	return i.TitleText
}

func (i ListItem) Description() string {
	return i.DescriptionText
}

func (i ListItem) FilterValue() string {
	return i.Title()
}

type Panel = string

type Tab struct {
	panels []Panel
}

type Manager struct {
	tabIndex     int
	DirPath      string
	List         list.Model
	ListDelegate list.DefaultDelegate
	Viewport     viewport.Model
	Form         *huh.Form
}

type model struct {
	IsTouched      bool
	n_panels       int
	Tabs           components.TabModel
	IdeaManager    Manager
	ProjectManager Manager
	BookManager    Manager
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
