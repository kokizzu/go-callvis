package main

import (
	"fmt"
)

var (
	version = "v0.7.1"
	commit  = "(unknown)"
)

func Version() string {
	return fmt.Sprintf("%s built from git %s", version, commit)
}
