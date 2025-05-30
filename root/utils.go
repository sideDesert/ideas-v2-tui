package root

import (
	"fmt"
	"log"
	"os"
	"path/filepath"
	"strings"
)

type Data struct {
	filePath string
	title    string
	hunk     string
}

func loadProjectData() []Project {
	projects := make([]Project, 0)
	data, err := loadData(projectsFolder)
	if err != nil {
		log.Println("Error [loadIdeaData]: ", err)
	}
	for _, d := range data {
		i := Project{
			TitleText:       d.title,
			FilePath:        d.filePath,
			DescriptionText: d.hunk,
		}
		projects = append(projects, i)
	}
	return projects
}

func loadBookData() []Book {
	books := make([]Book, 0)
	data, err := loadData(booksFolder)
	if err != nil {
		log.Println("Error [loadIdeaData]: ", err)
	}
	for _, d := range data {
		i := Book{
			TitleText:       d.title,
			FilePath:        d.filePath,
			DescriptionText: d.hunk,
		}
		books = append(books, i)
	}
	return books
}

func loadIdeaData() []Idea {
	ideas := make([]Idea, 0)
	data, err := loadData(ideasFolder)
	if err != nil {
		log.Println("Error [loadIdeaData]: ", err)
	}
	for _, d := range data {
		i := Idea{
			TitleText:       d.title,
			FilePath:        d.filePath,
			DescriptionText: d.hunk,
		}
		ideas = append(ideas, i)
	}
	return ideas
}

func loadData(dirPath string) ([]Data, error) {
	dataBuffer := make([]Data, 0)
	entries, err := os.ReadDir(dirPath)
	if err != nil {
		return dataBuffer, err
	}

	for _, entry := range entries {
		if filename := entry.Name(); !entry.IsDir() {
			fileData, err := os.ReadFile(filepath.Join(dirPath, filename))
			if err != nil {
				log.Println("Error [loadData]: ", err)
				continue
			}

			hunk := string(fileData)
			title := ""

			// Take care of extenstion
			split := strings.Split(filename, ".")
			if len(split) != 0 {
				title = split[0]
			}

			// Take care of duplication number
			split = strings.Split(title, "_")
			if len(split) != 0 {
				title = split[0]
			}
			fp := filepath.Join(dirPath, filename)
			dataNode := Data{
				filePath: fp,
				title:    title,
				hunk:     hunk,
			}
			dataBuffer = append(dataBuffer, dataNode)
		}
	}
	return dataBuffer, nil
}

func getUniqueFileName(dirpath string, filename string, extension string) string {
	fname := filename + "." + extension
	candidate := filepath.Join(dirpath, fname)
	// CHECK FOR EXISTENCE
	for index := 0; ; index++ {
		if _, err := os.Stat(candidate); os.IsNotExist(err) {
			return candidate // file doesn't exist, safe to use
		}
		candidate = filepath.Join(dirpath, fmt.Sprintf("%s_%d.%s", filename, index, extension))
	}
}

func get_mode(mode int) string {
	if mode == Read {
		return "R"
	} else if mode == Write {
		return "W"
	} else if mode == Edit {
		return "E"
	} else {
		return "U"
	}
}
