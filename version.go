package up

import (
	"fmt"
	"runtime"
)

var (
	CurrentCommit  = ""
	CurrentVersion = "1.0.0"
	BuildDate      = ""
	GoVersion      = runtime.Version()
	Platform       = fmt.Sprintf("%s/%s", runtime.GOOS, runtime.GOARCH)
)
