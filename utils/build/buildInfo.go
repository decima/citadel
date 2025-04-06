package build

import (
	"runtime/debug"
	"sync"
)

var version string = "dev"
var goVersion string = "dev"
var once sync.Once

func Initialize(v ...string) {
	once.Do(func() {
		if len(v) != 1 {
			panic("Initialize function requires exactly one argument")
		}
		version = v[0]

		buildInfo, ok := debug.ReadBuildInfo()
		if ok {
			goVersion = buildInfo.GoVersion
		}

	})
}

func Version() string {
	return version
}

func GoVersion() string {
	return goVersion
}
