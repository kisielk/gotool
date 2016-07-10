// +build go1.4

package gotool

import (
	"go/build"
	"path/filepath"
	"runtime"
)

var gorootSrcPkg = filepath.Join(runtime.GOROOT(), "src")

func shouldIgnoreImport(p *build.Package) bool {
	return true
}
