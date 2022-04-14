package app

// Conf struct
var Conf struct {
	ConfigFile string   // build process is relative to this config
	WorkingDir string   // working directory is the base directory of the config file
	CleanDirs  []string // is set, this directory will be deleted with clean
	Process    []ProcessStruct
}

// ProcessStruct for config
type ProcessStruct struct {
	Type string
	Name string
	// Src files
	Src []string
	// Dist directory
	Dist string
	// DistFile is the combined filename for all matching files if specified (JS/CSS only)
	DistFile string
	// JSBundle determines whether the dist JS should be bundled
	JSBundle bool
}

// YamlConf is the yaml struct
type yamlConf struct {
	Clean  []string `yaml:"clean"`
	Styles []struct {
		Name string   `yaml:"name"`
		Src  []string `yaml:"src"`
		Dist string   `yaml:"dist"`
	} `yaml:"styles"`
	Scripts []struct {
		Name   string   `yaml:"name"`
		Src    []string `yaml:"src"`
		Dist   string   `yaml:"dist"`
		Bundle bool     `yaml:"bundle"`
	} `yaml:"scripts"`
	Copy []struct {
		Name string   `yaml:"name"`
		Src  []string `yaml:"src"`
		Dist string   `yaml:"dist"`
	} `yaml:"copy"`
}

type watchMap struct {
	Path          string
	ProcessStruct ProcessStruct
}

// FileMap struct maps the file to the respective dist directory
type FileMap struct {
	InFile  string
	OutPath string
}
