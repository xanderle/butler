package main

import (
	"flag"
	butler "github.com/xanderle/butler/pkgs"
	"log"
	"os"
)




func main() {
	cleanPtr := flag.Bool("clean", false, "remove all symlinks")
	configPtr := flag.String("config", "host.toml", "toml config to use")
    dotfilePtr := flag.String("dots", "dotfiles", "path to dotfile directory")
	overwritePtr := flag.Bool("overwrite", false, "force overwrite destination dots")

	dryRunPtr := flag.Bool("dry-run", false, "won't symlink files out")
	flag.Parse()
	butler.GlobalConfig.DotfilePath = *dotfilePtr
	butler.GlobalConfig.DryRun = *dryRunPtr


	//os.Setenv("DRYRUN", strconv.FormatBool(*dryRunPtr))

	cfg := butler.TomlConfig{}
	cfg.New(*configPtr)

	apps := []string{}
	log.Println(os.Args)
	if len(flag.Args()) > 0{
		apps = flag.Args()
		// apply action for each param
	 } else {
		apps = cfg.GetAppNames()
	}

	for _, program := range apps {
		filesConfig, err := cfg.GetFileConfigForApp(program)
		if err != nil {
			log.Fatal(err.Error())
		}
		if *cleanPtr {
			for _, v := range filesConfig {
				butler.UnlinkFile(v)
			}
		} else {
			for k, v := range filesConfig {
				butler.SymlinkFile(k, v, *overwritePtr)
			}
		}
	}
}
