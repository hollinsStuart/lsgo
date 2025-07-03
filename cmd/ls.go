package cmd

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/hollinsStuart/lsgo/fileops"
	"github.com/hollinsStuart/lsgo/table"
	"github.com/spf13/cobra"
	"os"
	"path/filepath"
	"sort"
)

func Run(cmd *cobra.Command, args []string) {
	fmt.Printf("oneline: %v, long: %v, all: %v\n", oneLine, longList, allFiles)

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

	files, err := fileops.GetFiles(absPath)
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
		sort.Slice(files, func(i, j int) bool {
			if files[i].EType == fileops.Dir && files[j].EType != fileops.Dir {
				return true
			}
			if files[i].EType != fileops.Dir && files[j].EType == fileops.Dir {
				return false
			}
			return files[i].Name < files[j].Name
		})
		if tableOutput {
			table.PrintTable(files)
		} else {
			// TODO: print as list
		}

	}
}
