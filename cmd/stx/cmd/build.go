package cmd

import (
	"fmt"
//	"os"
//	"io"
	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
	//"github.com/zulcss/edgeos/pkg/image"
)

var (
	BaseImage	bool
	verbose		string
)

var BuildCommand = &cobra.Command{
	Use:   "build",
	Short: "Build an image",
	Long:  "",
	//Args:  cobra.MatchAll(cobra.ExactArgs(2)),
	Run: func(BuildCommand *cobra.Command, args []string) {
		/*if args[0] == "image" || args[0] == "iso" {
			image.Build(args[0], args[1])
		}*/
		fmt.Println(BaseImage)
	},
}

func init() {
	BuildCommand.PersistentFlags().BoolVarP(&BaseImage, "base", "", false, "build base starlingx image")
	BuildCommand.PersistentFlags().StringVarP(&verbose, "verbose", "v", log.WarnLevel.String(), 
		"Log level (debug: (debug, info, warn, error, fatal, panic")
     
	rootCmd.AddCommand(BuildCommand)
}

/*
func setUpLogs(out io.Writer, level string) error {
	log.SetLevel(out)
	lvl, err := log.ParseLevel(level)
	if err != nil {
		return nil
	}
	log.SetLevel(lvl)
	return nil
}
*/
