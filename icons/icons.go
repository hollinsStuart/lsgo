package icons

import "path/filepath"

func NerdIconForFile(name string, isDir bool) string {
	if isDir {
		if name == ".git" {
			return "" // git dir
		}
		return "" // folder
	}

	// exact matches
	switch name {
	case "Makefile", "makefile":
		return "" // make
	case "CMakeLists.txt":
		return "" // config
	case ".gitignore":
		return ""
	case "Dockerfile":
		return ""
	}

	// by extension
	switch filepath.Ext(name) {
	case ".go", ".mod", ".sum":
		return "" // Go
	case ".rs":
		return "" // Rust
	case ".py":
		return "" // Python
	case ".lua":
		return "" // Lua
	case ".c":
		return "" // C
	case ".h", ".hpp":
		return "" // C header
	case ".cpp", ".cc", ".cxx":
		return "" // C++
	case ".js":
		return "" // JavaScript
	case ".ts":
		return "" // TypeScript
	case ".jsx", ".tsx":
		return "" // React JSX, TSX
	case ".java":
		return "" // Java
	case ".kt", ".kts":
		return "" // Kotlin
	case ".rb":
		return "" // Ruby
	case ".php":
		return "" // PHP
	case ".html", ".htm":
		return "" // HTML
	case ".css":
		return "" // CSS
	case ".scss", ".sass":
		return "" // SCSS
	case ".json":
		return "" // JSON
	case ".yaml", ".yml":
		return "" // YAML
	case ".sh", ".bash":
		return "" // Shell
	case ".md":
		return "󰂺"
	case ".txt":
		return ""
	default:
		return "" // generic file
	}
}
