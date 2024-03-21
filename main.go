/*
Copyright © 2024 tomo0611 <tomo0611@hotmail.com>
*/
package main

import "github.com/tomo0611/words-play/cmd"

var (
	version  = "v0.0.1"
	revision = "000000"
)

func main() {
	cmd.Version = version
	cmd.Revision = revision

	cmd.Execute()
}
