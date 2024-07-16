package cmd

import (
	"github.com/joosejunsheng/jhft/cmd/core"
	"github.com/joosejunsheng/jhft/version"
	"github.com/spf13/cobra"
)

func Run() {

}

var versionCmd = &cobra.Command{
	Use:   "version",
	Short: "print version.",
	Long:  "print version.",
	Args:  cobra.NoArgs,
	Run: func(cmd *cobra.Command, args []string) {
		version.Printer()
	},
}

func Execute() {
	var rootCmd = &cobra.Command{Use: "hft"}
	rootCmd.AddCommand(core.CoreCmd, versionCmd)
	rootCmd.Execute()
}
