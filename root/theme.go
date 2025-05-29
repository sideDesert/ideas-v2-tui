package root

import (
	"sideDesert/ideasv2/colors"

	"github.com/charmbracelet/lipgloss"
)

type ListTheme struct {
	selectedTitle lipgloss.Style
	selectedDesc  lipgloss.Style
	normalTitle   lipgloss.Style
	normalDesc    lipgloss.Style
	blurTitle     lipgloss.Style
	blurDesc      lipgloss.Style
}

type TabPanelTheme struct {
	activeTitle lipgloss.Style
	normalTitle lipgloss.Style
	active      lipgloss.Style
	inactive    lipgloss.Style
	blur        lipgloss.Style
}

type Theme struct {
	inputStyle lipgloss.Style
	titleStyle lipgloss.Style
	list       ListTheme
	tabPanel   TabPanelTheme
}

func DefaultTheme() Theme {
	return Theme{
		inputStyle: lipgloss.NewStyle(),
		titleStyle: lipgloss.NewStyle(),
		tabPanel: TabPanelTheme{
			activeTitle: lipgloss.
				NewStyle().
				MarginRight(2).
				PaddingLeft(1).
				Background(lipgloss.Color(colors.ActiveTab)).
				Foreground(lipgloss.Color(colors.White)),
			normalTitle: lipgloss.
				NewStyle().
				MarginRight(2).
				PaddingLeft(1).
				Background(lipgloss.Color(colors.InactiveTab)).
				Foreground(lipgloss.Color(colors.InactiveText)),
		},
		list: ListTheme{
			selectedTitle: lipgloss.NewStyle().
				PaddingLeft(2).                                // or use 0 and let the list add the marker
				Foreground(lipgloss.Color(colors.ActiveText)). // Green text
				PaddingLeft(2),
			selectedDesc: lipgloss.NewStyle().
				PaddingLeft(2). // or use 0 and let the list add the marker
				Foreground(lipgloss.Color(colors.ActiveSubduedText)).
				PaddingLeft(2).
				Italic(true),
			normalTitle: lipgloss.NewStyle().
				PaddingLeft(2).                          // or use 0 and let the list add the marker
				Foreground(lipgloss.Color(colors.Grey)), // Green text
			normalDesc: lipgloss.NewStyle().
				PaddingLeft(2). // or use 0 and let the list add the marker
				Foreground(lipgloss.Color(colors.DarkGrey)).
				Italic(true),
			blurTitle: lipgloss.NewStyle().
				PaddingLeft(2). // or use 0 and let the list add the marker
				Foreground(lipgloss.Color(colors.DarkGrey)),
			blurDesc: lipgloss.NewStyle().
				PaddingLeft(2).                                 // or use 0 and let the list add the marker
				Foreground(lipgloss.Color(colors.DarkestGrey)), // Green text
		},
	}
}
