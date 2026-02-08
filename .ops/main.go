package main

import (
	"os"

	"labs.lesiw.io/ops/golib"
	"lesiw.io/ops"
)

func main() {
	golib.Targets = []golib.Target{
		{Goos: "linux", Goarch: "386"},
		{Goos: "linux", Goarch: "amd64"},
		{Goos: "linux", Goarch: "arm"},
		{Goos: "linux", Goarch: "arm64"},
		{Goos: "darwin", Goarch: "amd64"},
		{Goos: "darwin", Goarch: "arm64"},
		{Goos: "windows", Goarch: "386"},
		{Goos: "windows", Goarch: "arm"},
		{Goos: "windows", Goarch: "amd64"},
	}
	if len(os.Args) < 2 {
		os.Args = append(os.Args, "check")
	}
	ops.Handle(golib.Ops{})
}
