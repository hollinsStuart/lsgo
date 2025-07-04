package fileops

import (
	"fmt"
	"os"
	"os/user"
	"strings"
	"syscall"
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

	Mode     os.FileMode `json:"-"`
	Owner    string      `json:"owner"`
	Group    string      `json:"group"`
	NumLinks uint64      `json:"links"`
}

func HumanBytes(bytes int64) string {
	const unit = 1000
	if bytes < unit {
		return fmt.Sprintf("%4d", bytes)
	}
	div, exp := int64(unit), 0
	for n := bytes / unit; n >= unit && exp < 5; n /= unit {
		div *= unit
		exp++
	}

	value := float64(bytes) / float64(div)
	if value >= 10 {
		// like 123k
		return fmt.Sprintf("%3.0f%c", value, "kMGTPE"[exp])
	} else {
		// like 1.2k
		return fmt.Sprintf("%3.1f%c", value, "kMGTPE"[exp])
	}
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

		modified := info.ModTime().Format("2 Jan 15:04")

		// Try to get Unix owner/group/nlink info
		var owner, group string
		var numLinks uint64

		if stat, ok := info.Sys().(*syscall.Stat_t); ok {
			numLinks = uint64(stat.Nlink)

			uid := fmt.Sprint(stat.Uid)
			if u, err := user.LookupId(uid); err == nil {
				owner = u.Username
			} else {
				owner = uid
			}

			gid := fmt.Sprint(stat.Gid)
			if g, err := user.LookupGroupId(gid); err == nil {
				group = g.Name
			} else {
				group = gid
			}
		}

		data = append(data, FileEntry{
			Name:     entry.Name(),
			EType:    eType,
			LenBytes: info.Size(),
			Modified: modified,
			Mode:     info.Mode(),
			Owner:    owner,
			Group:    group,
			NumLinks: numLinks,
		})
	}
	return data, nil
}

// FilterDotFiles filters out entries that start with '.'
// If showAll is true (like -a), it returns all files unchanged.
func FilterDotFiles(files []FileEntry) []FileEntry {
	var filtered []FileEntry
	for _, file := range files {
		if strings.HasPrefix(file.Name, ".") {
			continue
		}
		filtered = append(filtered, file)
	}
	return filtered
}

func RichFileName(file FileEntry) string {
	if file.EType == Dir {

	} else if file.EType == File {

	} else {
		// Something went wrong
	}
	return ""
}
