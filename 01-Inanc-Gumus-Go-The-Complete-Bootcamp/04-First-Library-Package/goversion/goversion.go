package goversion

import "runtime"

func GetVersion() string {
	return runtime.Version()
}
