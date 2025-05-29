package root

import (
	"sideDesert/ideasv2/colors"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func NewList(items []list.Item) list.Model {
	delegate := list.NewDefaultDelegate()

	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		PaddingLeft(2).                                // or use 0 and let the list add the marker
		Foreground(lipgloss.Color(colors.ActiveText)). // Green text
		Bold(true).
		PaddingLeft(2)

	delegate.Styles.SelectedDesc = lipgloss.NewStyle().
		PaddingLeft(2). // or use 0 and let the list add the marker
		Foreground(lipgloss.Color(colors.InactiveText)).
		PaddingLeft(2).
		Italic(true)

	delegate.Styles.DimmedTitle = lipgloss.NewStyle().
		PaddingLeft(2).                           // or use 0 and let the list add the marker
		Foreground(lipgloss.Color(colors.Green)). // Green text
		Bold(true).
		PaddingLeft(1)

	delegate.Styles.DimmedDesc = lipgloss.NewStyle().
		PaddingLeft(2). // or use 0 and let the list add the marker
		Foreground(lipgloss.Color(colors.Green)).
		PaddingLeft(1).
		Italic(true)

	lm := list.New(items, delegate, 0, 0)
	lm.SetShowTitle(false)
	lm.SetShowHelp(false)
	lm.FilterInput.Focus()

	return lm
}
