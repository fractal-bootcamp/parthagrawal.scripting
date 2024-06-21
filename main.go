package main

import (
	"github.com/fractal-bootcamp/parthagrawal.scripting/cmd"
	"github.com/fractal-bootcamp/parthagrawal.scripting/data"
)

func main() {
	data.OpenDatabase()
	cmd.Execute()

}
