package root

import (
	"encoding/json"
	"log"
	"os"
)

func loadData(filepath string, buffer any) {
	data, err := os.ReadFile(filepath)
	if err != nil {
		println("There was an error loading data")
		panic(err)
	}
	err = json.Unmarshal(data, buffer)
	if err != nil {
		println("%s", filepath)
		log.Fatal("Couldn't read JSON file")
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
