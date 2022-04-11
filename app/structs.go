package app

import (
	"path"
	"strings"

	"github.com/axllent/golp/utils"
	fg "github.com/goreleaser/fileglob"
)

// Conf struct
var Conf struct {
	ConfigFile          string   // build process is relative to this config
	CleanDirs           []string // is set, this directory will be deleted with clean
	AbortOnProcessError bool
	// Watch               []string
	Process []ProcessStruct
}

// ProcessStruct for config
type ProcessStruct struct {
	Type string
	Src  []string
	Dist string
}

// YamlConf is the yaml struct
type yamlConf struct {
	Clean  []string `yaml:"clean"`
	Styles []struct {
		Src  []string `yaml:"src"`
		Dist string   `yaml:"dist"`
	} `yaml:"styles"`
	Scripts []struct {
		Src  []string `yaml:"src"`
		Dist string   `yaml:"dist"`
	} `yaml:"scripts"`
	Copy []struct {
		Src  []string `yaml:"src"`
		Dist string   `yaml:"dist"`
	} `yaml:"copy"`
}

type watchMap struct {
	Path          string
	ProcessStruct ProcessStruct
}

// Files returns all files matching the glob pattern
// should maybe use https://github.com/goreleaser/fileglob
func (p ProcessStruct) Files() map[string]string {
	paths := map[string]string{}
	for _, pth := range p.Src {
		matches, err := fg.Glob(pth)
		// matches, err := filepath.Glob(pth)
		if err == nil {
			subDir := ""
			subDirFrom := ""

			if strings.Contains(pth, "*") {
				parts := strings.Split(pth, "*")
				subDirFrom = parts[0]
			}

			for _, f := range matches {
				if utils.IsFile(f) {
					if subDirFrom != "" {
						if strings.HasPrefix(f, subDirFrom) {
							if len(path.Dir(f)) > len(subDirFrom) {
								subDir = path.Dir(f)[len(subDirFrom):]
							}
						}
					}
					paths[f] = subDir
				}
			}
		}
	}

	return paths
}
