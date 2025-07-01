package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"github.com/fatih/color"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
	"os"
	"path/filepath"
)

type EntryType string

const (
	File EntryType = "File"
	Dir  EntryType = "Dir"
)

type FileEntry struct {
	Name     string    `json:"name"`
	EType    EntryType `json:"type"`
	LenBytes int64     `json:"bytes"`
	Modified string    `json:"modified"`
}

func getFiles(path string) ([]FileEntry, error) {
	var data []FileEntry

	entries, err := os.ReadDir(path)
	if err != nil {
		return nil, err
	}

	for _, entry := range entries {
		info, err := entry.Info()
		if err != nil {
			continue
		}
		eType := File
		if entry.IsDir() {
			eType = Dir
		}
		modified := info.ModTime().Format("Mon Jan 2 2006")
		data = append(data, FileEntry{
			Name:     entry.Name(),
			EType:    eType,
			LenBytes: info.Size(),
			Modified: modified,
		})
	}
	return data, nil
}

func printTable(files []FileEntry) {
	data := make([][]string, len(files))
	for i, f := range files {
		data[i] = []string{
			f.Name,
			string(f.EType),
			fmt.Sprintf("%d", f.LenBytes),
			f.Modified,
		}
	}

	colorCfg := renderer.ColorizedConfig{
		Header: renderer.Tint{
			FG: renderer.Colors{color.FgGreen, color.Bold},
			BG: renderer.Colors{color.BgHiWhite},
		},
		Column: renderer.Tint{
			FG: renderer.Colors{color.FgCyan}, // default
			Columns: []renderer.Tint{
				{FG: renderer.Colors{color.FgMagenta}},  // Name
				{},                                      // Type
				{FG: renderer.Colors{color.FgHiYellow}}, // Bytes
				{FG: renderer.Colors{color.FgHiRed}},    // Modified
			},
		},
		Footer: renderer.Tint{
			FG: renderer.Colors{color.FgYellow, color.Bold},
		},
		Border:    renderer.Tint{FG: renderer.Colors{color.FgWhite}},
		Separator: renderer.Tint{FG: renderer.Colors{color.FgWhite}},
	}

	table := tablewriter.NewTable(os.Stdout,
		tablewriter.WithRenderer(renderer.NewColorized(colorCfg)),
		tablewriter.WithConfig(tablewriter.Config{
			Row: tw.CellConfig{
				Formatting:   tw.CellFormatting{AutoWrap: tw.WrapNormal},
				Alignment:    tw.CellAlignment{Global: tw.AlignLeft},
				ColMaxWidths: tw.CellWidth{Global: 25},
			},
			Footer: tw.CellConfig{
				Alignment: tw.CellAlignment{Global: tw.AlignRight},
			},
		}),
	)

	table.Header([]string{"Name", "Type", "Bytes", "Modified"})
	err := table.Bulk(data)
	if err != nil {
		return
	}
	table.Footer([]string{"", "Total Files", fmt.Sprintf("%d", len(files)), ""})
	err = table.Render()
	if err != nil {
		return
	}
}

func nerdIconForFile(name string, isDir bool) string {
	if isDir {
		if name == ".git" {
			return "" // special git dir
		}
		return "" // generic folder
	}

	// handle special filenames first
	switch name {
	case "Makefile", "makefile":
		return "" // icon for make
	case "CMakeLists.txt":
		return "" // generic config icon
	case ".gitignore":
		return ""
	}

	// then by extension
	switch filepath.Ext(name) {
	case ".go", ".mod", ".sum":
		return "" // Go
	case ".rs":
		return "" // Rust
	case ".py":
		return "" // Python
	case ".c":
		return "" // C
	case ".h", ".hpp":
		return "" // header
	case ".cpp", ".cc", ".cxx":
		return "" // C++
	case ".md":
		return "󰂺"
	case ".txt":
		return "" // text
	default:
		return "" // generic file
	}
}

func iconForFile(name string, isDir bool) string {
	if isDir {
		return "📁"
	}
	// Extensions
	switch filepath.Ext(name) {
	case ".go", ".rs", ".py", ".js", ".ts", ".cpp", ".c", ".h":
		return "🔧"
	case ".md", ".txt":
		return "📝"
	case ".zip", ".tar", ".gz", ".rar":
		return "📦"
	case ".png", ".jpg", ".jpeg", ".gif", ".svg":
		return "🖼"
	case ".mp3", ".wav", ".flac":
		return "🎵"
	case ".mp4", ".mkv", ".webm":
		return "🎬"
	case ".db", ".sqlite":
		return "🗃"
	default:
		return "📄"
	}
}

func main() {
	var jsonOutput bool
	flag.BoolVar(&jsonOutput, "json", false, "output in JSON format")
	flag.Parse()

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
		printTable(files)
	}
}
