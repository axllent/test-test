package app

import (
	"os"
	"path/filepath"
	"strings"

	"github.com/bep/golibsass/libsass"
)

func compileStyles(content, outFile, inFile string) error {
	wi, err := os.Create(filepath.Clean(outFile))
	if err != nil {
		return err
	}
	/* #nosec G307 */
	defer wi.Close()

	options := libsass.Options{}
	options.IncludePaths = []string{"."}

	extension := strings.ToLower(filepath.Ext(inFile))

	if Minify {
		options.OutputStyle = libsass.CompressedStyle
	} else {
		options.OutputStyle = libsass.ExpandedStyle
		if extension != ".css" {
			options.SourceMapOptions = libsass.SourceMapOptions{
				EnableEmbedded: false,
				// InputPath:      inFile,
				OutputPath: outFile,
				Filename:   outFile + ".map",
				// Root: ".",
			}
		}
	}

	transpiler, _ := libsass.New(options)

	result, err := transpiler.Execute(content)
	if err != nil {
		return err
	}

	if _, err := wi.WriteString(result.CSS); err != nil {
		return err
	}

	if result.SourceMapFilename != "" && extension != ".css" {
		wi, err := os.Create(filepath.Clean(outFile + ".map"))
		if err != nil {
			return err
		}
		/* #nosec G307 */
		defer wi.Close()

		if _, err := wi.WriteString(result.SourceMapContent); err != nil {
			return err
		}
	}

	return nil
}
