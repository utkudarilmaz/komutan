package main

import (
	"komutan/cmd"
	"komutan/core"
)

func init() {
	core.InitLogger()
}

func main() {
	cmd.Execute()
}
