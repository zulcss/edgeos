package cmd

import (
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"

	"github.com/zulcss/edgeos/pkg/image"
)

var (
	DockerBuildDir	string
)

var BuildBaseCommand = &cobra.Command{
	Use:   "base",
	Short: "Build an image",
	Long:  "",
	Run: func(BuildBaseCommand *cobra.Command, args []string) {
		log.Info("build base image")
		image.Build(DockerBuildDir)
	},
}

func init() {
	BuildCommand.AddCommand(BuildBaseCommand)
	BuildCommand.PersistentFlags().StringVarP(&DockerBuildDir, "base-dir", "", "images", "/usr/lib/stx/images")
}
