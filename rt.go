package goods

import (
	"os"
	"path/filepath"
	"strings"
)

var (
	BinName    = ""
	BinExt     = ""
	BinNameExt = ""
	BinDir     = ""
	BinPath    = ""
)

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
