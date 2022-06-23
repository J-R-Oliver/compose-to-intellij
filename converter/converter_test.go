package converter

import (
	"fmt"
	"reflect"
	"testing"

	"github.com/compose-spec/compose-go/types"
)

func ExampleIntellijEnvVars() {
	s := new(string)
	*s = "8080"

	l := new(string)
	*l = "debug"

	p := &types.Project{
		Services: []types.ServiceConfig{
			{
				Name: "application",
				Environment: map[string]*string{
					"PORT":      s,
					"LOG_LEVEL": l,
				},
			},
		},
	}

	e := IntellijEnvVars(p)
	fmt.Println(e)
	// Output: map[application:PORT=8080;LOG_LEVEL=debug;]
}

func TestIntellijEnvVars(t *testing.T) {
	s := new(string)
	*s = "8080"

	l := new(string)
	*l = "debug"

	v := new(string)
	*v = "true"

	p1 := &types.Project{
		Services: []types.ServiceConfig{
			{
				Name: "application",
				Environment: map[string]*string{
					"PORT": s,
				},
			},
		},
	}

	p2 := &types.Project{
		Services: []types.ServiceConfig{
			{
				Name: "application",
				Environment: map[string]*string{
					"PORT":      s,
					"LOG_LEVEL": l,
				},
			},
		},
	}

	p3 := &types.Project{
		Services: []types.ServiceConfig{
			{
				Name: "application",
				Environment: map[string]*string{
					"PORT":      s,
					"LOG_LEVEL": l,
				},
			},
			{
				Name: "database",
				Environment: map[string]*string{
					"VERBOSE": v,
				},
			},
		},
	}

	tests := []struct {
		name string
		p    *types.Project
		want map[string]string
	}{
		{"When passed one service with one env var then returns correct map", p1, map[string]string{"application": "PORT=8080;"}},
		{"When passed one service with multiple env var then returns correct map", p2, map[string]string{"application": "PORT=8080;LOG_LEVEL=debug;"}},
		{"When passed multiple services with multiple env vars then returns correct map", p3, map[string]string{"application": "PORT=8080;LOG_LEVEL=debug;", "database": "VERBOSE=true;"}},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := IntellijEnvVars(tt.p); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("IntellijEnvVars() = %v, want %v", got, tt.want)
			}
		})
	}
}
