package util

import "os"

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
