package root

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/huh"
)

const (
	NullState = iota
	SaveAndExit
	Exit
)

func (m *Manager) HandleUpdateForm(msg tea.Msg) (tea.Cmd, int) {
	newForm, cmd := m.Form.Update(msg)
	state := NullState
	if form, ok := newForm.(*huh.Form); ok {
		m.Form = form
	} else {
		fmt.Println("NO form updates")
	}

	// Check if form is completed
	if m.Form.State == huh.StateCompleted {
		title := m.Form.GetString("title")
		desc := m.Form.GetString("desc")
		confirm := m.Form.GetBool("confirm")

		tryFp := getUniqueFileName(m.DirPath, title, "md")

		newListItem := ListItem{
			TitleText:       title,
			FilePath:        tryFp,
			DescriptionText: desc,
		}

		DEBUG := false
		if DEBUG {
			f, _ := os.OpenFile("debug.log", os.O_APPEND|os.O_CREATE|os.O_WRONLY, 0644)
			fmt.Fprintf(f, "DEBUG: title='%s', desc='%s', confirm=%v\n", title, desc, confirm)
			f.Close()
		}

		if confirm {
			m.List.InsertItem(len(m.List.Items()), newListItem)
			state = SaveAndExit
		} else {
			state = Exit
		}
	}

	return cmd, state
}

func (m Manager) Init() []tea.Cmd {
	cmds := make([]tea.Cmd, 0)
	cmd := m.Viewport.Init()
	cmds = append(cmds, cmd)

	cmd = m.Form.Init()
	cmds = append(cmds, cmd)

	return cmds
}

func (m *Manager) SaveLatestFile() {
	listItems := m.List.Items()
	if len(listItems) != 0 {
		if listItem, ok := listItems[len(listItems)-1].(ListItem); ok {
			fp := listItem.FilePath
			data := []byte(listItem.Description())
			err := os.WriteFile(fp, data, 0664)
			if err != nil {
				log.Println("Error [*Manager.SaveFiles()]: ", err)
			}
		}
	}
}
