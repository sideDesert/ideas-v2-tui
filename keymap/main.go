package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up    key.Binding
	Down  key.Binding
	Left  key.Binding
	Right key.Binding

	NextTab   key.Binding
	PrevTab   key.Binding
	NextPanel key.Binding
	PrevPanel key.Binding

	AddMode  key.Binding
	ReadMode key.Binding
	EditMode key.Binding

	DeleteItem key.Binding
	RenameItem key.Binding
	EditItem   key.Binding

	Help   key.Binding
	Submit key.Binding
	Quit   key.Binding

	Confirm key.Binding
	Cancel  key.Binding
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.DeleteItem, k.RenameItem, k.EditItem},
		{k.Up, k.Down},             // first column
		{k.Help, k.Quit, k.Submit}, // second column
		{k.NextTab, k.PrevTab, k.NextPanel, k.PrevPanel},
	}
}
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Quit, k.Help, k.Up, k.Down, k.NextPanel, k.PrevPanel}
}

var Keys = KeyMap{
	Up: key.NewBinding(
		key.WithKeys("up", "k"),
		key.WithHelp("↑/k", "move up"),
	),
	Down: key.NewBinding(
		key.WithKeys("down", "j"),
		key.WithHelp("↓/j", "move down"),
	),
	Help: key.NewBinding(
		key.WithKeys("?"),
		key.WithHelp("?", "toggle help"),
	),
	Quit: key.NewBinding(
		key.WithKeys("q", "ctrl+c"),
		key.WithHelp("q", "quit"),
	),
	Submit: key.NewBinding(
		key.WithKeys("enter", " "),
		key.WithHelp("enter", "toggle done"),
	),
	NextTab: key.NewBinding(
		key.WithKeys("J", "ctrl+tab"),
		key.WithHelp("J", "Shift to right tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("K", "ctrl+shift+tab"),
		key.WithHelp("K", "Shift to left tab"),
	),
	NextPanel: key.NewBinding(
		key.WithKeys("L"),
		key.WithHelp("L/→", "Shift to right tab"),
	),
	PrevPanel: key.NewBinding(
		key.WithKeys("H"),
		key.WithHelp("H/←", "Shift to left tab"),
	),
	AddMode: key.NewBinding(
		key.WithKeys("a", "i", "A", "I"),
		key.WithHelp("a/i", "Append/Add mode"),
	),
	EditMode: key.NewBinding(
		key.WithKeys("c", "e"),
		key.WithHelp("c/e", "Edit/Change mode"),
	),
	ReadMode: key.NewBinding(
		key.WithKeys("esc", "ctrl+o"),
		key.WithHelp("esc", "Read mode"),
	),
	DeleteItem: key.NewBinding(
		key.WithKeys("x", "D"),
		key.WithHelp("x/D", "Delete Item"),
	),
	Confirm: key.NewBinding(
		key.WithKeys("y", "Y"),
		key.WithHelp("y", "Confirm"),
	),
	Cancel: key.NewBinding(
		key.WithKeys("n", "N"),
		key.WithHelp("n", "Cancel"),
	),
}
