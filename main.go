package main

import (
	"fmt"
	"os"
	app "sideDesert/ideasv2/root"

	tea "github.com/charmbracelet/bubbletea"
)

func main() {

	p := tea.NewProgram(app.InitialModel(), tea.WithAltScreen())

	if _, err := p.Run(); err != nil {
		fmt.Printf("Alas, there's been an error: %v", err)
		os.Exit(1)
	}

}
