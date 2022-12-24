package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	//"github.com/zulcss/edgeos/pkg/image"
)

var (
	BaseImage	bool
)

var BuildCommand = &cobra.Command{
	Use:   "build",
	Short: "Build an image",
	Long:  "",
	Run: func(BuildCommand *cobra.Command, args []string) {
		log.Info("test")
	},
}

func init() {
	BuildCommand.PersistentFlags().BoolVarP(&BaseImage, "base", "", false, "build base starlingx image")     
	rootCmd.AddCommand(BuildCommand)
}
