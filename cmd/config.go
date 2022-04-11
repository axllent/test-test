package cmd

import (
	"fmt"

	"github.com/spf13/cobra"
)

// configCmd represents the config command
var configCmd = &cobra.Command{
	Use:   "config",
	Short: "View an example config file",
	Long: `View an example config file.
	
This file is typically saved in your project root directory as golp.yaml

All styles, scripts, and assets are relative to this file.`,
	Args: cobra.ExactArgs(0),
	Run: func(cmd *cobra.Command, args []string) {
		fmt.Println(`# optional - will be deleted on build / clean
clean: 
  - themes/site/dist

# SASS & CSS files
styles:
  - src:
      - themes/site/src/sass/*.scss
      - themes/site/src/sass/**.css
      - node_modules/@dashboardcode/bsmultiselect/dist/css/BsMultiSelect.css
    dist:  themes/site/dist/css

# JavaScript files
scripts:
  - src:
      - node_modules/@popperjs/core/dist/umd/popper.min.js
      - node_modules/bootstrap/dist/js/bootstrap.min.js
      - node_modules/axios/dist/axios.min.js
      - node_modules/@dashboardcode/bsmultiselect/dist/js/BsMultiSelect.min.js
      - node_modules/vuedraggable/dist/vuedraggable.umd.min.js
      - node_modules/sortablejs/Sortable.min.js
      - node_modules/vue/dist/vue.global.prod.js
      - themes/site/src/js/**.js
    dist: themes/site/dist/js

# All other files
copy:
  - src:
      - themes/site/src/images/**
    dist: themes/site/dist/images
  - src: 
      - themes/site/src/fonts/**
    dist: themes/site/dist/fonts/`)
	},
}

func init() {
	rootCmd.AddCommand(configCmd)
}
