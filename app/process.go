package app

import (
	"fmt"
	"os"
	"path"
	"path/filepath"
	"strings"

	"github.com/axllent/golp/utils"
	"github.com/evanw/esbuild/pkg/api"
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

	if p.DistFile != "" {
		imports := []string{}
		for _, f := range files {
			extension := strings.ToLower(filepath.Ext(f.InFile))

			if extension == ".css" {
				c, err := utils.FileGetContents(f.InFile)
				if err != nil {
					return err
				}

				imports = append(imports, string(c))
			} else {
				imports = append(imports, fmt.Sprintf(`@import "%s";`, f.InFile))
			}
		}

		if !utils.IsDir(p.Dist) {
			/* #nosec G301 */
			if err := os.MkdirAll(p.Dist, 0755); err != nil {
				return err
			}
		}

		sassImport := strings.Join(imports, "\n")

		out := path.Join(p.Dist, p.DistFile)

		if err := compileStyles(sassImport, out, ""); err != nil {
			return err
		}

		Log().Debugf("processed %d SASS files to %s", len(files), out)
		Log().Infof("'%s' compiled in %v", p.Name, sw.Elapsed())

		return nil
	}

	for _, f := range files {
		filename := filepath.Base(f.InFile)
		d := path.Join(p.Dist, f.OutPath)
		if !utils.IsDir(d) {
			/* #nosec G301 */
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}

		out := path.Join(d, filename)
		extension := strings.ToLower(filepath.Ext(filename))

		if extension == ".scss" || extension == ".sass" || extension == ".css" {
			out = out[0:len(out)-len(extension)] + ".css"

			content := fmt.Sprintf(`@import "%s";`, f.InFile)

			if extension == ".css" {
				c, err := utils.FileGetContents(f.InFile)
				if err != nil {
					return err
				}

				content = c
			}

			if err := compileStyles(string(content), out, f.InFile); err != nil {
				return err
			}
		} else {
			Log().Warningf("unsupported stylesheet file extension: %s", f)
		}
	}

	Log().Infof("'%s' compiled in %v", p.Name, sw.Elapsed())

	return nil
}

func (p ProcessStruct) processScripts() error {
	sw := utils.StartTimer()

	files := p.Files()

	if p.DistFile != "" {
		imports := []string{}
		for _, f := range files {
			imports = append(imports, f.InFile)
		}

		if !utils.IsDir(p.Dist) {
			/* #nosec G301 */
			if err := os.MkdirAll(p.Dist, 0755); err != nil {
				return err
			}
		}

		out := path.Join(p.Dist, p.DistFile)

		options := api.BuildOptions{
			Stdin: &api.StdinOptions{
				Contents: "",
			},
			Inject:         imports,
			Outfile:        out,
			Write:          true,
			AllowOverwrite: true,
			Format:         api.FormatCommonJS,
			SourcesContent: api.SourcesContentExclude,
		}

		if p.JSBundle {
			options.Bundle = true
		}

		if Minify {
			options.MinifyWhitespace = true
			options.MinifyIdentifiers = true
			options.MinifySyntax = true
		} else {
			options.Sourcemap = api.SourceMapLinked
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

		Log().Debugf("compiled %d JS files to %s", len(files), out)
		Log().Infof("'%s' compiled in %v", p.Name, sw.Elapsed())
		return nil
	}

	for _, f := range files {
		filename := filepath.Base(f.InFile)
		d := path.Join(p.Dist, f.OutPath)
		if !utils.IsDir(d) {
			/* #nosec G301 */
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}
		out := path.Join(d, filename)

		options := api.BuildOptions{
			EntryPoints:    []string{f.InFile},
			Outfile:        out,
			Write:          true,
			AllowOverwrite: true,
			SourcesContent: api.SourcesContentExclude,
		}

		if p.JSBundle {
			options.Bundle = true
		}

		if Minify {
			options.MinifyWhitespace = true
			options.MinifyIdentifiers = true
			options.MinifySyntax = true
		} else {
			options.Sourcemap = api.SourceMapLinked
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

		Log().Debugf("compiled %s to %s", f.InFile, out)
	}

	Log().Infof("'%s' compiled in %v", p.Name, sw.Elapsed())

	return nil
}

func (p ProcessStruct) processCopy() error {
	sw := utils.StartTimer()

	files := p.Files()

	for _, f := range files {
		filename := filepath.Base(f.InFile)
		d := path.Join(p.Dist, f.OutPath)
		if !utils.IsDir(d) {
			/* #nosec G301 */
			if err := os.MkdirAll(d, 0755); err != nil {
				return err
			}
		}
		out := path.Join(d, filename)

		if err := utils.Copy(f.InFile, out); err != nil {
			return err
		}

		Log().Debugf("copied %s to %s", f.InFile, out)
	}

	Log().Infof("'%s' copied in %v", p.Name, sw.Elapsed())

	return nil
}
