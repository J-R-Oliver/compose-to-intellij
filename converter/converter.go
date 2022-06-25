package converter

import (
	"fmt"

	"github.com/compose-spec/compose-go/types"
)

// IntellijEnvVars converts environment variables configured in types.Services in types.Project to a map. The map keys
// are the service names and the values are environment variables in IntelliJ string format.
func IntellijEnvVars(p *types.Project) map[string]string {
	e := make(map[string]string)

	for _, s := range p.Services {
		g := ""

		for n, v := range s.Environment {
			g += fmt.Sprintf("%s=%s;", n, *v)
		}

		e[s.Name] = g
	}

	return e
}
