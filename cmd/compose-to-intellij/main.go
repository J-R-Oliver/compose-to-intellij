package main

import (
	"github.com/J-R-Oliver/compose-to-intellij/internal"
	"github.com/J-R-Oliver/go-commando"
)

func main() {
	program := commando.NewProgram()

	program.
		Name("converter").
		Description("Command line application to convert Docker Compose environment variables to IntelliJ format.").
		Version("1.0.0").
		Option("i", "input", "input", "", "./docker-compose.yml").
		Action(internal.Action).
		Parse()
}
