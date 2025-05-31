package root

import (
	"os"
	"path/filepath"
)

const (
	Ideas = iota
	Books
)

const (
	Read = iota
	Write
	Edit
	Delete
)

var (
	home, _        = os.UserHomeDir()
	folderPath     = filepath.Join(home, "ideas")
	ideasFolder    = filepath.Join(folderPath, "ideas")
	booksFolder    = filepath.Join(folderPath, "books")
	projectsFolder = filepath.Join(folderPath, "projects")
	editor         = "nvim"
)
