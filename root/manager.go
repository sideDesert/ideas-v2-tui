package root

import (
	"fmt"
	"log"
	"os"

	tea "github.com/charmbracelet/bubbletea"
	"github.com/charmbracelet/glamour"
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

func (m *Manager) RemoveItem(index int) {
	n := len(m.List.Items())
	if index >= n {
		return
	}
	if item, ok := m.List.Items()[index].(ListItem); ok {
		fp := item.FilePath
		if err := os.Remove(fp); err != nil {
			log.Println("Error[*Manager.RemoveItem]: Could not delete file ", fp, err)
		}
		m.List.RemoveItem(index)
	}
}

func (m *Manager) GetActiveFilepath() string {
	n := len(m.List.Items())
	if n == 0 {
		return ""
	}

	if item, ok := m.List.Items()[m.List.Index()].(ListItem); ok {
		return item.FilePath
	}

	return ""
}

func (m *model) editDescription() tea.Cmd {
	return tea.ExecProcess(editorCmd(m.getManager().GetActiveFilepath()), nil)
}

func (m *Manager) GetActiveListItem() *ListItem {
	list := m.List
	n := len(list.Items())
	index := list.Index()
	if n == 0 {
		return nil
	}
	if listItem, ok := list.Items()[index].(ListItem); ok {
		return &listItem
	}
	return nil
}

func (m *Manager) SetViewportContent(content string) {
	out, err := glamour.Render(content, "dark")
	if err != nil {
		fmt.Println("Error[*Manager.SetViewportContent]:", err)
		return
	}
	m.Viewport.SetContent(out)
}
