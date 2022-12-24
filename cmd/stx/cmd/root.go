package cmd

import (
	"os"

	"github.com/spf13/cobra"
	log "github.com/sirupsen/logrus"
)

var (
	debug	bool
)


var rootCmd = &cobra.Command{
	Use:   "stx",
	Short: "StarlingX cli",
	Long: "",
	Run: func(cmd *cobra.Command, args []string) { },
}

func Execute() {
	err := rootCmd.Execute()
	if err != nil {
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().BoolVarP(&debug, "debug", "d", false, "Enable debug output")

}

func initConfig() {
	log.SetOutput(os.Stdout)
	if debug {
		log.SetLevel(log.DebugLevel)
	}
}

