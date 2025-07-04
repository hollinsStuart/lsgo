package cmd

import (
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"sort"

	"github.com/fatih/color"
	"github.com/hollinsStuart/lsgo/fileops"
	"github.com/hollinsStuart/lsgo/icons"
	"github.com/hollinsStuart/lsgo/output"
	"github.com/spf13/cobra"
)

func Run(cmd *cobra.Command, args []string) {

	// get a path from positional arg
	if len(args) > 0 {
		path = args[0]
	}

	// resolve to absolute
	absPath, err := filepath.Abs(path)
	if err != nil {
		color.Red("Error resolving path: %v", err)
		return
	}

	if _, err := os.Stat(absPath); os.IsNotExist(err) {
		color.Red("Path does not exist!")
		return
	}

	// load files
	files, err := fileops.GetFiles(absPath)
	if err != nil {
		color.Red("Error reading directory: %v", err)
		return
	}

	// filter dot files
	if !allFiles {
		files = fileops.FilterDotFiles(files)
	}

	// json?
	if jsonOutput {
		jsonData, err := json.MarshalIndent(files, "", "  ")
		if err != nil {
			color.Red("JSON serialization error: %v", err)
			return
		}
		fmt.Println(string(jsonData))
		return
	}

	// sort dirs first
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
		output.PrintTable(files)
		return
	}

	if longOutput {
		output.PrintLong(files)
	} else if oneLineOutput {
		for _, f := range files {
			icon := icons.NerdIconForFile(f.Name, f.EType == fileops.Dir)
			fmt.Printf("%s %s\n", icon, f.Name)
		}
	} else {
		output.PrintDefault(files)
	}
	fmt.Println(color.New(color.Bold).Sprint("Bold text"))
	fmt.Println(color.New(color.Italic).Sprint("Italic text"))
	fmt.Println(color.New(color.Underline).Sprint("Underlined text"))
}
