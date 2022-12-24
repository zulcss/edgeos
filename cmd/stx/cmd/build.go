package cmd

import (
	"github.com/spf13/cobra"
)

var BuildCommand = &cobra.Command{
	Use:   "build",
	Short: "Build an image",
	Long:  "",
}

func init() {
	rootCmd.AddCommand(BuildCommand)
}
