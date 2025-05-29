package root

import (
	"github.com/charmbracelet/bubbles/list"
)

func NewList(items []list.Item, delegate list.DefaultDelegate) list.Model {

	lm := list.New(items, delegate, 0, 0)
	lm.SetShowTitle(false)
	lm.SetShowHelp(false)
	lm.FilterInput.Focus()

	return lm
}
