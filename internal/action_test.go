package internal

import (
	"io/ioutil"
	"os"
	"strings"
	"testing"
)

func TestAction(t *testing.T) {
	r, w, _ := os.Pipe()
	os.Stdout = w

	Action([]string{"application"}, map[string]string{"input": "./testdata/docker-compose.yml"})

	w.Close()

	out, _ := ioutil.ReadAll(r)
	s := string(out)

	e := "docker-compose to IntelliJ Environment variable convert\n\nFile input: ./testdata/docker-compose.yml\n\nService                 IntelliJ Env String\napplication          -> "
	if !strings.Contains(s, e) {
		t.Errorf("Expected:\n%s\nGot:\n%s", e, out)
	}

	e = "CONTEXT_PATH=/;"
	if !strings.Contains(s, e) {
		t.Errorf("Expected to contain: %s\nGot:\n%s", e, out)
	}

	e = "PORT=8080;"
	if !strings.Contains(s, e) {
		t.Errorf("Expected to contain: %s\nGot:\n%s", e, out)
	}
}

func Test_printOutput(t *testing.T) {
	type args struct {
		input     string
		arguments []string
		envVars   map[string]string
	}

	type serviceText struct {
		serviceName string
		envVars     string
	}

	type want struct {
		heading string
		sT      []serviceText
	}

	a1 := args{
		"./docker-compose.yml",
		[]string{},
		map[string]string{"application": "PORT=8080;LOG_LEVEL=debug;", "database": "VERBOSE=true;"},
	}
	w1 := want{
		heading: "docker-compose to IntelliJ Environment variable convert\n\nFile input: ./docker-compose.yml\n\nService                 IntelliJ Env String\n",
		sT: []serviceText{
			{"application", "PORT=8080;LOG_LEVEL=debug;"},
			{"database", "VERBOSE=true;"},
		},
	}

	a2 := args{
		"./docker-compose.yml",
		[]string{"application"},
		map[string]string{"application": "PORT=8080;LOG_LEVEL=debug;", "database": "VERBOSE=true;"},
	}
	w2 := want{
		heading: "docker-compose to IntelliJ Environment variable convert\n\nFile input: ./docker-compose.yml\n\nService                 IntelliJ Env String\n",
		sT: []serviceText{
			{"application", "PORT=8080;LOG_LEVEL=debug;"},
		},
	}

	tests := []struct {
		name string
		args args
		want want
	}{
		{"When passed arguments with length zero then prints all env vars", a1, w1},
		{"When passed arguments then prints specified env vars", a2, w2},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			r, w, _ := os.Pipe()
			os.Stdout = w

			printOutput(tt.args.input, tt.args.arguments, tt.args.envVars)

			w.Close()

			out, _ := ioutil.ReadAll(r)
			s := string(out)

			if !strings.Contains(s, tt.want.heading) {
				t.Errorf("Expected:\n%s\nGot:\n%s", tt.want.heading, out)
			}

			for _, w := range tt.want.sT {
				if !strings.Contains(s, w.serviceName) {
					t.Errorf("Expected to contain: %s\nGot:\n%s", w.serviceName, out)
				}

				if !strings.Contains(s, w.envVars) {
					t.Errorf("Expected to contain: %s\nGot:\n%s", w.envVars, out)
				}
			}
		})
	}
}
