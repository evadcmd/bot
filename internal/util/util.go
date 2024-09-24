package util

import (
	"path"
	"runtime"
)

var RootPath string

func init() {
	_, f, _, _ := runtime.Caller(0)
	RootPath = path.Join(path.Dir(f), "../../")
}
