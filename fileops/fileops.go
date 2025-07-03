package fileops

import "fmt"

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

func HumanBytes(bytes int64) string {
	const unit = 1000 // Use 1024 for binary (KiB, MiB, etc)
	if bytes < unit {
		return fmt.Sprintf("%d", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit; n /= unit {
		div *= unit
		exp++
	}
	return fmt.Sprintf("%.1f%c", float64(bytes)/float64(div), "kMGTPE"[exp])
}
