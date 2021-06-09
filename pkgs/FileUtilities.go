package butler

import (
	"errors"
	"io"
	"log"
	"os"
	"path"
	"path/filepath"
)

func Exists(path string) bool {
	_, err := os.Lstat(path)
	if err != nil {
		return false
	}
	return true
}

func Copy(src string, dest string) error {
	srcFile, err := os.Open(src)
	if err != nil {
		return err
	}
	defer srcFile.Close()

	dir := path.Dir(dest)
	if err := os.MkdirAll(dir, 0755); err != nil {
		return err
	}

	destFile, err := os.Create(dest)
	if err != nil  {
		return err
	}
	_, err = io.Copy(destFile, srcFile)
	if err != nil {
		return err
	}

	destFile.Sync()
	return nil
}

// symlinks source to destination with optional overwrite
// fails if file already exists and overwrite is false

func SymlinkFile(source string, destination string, overwrite bool) error {
	dryRun := GlobalConfig.DryRun
	exists := Exists(destination)

	absSrc, err := filepath.Abs(source)
	if err != nil {
		return err
	}
	if exists && overwrite {
		log.Println("Removing destination " + destination)
		if !dryRun {
			os.Remove(destination)
		}

	}
	if !Exists(destination) {
		if !dryRun {
			log.Println("Symlinking " + absSrc + " to " +destination)
			dir, _ := path.Split(destination)
			if err := os.MkdirAll(dir, 0755); err != nil {
				log.Println(err.Error())
				return err
			}

			if err := os.Symlink(absSrc, destination); err != nil {
				log.Println(err.Error())
				return err
			}
		}

		log.Println("Symlinked "+ destination)
		return nil
	}

	log.Println("File exists, not replacing " + destination)
	return errors.New("File exists, not replacing")
}

func UnlinkFile(filePath string) error {
	dryRun := GlobalConfig.DryRun

	fileInfo, err := os.Lstat(filePath)
	if err != nil {
		log.Println("Failed to stat file "+ filePath)
		return err
	}
	if fileInfo.Mode()&os.ModeSymlink == os.ModeSymlink {
		log.Println("Cleaning "+filePath)
		if !dryRun {
			os.Remove(filePath)
		}
	}
	return nil
}
