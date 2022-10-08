package goods

import (
	"fmt"
	"io"
	"os"
	"path"
)

/* check file exists */
func FileExists(filename string) (yes bool, err error) {
	fi, err := os.Stat(filename)
	if err != nil {
		if os.IsNotExist(err) {
			yes = false
			err = nil
			return
		}
		return
	}
	yes = fi.Mode().IsRegular()
	return
}

func EnsureDir(path string) (err error) {
	return os.MkdirAll(path, os.ModePerm)
}

func Rename(frompath, topath string) (err error) {
	distdir := path.Dir(topath)
	err = EnsureDir(distdir)
	if err != nil {
		return
	}
	return os.Rename(frompath, topath)
}

func Copy(frompath, topath string) (err error) {
	distdir := path.Dir(topath)
	err = EnsureDir(distdir)
	if err != nil {
		return
	}
	sourceFileStat, err := os.Stat(frompath)
	if err != nil {
		return
	}

	if !sourceFileStat.Mode().IsRegular() {
		return fmt.Errorf("%s is not a regular file", frompath)
	}

	source, err := os.Open(frompath)
	if err != nil {
		return err
	}
	defer source.Close()

	destination, err := os.Create(topath)
	if err != nil {
		return err
	}
	defer destination.Close()
	_, err = io.Copy(destination, source)
	return
}

func Symlink(target string, symlink string) (err error) {
	distdir := path.Dir(symlink)
	err = EnsureDir(distdir)
	if err != nil {
		return
	}
	return os.Symlink(distdir, symlink)
}
