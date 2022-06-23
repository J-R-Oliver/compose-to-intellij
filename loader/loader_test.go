package loader

import (
	"reflect"
	"testing"

	"github.com/compose-spec/compose-go/types"
)

func TestCompose(t *testing.T) {
	s := new(string)
	*s = "8080"

	l := new(string)
	*l = "/"

	w := new(string)
	*w = "9090"

	x := new(string)
	*x = "verbose"

	a := map[string]*string{
		"PORT":         s,
		"CONTEXT_PATH": l,
	}

	d := map[string]*string{
		"PORT":   w,
		"OUTPUT": x,
	}

	tests := []struct {
		name     string
		filePath string
		aEnv     types.MappingWithEquals
		dEnv     types.MappingWithEquals
		wantErr  bool
	}{
		{"When passed a filepath to a valid compose file then returns a *types.Project", "./testdata/docker-compose.valid.yml", a, d, false},
		{"When passed a filepath to an invalid compose file then returns error", "./testdata/docker-compose.invalid.yml", nil, nil, true},
		{"When passed a filepath to a missing compose file then returns error", "./testdata/docker-compose.missing.yml", nil, nil, true},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := Compose(tt.filePath)
			if (err != nil) != tt.wantErr {
				t.Errorf("Compose() error = %v, wantErr %v", err, tt.wantErr)
				return
			}

			if err != nil {
				return
			}

			if len(got.Services) != 2 {
				t.Errorf("Compose() got = %v, want %v", len(got.Services), 2)
			}

			for _, s := range got.Services {
				if s.Name == "application" {
					if !reflect.DeepEqual(s.Environment, tt.aEnv) {
						t.Errorf("Compose() = %v, want %v", s.Environment, tt.aEnv)
					}
					break
				}

				if s.Name == "database" {
					if !reflect.DeepEqual(s.Environment, tt.dEnv) {
						t.Errorf("Compose() = %v, want %v", s.Environment, tt.dEnv)
					}
					break
				}

				t.Errorf("Compose() = %v, want service name 'application' or 'database'", s.Name)
			}
		})
	}
}
