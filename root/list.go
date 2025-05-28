package root

import (
	"sideDesert/ideasv2/colors"

	"github.com/charmbracelet/bubbles/list"
	"github.com/charmbracelet/lipgloss"
)

func NewList(items []list.Item, title string, titleStyle lipgloss.Style) list.Model {
	delegate := list.NewDefaultDelegate()
	delegate.Styles.SelectedTitle = lipgloss.NewStyle().
		PaddingLeft(2).                           // or use 0 and let the list add the marker
		Foreground(lipgloss.Color(colors.Green)). // Green text
		Bold(true).
		Border(lipgloss.NormalBorder(), false, false, false, true). // left, top, right, bottom
		BorderForeground(lipgloss.Color("10")).                     // Optional color
		PaddingLeft(1)

	delegate.Styles.SelectedDesc = lipgloss.NewStyle().
		PaddingLeft(2). // or use 0 and let the list add the marker
		Foreground(lipgloss.Color("#AAAAAA")).
		Border(lipgloss.NormalBorder(), false, false, false, true). // left, top, right, bottom
		BorderForeground(lipgloss.Color("10")).                     // Optional color
		PaddingLeft(1).
		Italic(true)

	lm := list.New(items, delegate, 40, 30)
	lm.Styles.Title = titleStyle
	lm.Title = title
	lm.SetShowHelp(false)
	return lm
}
