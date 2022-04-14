package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"
	"strings"

	"github.com/axllent/golp/utils"
	"gopkg.in/yaml.v3"

	fg "github.com/goreleaser/fileglob"
)

var (
	// Minify determines whether to minify the styles and scripts
	Minify bool
)

// ParseConfig reads a yaml file and returns a Conf struct
func ParseConfig() error {
	var yml = yamlConf{}

	if !utils.IsFile(Conf.ConfigFile) {
		return fmt.Errorf("Config %s does not exist", Conf.ConfigFile)
	}

	buf, err := ioutil.ReadFile(Conf.ConfigFile)
	if err != nil {
		return err
	}

	err = yaml.Unmarshal(buf, &yml)
	if err != nil {
		return err
	}

	Conf.WorkingDir, err = filepath.Abs(filepath.Dir(Conf.ConfigFile))
	if err != nil {
		return err
	}

	for _, c := range yml.Clean {
		Conf.CleanDirs = append(Conf.CleanDirs, filepath.Join(Conf.WorkingDir, c))
	}

	for _, p := range yml.Styles {
		c := ProcessStruct{}
		c.Type = "styles"
		c.Name = p.Name
		if c.Name == "" {
			c.Name = c.Type
		}
		c.Src = p.Src
		if strings.HasSuffix(p.Dist, ".css") {
			c.DistFile = filepath.Base(p.Dist)
			p.Dist = filepath.Dir(p.Dist)
		}
		c.Dist = filepath.Join(Conf.WorkingDir, p.Dist)
		Conf.Process = append(Conf.Process, c)
	}

	for _, p := range yml.Scripts {
		c := ProcessStruct{}
		c.Type = "scripts"
		c.Name = p.Name
		if c.Name == "" {
			c.Name = c.Type
		}
		c.Src = p.Src
		if strings.HasSuffix(p.Dist, ".js") {
			c.DistFile = filepath.Base(p.Dist)
			p.Dist = filepath.Dir(p.Dist)
		}
		c.Dist = filepath.Join(Conf.WorkingDir, p.Dist)
		c.JSBundle = p.Bundle

		Conf.Process = append(Conf.Process, c)
	}

	for _, p := range yml.Copy {
		c := ProcessStruct{}
		c.Type = "copy"
		c.Name = p.Name
		if c.Name == "" {
			c.Name = c.Type
		}
		c.Src = p.Src
		c.Dist = filepath.Join(Conf.WorkingDir, p.Dist)
		Conf.Process = append(Conf.Process, c)
	}

	return nil
}

// Files returns all files matching the glob pattern
// should maybe use https://github.com/goreleaser/fileglob
func (p ProcessStruct) Files() []FileMap {

	fm := []FileMap{}
	exists := map[string]bool{}

	for _, pth := range p.Src {
		fullpth := filepath.Join(Conf.WorkingDir, pth)
		matches, err := fg.Glob(fullpth, fg.MaybeRootFS)
		if err == nil {
			subDir := ""
			subDirFrom := ""

			if strings.Contains(fullpth, "*") {
				parts := strings.Split(fullpth, "*")
				subDirFrom = parts[0]
			}

			for _, f := range matches {
				if utils.IsFile(f) {
					// only add each file once
					if _, ok := exists[f]; ok {
						continue
					}

					if subDirFrom != "" {
						if strings.HasPrefix(f, subDirFrom) {
							if len(filepath.Dir(f)) > len(subDirFrom) {
								subDir = filepath.Dir(f)[len(subDirFrom):]
							}
						}
					}

					fm = append(fm, FileMap{InFile: f, OutPath: subDir})

					exists[f] = true
				}
			}
		} else {
			Log().Error(err)
		}
	}

	return fm
}
