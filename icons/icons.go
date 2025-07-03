package icons

import "path/filepath"

func NerdIconForFile(name string, isDir bool) string {
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

func RichIconForFile(name string, isDir bool) string {
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
