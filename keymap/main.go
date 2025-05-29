package keymap

import "github.com/charmbracelet/bubbles/key"

type KeyMap struct {
	Up         key.Binding
	Down       key.Binding
	Left       key.Binding
	Right      key.Binding
	Help       key.Binding
	Quit       key.Binding
	Submit     key.Binding
	NextTab    key.Binding
	PrevTab    key.Binding
	NextPanel  key.Binding
	PrevPanel  key.Binding
	DeleteItem key.Binding
	AddMode    key.Binding
	ReadMode   key.Binding
	EditMode   key.Binding
}

func (k KeyMap) FullHelp() [][]key.Binding {
	return [][]key.Binding{
		{k.Up, k.Down, k.Left, k.Right}, // first column
		{k.Help, k.Quit, k.Submit},      // second column
		{k.NextTab, k.NextPanel},
	}
}
func (k KeyMap) ShortHelp() []key.Binding {
	return []key.Binding{k.Help, k.Quit}
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
	Left: key.NewBinding(
		key.WithKeys("left", "h"),
		key.WithHelp("←/h", "move left"),
	),
	Right: key.NewBinding(
		key.WithKeys("right", "l"),
		key.WithHelp("→/l", "move right"),
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
		key.WithHelp("shift+j", "Shift to right tab"),
	),
	PrevTab: key.NewBinding(
		key.WithKeys("K", "ctrl+shift+tab"),
		key.WithHelp("shift+k", "Shift to left tab"),
	),
	NextPanel: key.NewBinding(
		key.WithKeys("L"),
		key.WithHelp("shift+l", "Shift to right tab"),
	),
	PrevPanel: key.NewBinding(
		key.WithKeys("H"),
		key.WithHelp("shift+h", "Shift to left tab"),
	),
	AddMode: key.NewBinding(
		key.WithKeys("a", "i", "A", "I"),
		key.WithHelp("a", "Append/Add mode"),
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
		key.WithKeys("d", "D"),
		key.WithHelp("d", "Delete Item"),
	),
}
