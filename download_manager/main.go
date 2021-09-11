package main

import (
	"dlmm/command"
	"dlmm/manager"
)

func main() {
	c := &command.Command{}

	c.Init()

	manager := &manager.Manager{}

	manager.Init(c)

	manager.Download()
}
