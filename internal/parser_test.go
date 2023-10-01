package internal

import (
	"fmt"
	"reflect"
	"slices"
	"testing"
)

func TestParser_parseGadget(t *testing.T) {
	type parser struct {
		settings Settings
		args     []string
	}
	type args struct {
		filename string
	}

	tests := []struct {
		name    string
		parser  parser
		args    args
		want    Gadget
		wantErr bool
	}{
		{
			name: "Gadget not found",
			parser: parser{
				settings: Settings{
					GadgetPath: "./testdata/gadgets/",
				},
			},
			args: args{
				filename: "java",
			},
			want:    Gadget{},
			wantErr: true,
		},
		{
			name: "Success",
			parser: parser{
				settings: Settings{
					GadgetPath: "./testdata/gadgets/",
				},
			},
			args: args{
				filename: "cpp",
			},
			want: Gadget{
				Commands: Commands{"mkdir _APPNAME"},
				Dirs:     []string{"bin", "include", "src"},
				Files: Files{
					"src": {
						Filepath: "src/main.cpp",
						Template: true,
					},
				},
				SubCommands: SubCommands{
					"basic": {
						Name:     "basic",
						Override: false,
						Parallel: true,
						Files: Files{
							"makefile": {
								Filepath: "Makefile",
								Template: true,
							},
						},
						Help: "This is a basic setup",
					},
				},
				Help: "Create C++ app",
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				settings: tt.parser.settings,
				args:     tt.parser.args,
			}

			gadget, err := p.parseGadget(tt.args.filename)

			if !reflect.DeepEqual(gadget, tt.want) {
				t.Errorf("Parser.parseGadget() gadget = %v, want = %v", gadget, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseGadget() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_getHelp(t *testing.T) {
	type parser struct {
		settings Settings
		args     []string
	}
	tests := []struct {
		name   string
		parser parser
		want   []string
	}{
		{
			name:   "No gadgets found",
			parser: parser{},
			want:   nil,
		},
		{
			name: "Gadget help menu",
			parser: parser{
				settings: Settings{
					GadgetPath: "./testdata/gadgets/",
				},
			},
			want: []string{fmt.Sprintf("\n%30s   - %s", "cpp", "Create C++ app")},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				settings: tt.parser.settings,
				args:     tt.parser.args,
			}

			if got := p.getHelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getHelp() = %v, want = %v", got, tt.want)
			}
		})
	}
}

func TestParser_getSubHelp(t *testing.T) {
	type parser struct {
		settings Settings
		args     []string
	}
	type args struct {
		filename string
	}

	tests := []struct {
		name    string
		parser  parser
		args    args
		want    []string
		wantErr bool
	}{
		{
			name: "non existant gadget",
			parser: parser{
				settings: Settings{
					GadgetPath: "./testdata/gadgets/",
				},
			},
			args: args{
				filename: "some_name",
			},
			want:    nil,
			wantErr: true,
		},
		{
			name: "Success",
			parser: parser{
				settings: Settings{
					GadgetPath: "./testdata/gadgets/",
				},
			},
			args: args{
				filename: "cpp",
			},
			want:    []string{fmt.Sprintf("\n%31s   - %s", "basic", "This is a basic setup")},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				settings: tt.parser.settings,
				args:     tt.parser.args,
			}

			got, err := p.getSubHelp(tt.args.filename)

			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getSubHelp() = %v, want = %v", got, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.getSubHelp() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_parseFlagsAndPlaceholders(t *testing.T) {
	type parser struct {
		settings Settings
		args     []string
	}
	type args struct {
		gadget Gadget
	}

	tests := []struct {
		name   string
		parser parser
		args   args
		want1  bool
		want2  bool
		want3  bool
		want4  string
	}{
		{
			name:   "No flags",
			parser: parser{},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Exclude: false,
						},
					},
				},
			},
			want1: false,
			want2: false,
			want3: false,
			want4: "",
		},
		{
			name: "all and verbose flags",
			parser: parser{
				args: []string{"all", "vv"},
			},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Exclude: false,
						},
					},
				},
			},
			want1: true,
			want2: false,
			want3: true,
			want4: "",
		},
		{
			name: "appname",
			parser: parser{
				args: []string{"_APPNAME", "myapp"},
			},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Exclude: false,
						},
					},
				},
			},
			want1: false,
			want2: false,
			want3: false,
			want4: "myapp",
		},
		{
			name: "exclude test subcommand",
			parser: parser{
				args: []string{"e", "test"},
			},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Exclude: false,
						},
					},
				},
			},
			want1: false,
			want2: true,
			want3: false,
			want4: "",
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				settings: tt.parser.settings,
				args:     tt.parser.args,
			}

			all, verbose, appname := p.parseFlagsAndPlaceholders(&tt.args.gadget)

			if all != tt.want1 {
				t.Errorf("Parser.parseFlagsAndPlaceholders() all = %v, want = %v", all, tt.want1)
			}

			if tt.args.gadget.SubCommands["test"].Exclude != tt.want2 {
				t.Errorf(
					"Parser.parseFlagsAndPlaceholders() all = %v, want = %v",
					tt.args.gadget.SubCommands["test"].Exclude,
					tt.want2,
				)

			}

			if verbose != tt.want3 {
				t.Errorf(
					"Parser.parseFlagsAndPlaceholders() verbose = %v, want = %v",
					verbose,
					tt.want3,
				)
			}

			if appname != tt.want4 {
				t.Errorf(
					"Parser.parseFlagsAndPlaceholders() appname = %v, want = %v",
					appname,
					tt.want4,
				)
			}

		})
	}
}

