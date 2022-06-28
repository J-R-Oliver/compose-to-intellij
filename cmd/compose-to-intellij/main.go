package main

import (
	"github.com/J-R-Oliver/compose-to-intellij/internal"
	"github.com/J-R-Oliver/go-commando"
)

func main() {
	program := commando.NewProgram()

	program.
		Name("converter").
		Description("Command line application to convert Docker Compose environment variables to IntelliJ format. Optional arguments of compose service names can be passed to filter the output.").
		Version("1.0.0").
		Option("i", "input", "input", "filepath for docker-compose YAML file", "./docker-compose.yml").
		Action(internal.Action).
		Parse()
}
