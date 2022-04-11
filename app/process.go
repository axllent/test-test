package app

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/axllent/golp/utils"
	"github.com/evanw/esbuild/pkg/api"
	"github.com/wellington/go-libsass"
)

var processTypes = map[string]bool{"styles": true, "scripts": true, "copy": true}

// Process will process the ProcessStruct
func (p ProcessStruct) Process() error {
	switch p.Type {
	case "styles":
		// return nil
		return p.processStyles()
	case "scripts":
		// return nil
		return p.processScripts()
	case "copy":
		return p.processCopy()
	}

	return fmt.Errorf("Unknown process type: %s", p.Type)
}

func (p ProcessStruct) processStyles() error {
	sw := utils.StartTimer()

	files := p.Files()

	for f, subDir := range files {
		filename := path.Base(f)
		d := path.Join(p.Dist, subDir)
		if !utils.IsDir(d) {
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}

		out := path.Join(d, filename)
		extension := strings.ToLower(filepath.Ext(filename))

		if extension == ".scss" || extension == ".sass" || extension == ".css" {
			out = out[0:len(out)-len(extension)] + ".css"

			wi, err := os.Create(out)
			if err != nil {
				return err
			}
			defer wi.Close()

			fi, err := os.Open(f)
			if err != nil {
				return err
			}
			defer fi.Close()

			comp, err := libsass.New(wi, fi)
			if err != nil {
				return err
			}

			comp.Option(libsass.Path(f)) // path must be set for relative imports

			if Minify {
				comp.Option(libsass.OutputStyle(3)) // compress CSS
			} else {
				comp.Option(libsass.SourceMap(true, out+".map", "")) // add sourcemaps
			}

			if err := comp.Run(); err != nil {
				return err
			}

			Log().Debugf("processed %s to %s", f, out)
		} else {
			Log().Warningf("unsupported extension: %s", f)
		}

	}

	Log().Infof("'styles' compiled in %v", sw.Elapsed())
	return nil
}

func (p ProcessStruct) processScripts() error {
	sw := utils.StartTimer()

	files := p.Files()

	for f, subDir := range files {
		filename := path.Base(f)
		d := path.Join(p.Dist, subDir)
		if !utils.IsDir(d) {
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}
		out := path.Join(d, filename)

		options := api.BuildOptions{
			EntryPoints:    []string{f},
			Outfile:        out,
			Write:          true,
			AllowOverwrite: true,
		}

		if Minify {
			options.MinifyWhitespace = true
			options.MinifyIdentifiers = true
			options.MinifySyntax = true
		}

		result := api.Build(options)

		if len(result.Errors) > 0 {
			errorMsg := fmt.Sprintf("> Error %s:%d\n%s",
				result.Errors[0].Location.File,
				result.Errors[0].Location.Line,
				result.Errors[0].Text,
			)

			return fmt.Errorf("%s", errorMsg)
		}

		Log().Debugf("processed %s to %s", f, out)
	}

	Log().Infof("'scripts' compiled in %v", sw.Elapsed())

	return nil
}

func (p ProcessStruct) processCopy() error {
	sw := utils.StartTimer()

	files := p.Files()

	for f, subDir := range files {
		filename := path.Base(f)
		d := path.Join(p.Dist, subDir)
		if !utils.IsDir(d) {
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}
		out := path.Join(d, filename)

		if err := utils.Copy(f, out); err != nil {
			return err
		}

		Log().Debugf("copied %s to %s", f, out)
	}

	Log().Infof("'copy' completed in %v", sw.Elapsed())

	return nil
}
