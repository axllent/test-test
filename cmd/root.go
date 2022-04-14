package cmd

import (
	"fmt"
	"os"

	"github.com/axllent/golp/app"
	"github.com/spf13/cobra"
)

var cfgFile string

// rootCmd represents the base command when called without any subcommands
var rootCmd = &cobra.Command{
	Use:   "golp",
	Short: "Golp automates build workflows.",
	Long: `Golp automates build workflows, compiling SASS and JavaScript into configurable
"dist" directories. It also handles dynamic copying of static assets.

Golp is not a Gulp drop-in replacement, but aims to solve many of the same
problems that Gulp does. It is fast, simple, and runs from a single binary.

Documentation:
  https://github.com/axllent/golp`,
	// Uncomment the following line if your bare application
	// has an action associated with it:
	//	Run: func(cmd *cobra.Command, args []string) { },
}

// Execute adds all child commands to the root command and sets flags appropriately.
// This is called by main.main(). It only needs to happen once to the rootCmd.
func Execute() {
	if err := rootCmd.Execute(); err != nil {
		fmt.Println(err)
		os.Exit(1)
	}
}

func init() {
	rootCmd.PersistentFlags().StringVarP(&app.Conf.ConfigFile, "config", "c", "./golp.yaml", "config file")
	rootCmd.PersistentFlags().BoolVarP(&app.VerboseLogging, "verbose", "v", false, "verbose logging")

	rootCmd.CompletionOptions.HiddenDefaultCmd = true
	rootCmd.SetHelpCommand(&cobra.Command{Hidden: true})
}
