package cmd

import (
	"os"

	"github.com/axllent/golp/app"
	"github.com/axllent/golp/utils"
	"github.com/spf13/cobra"
)

// buildCmd represents the build command
var buildCmd = &cobra.Command{
	Use:   "build",
	Short: "Compile & copy your assets (single)",
	Long: `Compile & copy your assets in a single run.

By default SASS & JS files will include SourceMaps, which are used by browsers
to debug your code. Run with '-m' to disable this and minify the output.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {

		if err := app.ParseConfig(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}

		sw := utils.StartTimer()

		if err := app.DeleteDistDirs(); err != nil {
			app.Log().Error(err.Error())
			os.Exit(1)
		}

		for _, p := range app.Conf.Process {
			if err := p.Process(); err != nil {
				app.Log().Error(err.Error())
				os.Exit(1)
			}
		}

		app.Log().Infof("completed in %v", sw.Elapsed())
	},
}

func init() {
	rootCmd.AddCommand(buildCmd)

	buildCmd.Flags().BoolVarP(&app.Minify, "minify", "m", false, "minify dist styles & scripts")
}
