package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	oneLineOutput bool
	longOutput    bool
	allFiles      bool
	jsonOutput    bool
	tableOutput   bool
	path          string
)

var rootCmd = &cobra.Command{
	Use:   "lsgo",
	Short: "lsgo is a colorful ls replacement",
	Long:  `lsgo is like ls/eza, with icons, tables, human sizes, and more.`,
	Run:   Run,
}

func Execute() {
	if err := rootCmd.Execute(); err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&oneLineOutput, "oneline", "1", false, "display one entry per line")
	rootCmd.Flags().BoolVarP(&longOutput, "long", "l", false, "display extended file metadata as a table")
	rootCmd.Flags().BoolVarP(&allFiles, "all", "a", false, "show hidden and 'dot' files")
	rootCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "output as json")
	rootCmd.Flags().BoolVarP(&tableOutput, "table", "t", false, "output as table")
	rootCmd.Flags().StringVarP(&path, "path", "p", ".", "path to list")
}
