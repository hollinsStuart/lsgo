package fileops

import (
	"fmt"
	"os"
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

func GetFiles(path string) ([]FileEntry, error) {
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
