package cmd

import (
	"github.com/spf13/cobra"
	"os"
)

var (
	oneLine     bool
	longList    bool
	allFiles    bool
	jsonOutput  bool
	tableOutput bool
)

var rootCmd = &cobra.Command{
	Use:   "lsgo",
	Short: "lsgo is a colorful ls replacement",
	Long:  `lsgo is like ls/eza, with icons, tables, human sizes, and more.`,
	Run:   Run,
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.Flags().BoolVarP(&oneLine, "oneline", "1", false, "display one entry per line")
	rootCmd.Flags().BoolVarP(&longList, "long", "l", false, "display extended file metadata as a table")
	rootCmd.Flags().BoolVarP(&allFiles, "all", "a", false, "show hidden and 'dot' files")
	rootCmd.Flags().BoolVarP(&jsonOutput, "json", "j", false, "output as json")
	rootCmd.Flags().BoolVarP(&tableOutput, "table", "t", false, "output as table")
}
