package butler

import (
	"fmt"
	"github.com/aymerick/raymond"
	"os"
	"path"
)

const CacheDir = "./.cache/"
func Build(src string, variables map[string]string) (string, error) {
	outfile := path.Join(CacheDir, src)
	dir := path.Dir(outfile)

	os.MkdirAll(dir, 0755)
	fmt.Println("cache file "+outfile)

	if len(variables) == 0 {
		err := Copy(src, outfile)
		if err != nil {
			return "", nil
		}
		return outfile, nil
	}

	tpl, err := raymond.ParseFile(src)
	if err != nil {
		return "", err
	}
	result, err := tpl.Exec(variables)
	if err != nil {
		return "", err
	}

	err = os.WriteFile(outfile,[]byte(result), 0755)
	if err != nil {
		return "", err
	}

	return outfile, nil
}