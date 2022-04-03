/*
Copyright © 2022 Matt Carpenter <mattcarp88@gmail.com>

*/
package main

import (
	"log"

	"github.com/mattcarp12/go-gitm/cmd"
)

func main() {
	log.SetPrefix("[gitm]: ")
	log.SetFlags(0)
	cmd.Execute()
}
