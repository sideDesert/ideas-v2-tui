# Ideas v2
The second version of the Ideas App (originally written in Rust).
This version is written using the Charmstack in golang! I wanted to create a useful TUI for my projects and ideas, and I wanted it to be aesthetically pleasing. Hence I recreated the original Ideas app with this new framework!

A modern, extensible terminal user interface (TUI) for managing your ideas and books, built with [Bubble Tea](https://github.com/charmbracelet/bubbletea), [Huh](https://github.com/charmbracelet/huh), and [Lip Gloss](https://github.com/charmbracelet/lipgloss).

### Note: The project is very much a WIP. I am still building.

---

## Features

- **Tabbed Interface:** Switch between Ideas, Books, Projects, and Debugger tabs.
- **Panel Navigation:** Move between list and detail panels with keyboard shortcuts.
- **Add/Edit/Delete:** Quickly add, edit, or remove ideas and books.
- **Rich Forms:** Use beautiful, interactive forms for data entry.
- **Keyboard Driven:** All actions are accessible via intuitive key bindings.
- **Persistent Storage:** Data is saved to JSON files for easy backup and editing.

---

## Key Bindings

| Action         | Keys                        | Description                |
| -------------- | -------------------------- | -------------------------- |
| Move Up/Down   | `↑/k`, `↓/j`               | Navigate list              |
| Move Left/Right| `←/h`, `→/l`               | Switch panels              |
| Next Tab       | `L`, `ctrl+tab`            | Next tab                   |
| Prev Tab       | `H`, `ctrl+shift+tab`      | Previous tab               |
| Next Panel     | `ctrl+l`                   | Next panel                 |
| Prev Panel     | `ctrl+h`                   | Previous panel             |
| Add Mode       | `a`, `i`, `A`, `I`         | Add new item               |
| Edit Mode      | `c`, `e`                   | Edit selected item         |
| Read Mode      | `esc`, `ctrl+o`            | Return to read mode        |
| Delete Item    | `d`, `D`                   | Delete selected item       |
| Help           | `?`                        | Toggle help                |
| Quit           | `q`, `ctrl+c`              | Quit the application       |

---

## Data Files

- `ideas.json` — Stores your ideas.
- `books.json` — Stores your books.

Both are automatically updated as you use the app.

---

## Development

- **Main entry:** `main.go`
- **Core logic:** `root/`
- **Keymap:** `keymap/main.go`
- **Components:** `components/`
- **Styling:** Uses Lip Gloss for beautiful terminal output.

---

## Getting Started

1. **Install dependencies:**
   ```sh
   go mod tidy

2. **Run the app:**
   ```sh
   go run main.go
   ```
3. **Navigate with your keyboard and enjoy!

## Credits

- [Charmbracelet Bubble Tea](https://github.com/charmbracelet/bubbletea)
- [Charmbracelet Huh](https://github.com/charmbracelet/huh)
- [Charmbracelet Lip Gloss](https://github.com/charmbracelet/lipgloss)

## TODO
- [x] Make Forms use up same width and height adjustments
- [x] Make Form for projects
- [x] Move documents from using .JSON files to using normal files
- [x] Add Command to Open Document in NVIM
- [x] Update description after changes
- [x] Fix Delete button
- [x] Add Scroll view
- [ ] Add Rename Option
- [ ] Iterate on Project View design
- [ ] Add Upgrade to project button
- [ ] Add Line number and other beautification aspects to the markdown renderer
- [ ] Replace huh with custom forms (I don't like how the keys are being handled)
- [x] Refine Help

---

> Designed for productivity and joy in your terminal.

Let me know if you want to highlight any specific features, add screenshots, or include a section on contributing!
