package main

import (
	"log"

	"github.com/cblk/go-service/cmds"
)

var (
	sha = ""
)

func main() {
	log.Println("Version:", sha)
	cmds.Execute()
}
