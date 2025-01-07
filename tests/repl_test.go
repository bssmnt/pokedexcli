package tests

import (
	"pokedexcli/internal/repl"
	"reflect"
	"testing"
)

func TestCleanInput(t *testing.T) {
	cases := []struct {
		input    string
		expected []string
	}{
		{
			input:    "  hello  world  ",
			expected: []string{"hello", "world"},
		},
	}
	for _, c := range cases {
		actual := repl.CleanInput(c.input)
		if len(actual) != len(c.expected) {
			t.Errorf("actual %v != expected %v", actual, c.expected)
		}
		for i := range actual {
			word := actual[i]
			expectedWord := c.expected[i]
			if word != expectedWord {
				t.Errorf("word: %s, expected: %s", word, expectedWord)
			}
		}
	}
}

func TestREPL(t *testing.T) {
	var tests []struct {
		name string
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			repl.REPL()
		})
	}
}

func Test_callbackExplore(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CallbackExplore(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("callbackExplore() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandCatch(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CommandCatch(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("commandCatch() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandExit(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CommandExit(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("commandExit() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandHelp(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CommandHelp(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("commandHelp() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandMap(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CommandMap(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("commandMap() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandMapB(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CommandMapB(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("commandMapB() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_commandPokedex(t *testing.T) {
	type args struct {
		cfg  *repl.Config
		args []string
	}
	var tests []struct {
		name    string
		args    args
		wantErr bool
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if err := repl.CommandPokedex(tt.args.cfg, tt.args.args...); (err != nil) != tt.wantErr {
				t.Errorf("commandPokedex() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func Test_getCommands(t *testing.T) {
	var tests []struct {
		name string
		want map[string]repl.CliCommand
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := repl.GetCommands(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("getCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}
