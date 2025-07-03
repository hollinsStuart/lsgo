package output

import (
	"fmt"
	"github.com/fatih/color"
	"github.com/hollinsStuart/lsgo/fileops"
	"github.com/hollinsStuart/lsgo/icons"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
	"os"
)

func PrintTable(files []fileops.FileEntry) {
	data := make([][]string, len(files))
	for i, f := range files {
		icon := icons.NerdIconForFile(f.Name, f.EType == fileops.Dir)
		data[i] = []string{
			fmt.Sprintf("%s %s", icon, f.Name),
			string(f.EType),
			fileops.HumanBytes(f.LenBytes),
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

func PrintLong(files []fileops.FileEntry) {
	grey := color.New(color.FgHiBlack)
	for _, f := range files {
		perm := permissionsString(f.Mode)
		size := fileops.HumanBytes(f.LenBytes)
		if f.EType == fileops.Dir {
			size = grey.Sprint("   \u001B[1m-\u001B[0m")
		}
		links := fmt.Sprintf("%d", f.NumLinks)
		icon := icons.NerdIconForFile(f.Name, f.EType == fileops.Dir)

		// Example like eza/ls -l:
		// drwxr-xr-x  3 user  group  4.0k Jul  3 15:04   cmd
		fmt.Printf("\u001B[1m%s\u001B[0m %3s %-8s %-8s %4s %s %s %s\n",
			color.New(color.FgHiWhite).Sprint(perm),
			links,
			color.New(color.FgYellow).Sprint(f.Owner),
			color.New(color.FgCyan).Sprint(f.Group),
			size,
			color.New(color.FgHiMagenta).Sprint(f.Modified),
			icon,
			f.Name,
		)
	}

}

// permissionsString converts FileMode to string like drwxr-xr-x
func permissionsString(mode os.FileMode) string {
	s := ""
	switch {
	case mode.IsDir():
		s += "d"
	case mode&os.ModeSymlink != 0:
		s += "l"
	default:
		s += "-"
	}
	perms := []os.FileMode{
		0400, 0200, 0100, // owner
		0040, 0020, 0010, // group
		0004, 0002, 0001, // other
	}
	for i, p := range perms {
		if mode&p != 0 {
			switch i % 3 {
			case 0:
				s += "r"
			case 1:
				s += "w"
			case 2:
				s += "x"
			}
		} else {
			s += "-"
		}
	}
	return s
}
