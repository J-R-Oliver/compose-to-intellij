package internal

import (
	"fmt"
	"log"

	"github.com/J-R-Oliver/compose-to-intellij/converter"
	"github.com/J-R-Oliver/compose-to-intellij/loader"
)

func Action(arguments []string, options map[string]string) {
	i, ok := options["input"]
	if !ok {
		log.Fatalln("input for 'docker-compose.yml' is null")
	}

	p, err := loader.Compose(i)
	if err != nil {
		log.Fatalf("error loading '%s': %s", i, err)
	}

	e := converter.IntellijEnvVars(p)
	printOutput(i, arguments, e)
}

func printOutput(input string, arguments []string, envVars map[string]string) {
	fmt.Println("docker-compose to IntelliJ Environment variable convert")
	fmt.Printf("\nFile input: %s\n\n", input)
	fmt.Printf("%-24s%s\n", "Service", "IntelliJ Env String")

	if len(arguments) > 0 {
		for _, a := range arguments {
			fmt.Printf("%-20s -> %s\n", a, envVars[a])
		}
	} else {
		for k, v := range envVars {
			fmt.Printf("%-20s -> %s\n", k, v)
		}
	}
}
