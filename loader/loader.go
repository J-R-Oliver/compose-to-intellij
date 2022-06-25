package loader

import (
	"os"

	"github.com/compose-spec/compose-go/loader"
	"github.com/compose-spec/compose-go/types"
)

// Compose parses a valid docker-compose file.
func Compose(filePath string) (*types.Project, error) {
	f, err := os.ReadFile(filePath)
	if err != nil {
		return nil, err
	}

	c := types.ConfigDetails{
		Version:     "",
		WorkingDir:  "",
		ConfigFiles: []types.ConfigFile{{Content: f}},
		Environment: nil,
	}

	return loader.Load(c)
}
