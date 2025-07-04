package output

import (
	"fmt"
	"os"

	"github.com/fatih/color"
	"github.com/hollinsStuart/lsgo/fileops"
	"github.com/hollinsStuart/lsgo/icons"
	"github.com/olekukonko/tablewriter"
	"github.com/olekukonko/tablewriter/renderer"
	"github.com/olekukonko/tablewriter/tw"
)

func PrintDefault(files []fileops.FileEntry) {
	for _, f := range files {
		icon := icons.NerdIconForFile(f.Name, f.EType == fileops.Dir)
		if f.EType == fileops.Dir {
			fmt.Printf("%s  ", folderString(icon, f.Name))
		} else {
			fmt.Printf("%s %s  ", icon, f.Name)
		}
	}
	fmt.Println()
}

// @return returns a folder string with no spaces
func folderString(icon, name string) string {
	blue := color.New(color.FgBlue)
	fileString := blue.Add(color.Bold).Sprintf("%s %s", icon, name)
	return fileString
}

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
			size = grey.Add(color.Bold).Sprint("   -")
		}
		icon := icons.NerdIconForFile(f.Name, f.EType == fileops.Dir)

		// Example like eza/ls -l:
		// drwxr-xr-x    - hollins staff 3 Jul 21:43  cmd
		fmt.Printf("%s %4s %-8s %-8s %s %s %s\n",
			perm,
			size,
			color.New(color.FgYellow).Add(color.Bold).Sprint(f.Owner),
			color.New(color.FgCyan).Sprint(f.Group),
			color.New(color.FgHiMagenta).Sprint(f.Modified),
			icon,
			f.Name,
		)
	}
}

// permissionsString converts FileMode to string like drwxr-xr-x
func permissionsString(mode os.FileMode) string {
	yellow := color.New(color.FgYellow)
	blue := color.New(color.FgBlue)
	green := color.New(color.FgGreen)
	red := color.New(color.FgRed)
	grey := color.New(color.FgHiBlack)

	// HACK: MUST define regular chars first
	// normal versions
	rCharRegular := yellow.Sprint("r")
	wCharRegular := red.Sprint("w")
	xCharRegular := green.Sprint("x")

	// use bold versions
	rCharBold := yellow.Add(color.Bold).Sprint("r")
	wCharBold := red.Add(color.Bold).Sprint("w")
	xCharBold := green.Add(color.Bold).Sprint("x")

	dashChar := grey.Add(color.Bold).Sprint("-")
	lChar := blue.Sprint("l")
	dChar := blue.Add(color.Bold).Sprint("d")

	s := ""
	switch {
	case mode.IsDir():
		s += dChar
	case mode&os.ModeSymlink != 0:
		// TODO: we will need to handle symlinks properly
		s += lChar
	default:
		s += "."
	}
	perms := []os.FileMode{
		0400, 0200, 0100, // owner
		0040, 0020, 0010, // group
		0004, 0002, 0001, // other
	}
	for i, p := range perms {
		var out string
		if mode&p != 0 {
			switch i % 3 {
			case 0:
				out = rCharRegular
			case 1:
				out = wCharRegular
			case 2:
				out = xCharRegular
			}
		} else {
			out = dashChar
		}

		// make the first group (owner) bold
		if i < 3 {
			switch i % 3 {
			case 0:
				out = rCharBold
			case 1:
				out = wCharBold
			case 2:
				out = xCharBold
			}
		}

		s += out
	}
	return s
}
