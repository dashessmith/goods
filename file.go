package goods

import (
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
