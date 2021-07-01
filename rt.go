package util

import (
	"os"
	"path/filepath"
	"strings"
)

var BinName = ""
var BinExt = ""
var BinNameExt = ""
var BinDir = ""
var BinPath = ""

func init() {
	var err error
	BinPath, err = filepath.Abs(os.Args[0])
	if err != nil {
		panic(err)
	}
	BinPath = strings.Replace(BinPath, "\\", "/", -1)
	BinDir, BinNameExt = filepath.Split(BinPath)
	BinExt = filepath.Ext(BinNameExt)
	BinName = strings.TrimSuffix(BinNameExt, BinExt)
}
