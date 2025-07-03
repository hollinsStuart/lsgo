package types

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
