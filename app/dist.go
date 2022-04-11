package app

import (
	"os"

	"github.com/axllent/golp/utils"
)

// DeleteDistDirs deleted all dist folders
func DeleteDistDirs() error {
	for _, d := range Conf.CleanDirs {
		if utils.IsDir(d) {
			if err := os.RemoveAll(d); err != nil {
				return err
			}

			Log().Debugf("Deleted dist directory: %s", d)
		}
	}

	for _, p := range Conf.Process {
		if p.Dist != "" {
			if utils.IsDir(p.Dist) {
				if err := os.RemoveAll(p.Dist); err != nil {
					return err
				}

				Log().Debugf("Deleted dist directory: %s", p.Dist)
			}
		}
	}

	return nil
}
