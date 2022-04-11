package app

import (
	"fmt"
	"io/ioutil"
	"path/filepath"

	"github.com/axllent/golp/utils"
	"gopkg.in/yaml.v3"
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

	baseDir := filepath.Dir(Conf.ConfigFile)

	Conf.CleanDirs = yml.Clean

	for _, p := range yml.Styles {
		c := ProcessStruct{}
		c.Type = "styles"
		c.Src = p.Src
		c.Dist = filepath.Join(baseDir, p.Dist)
		Conf.Process = append(Conf.Process, c)
	}

	for _, p := range yml.Scripts {
		c := ProcessStruct{}
		c.Type = "scripts"
		c.Src = p.Src
		c.Dist = filepath.Join(baseDir, p.Dist)
		Conf.Process = append(Conf.Process, c)
	}

	for _, p := range yml.Copy {
		c := ProcessStruct{}
		c.Type = "copy"
		c.Src = p.Src
		c.Dist = filepath.Join(baseDir, p.Dist)
		Conf.Process = append(Conf.Process, c)
	}

	return nil
}
