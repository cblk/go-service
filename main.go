package main

import (
	"log"

	"go-service/cmds"
)

var (
	sha = ""
)

func main() {
	log.Println("Version:", sha)
	cmds.Execute()
}
