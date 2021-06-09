package butler

import (
	"errors"
	"fmt"
	"github.com/BurntSushi/toml"
	"log"
	"os"
	"path"
	"strings"
)
type globalConfig struct {
	DotfilePath string
	DryRun bool
}
var GlobalConfig = globalConfig{}
type TomlConfig map[string]Application


type Application struct {
	Files map[string]string
	Variables map[string]string
	Before string
	After string
}

func (cfg *TomlConfig) New(filepath string){
	if _, err := toml.DecodeFile(filepath, &cfg); err != nil {
		fmt.Println(err)
		return
	}
}

func (cfg *TomlConfig) GetFileConfigForApp(app string) (map[string]string, error) {
	result := map[string]string{}
	deref := *cfg
		if deref[app].Files != nil {
			variables := deref[app].Variables
			for file := range deref[app].Files {
				dest := deref[app].Files[file]

				homedir := os.Getenv("HOME")
				absDest := strings.Replace(dest, "~", homedir,1)

				src := path.Join("./dotfiles",app, file)
				//absSrc, err := filepath.Abs(src)

				cachePath, err := Build(src, variables)
				if err != nil {
					log.Fatalln(err)
				}
				result[cachePath] = absDest
			}
		} else {
			return nil, errors.New("Failed to find config for " + app)
		}

	return result, nil
}
func (cfg *TomlConfig) GetAppNames() []string {
	apps := make([]string,0,len(*cfg))
	for app := range *cfg {
		log.Println(app)
		apps = append(apps, app)
	}
	return apps
}