func TestParser_parseAll(t *testing.T) {
	type args struct {
		gadget      Gadget
		commands    Commands
		subCommands SubCommands
	}

	tests := []struct {
		name  string
		args  args
		want1 Commands
		want2 SubCommands
	}{
		{
			name: "Override gadget commands",
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Commands: Commands{"echo test"},
							Override: true,
							Exclude:  false,
						},
					},
				},
				commands:    Commands{"echo"},
				subCommands: SubCommands{},
			},
			want1: Commands{"echo test"},
			want2: SubCommands{},
		},
		{
			name: "Exclude subcommand",
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Commands: Commands{"echo test"},
							Exclude:  true,
						},
					},
				},
				commands:    Commands{"echo"},
				subCommands: SubCommands{},
			},
			want1: Commands{"echo"},
			want2: SubCommands{},
		},
		{
			name: "Parse subcommands",
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Commands: Commands{"echo test"},
							Override: false,
							Exclude:  false,
						},
					},
				},
				commands:    Commands{"echo"},
				subCommands: SubCommands{},
			},
			want1: Commands{"echo"},
			want2: SubCommands{
				"test": {
					Commands: Commands{"echo test"},
					Override: false,
					Exclude:  false,
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{}

			p.parseAll(tt.args.gadget, &tt.args.commands, tt.args.subCommands)

			if !slices.Equal(tt.args.commands, tt.want1) {
				t.Errorf("Parser.parseAll() commands = %v, want = %v", tt.args.commands, tt.want1)
			}

			if !reflect.DeepEqual(tt.args.subCommands, tt.want2) {
				t.Errorf(
					"Parser.parseAll() subcommands = %v, want = %v",
					tt.args.subCommands,
					tt.want2,
				)
			}
		})
	}
}

func TestParser_parseCmd(t *testing.T) {
	type parser struct {
		settings Settings
		args     []string
	}
	type args struct {
		gadget      Gadget
		commands    Commands
		subCommands SubCommands
	}

	tests := []struct {
		name   string
		parser parser
		args   args
		want1  Commands
		want2  SubCommands
	}{
		{
			name: "No matching subcommands",
			parser: parser{
				args: []string{"non_valid"},
			},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Commands: Commands{"echo test"},
						},
					},
				},
				commands:    Commands{"echo"},
				subCommands: SubCommands{},
			},
			want1: Commands{"echo"},
			want2: SubCommands{},
		},
		{
			name: "Override gadget commands",
			parser: parser{
				args: []string{"test"},
			},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Commands: Commands{"echo test"},
							Override: true,
						},
					},
				},
				commands:    Commands{"echo"},
				subCommands: SubCommands{},
			},
			want1: Commands{"echo test"},
			want2: SubCommands{},
		},
		{
			name: "Parse subcommands",
			parser: parser{
				args: []string{"test"},
			},
			args: args{
				gadget: Gadget{
					SubCommands: SubCommands{
						"test": {
							Commands: Commands{"echo test"},
						},
					},
				},
				commands:    Commands{"echo"},
				subCommands: SubCommands{},
			},
			want1: Commands{"echo"},
			want2: SubCommands{
				"test": {
					Commands: Commands{"echo test"},
				},
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				settings: tt.parser.settings,
				args:     tt.parser.args,
			}

			p.parseCmd(tt.args.gadget, &tt.args.commands, tt.args.subCommands)

			if !slices.Equal(tt.args.commands, tt.want1) {
				t.Errorf("Parser.parseCmd() commands = %v, want = %v", tt.args.commands, tt.want1)
			}

			if !reflect.DeepEqual(tt.args.subCommands, tt.want2) {
				t.Errorf(
					"Parser.parseCmd() subcommands = %v, want = %v",
					tt.args.subCommands,
					tt.want2,
				)
			}
		})
	}
}
