package main

import (
	"encoding/json"
	"fmt"
	"github.com/fatih/color"
	"github.com/hollinsStuart/lsgo/table"
	"github.com/hollinsStuart/lsgo/types"
	flag "github.com/spf13/pflag"
	"os"
	"path/filepath"
)

func getFiles(path string) ([]types.FileEntry, error) {
	var data []types.FileEntry

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		eType := types.File
		if entry.IsDir() {
			eType = types.Dir
		}
		modified := info.ModTime().Format("Mon Jan 2 2006")
		data = append(data, types.FileEntry{
			Name:     entry.Name(),
			EType:    eType,
			LenBytes: info.Size(),
			Modified: modified,
		})
	}
	return data, nil
}

func main() {
	var jsonOutput bool
	var recursive bool
	var human bool

	flag.BoolVarP(&jsonOutput, "json", "j", false, "output as JSON")
	flag.BoolVarP(&recursive, "recursive", "r", false, "recursively list")
	flag.BoolVarP(&human, "human", "h", false, "human-readable sizes")

	flag.Parse()

	fmt.Println("json:", jsonOutput)
	fmt.Println("recursive:", recursive)
	fmt.Println("human:", human)

	path := "."
	if flag.NArg() > 0 {
		path = flag.Arg(0)
	}

	absPath, err := filepath.Abs(path)
	if err != nil {
		color.Red("Error resolving path: %v", err)
		return
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		color.Red("Path does not exist!")
		return
	}

	files, err := getFiles(absPath)
	if err != nil {
		color.Red("Error reading directory: %v", err)
		return
	}

	if jsonOutput {
		jsonData, err := json.MarshalIndent(files, "", "  ")
		if err != nil {
			color.Red("JSON serialization error: %v", err)
			return
		}
		fmt.Println(string(jsonData))
	} else {
		table.PrintTable(files)
	}
}
