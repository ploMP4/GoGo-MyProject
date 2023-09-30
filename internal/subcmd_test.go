package internal

import "testing"

func TestSubcmd_runCommands(t *testing.T) {
	type subcmd struct {
		Name     string
		Commands Commands
		Override bool
		Parallel bool
		Exclude  bool
		Files    Files
		Help     string
	}

	tests := []struct {
		name    string
		app     App
		subcmd  subcmd
		wantErr bool
	}{
		{
			name: "Unable to execute command",
			app: App{
				spinner: loadSpinner(),
			},
			subcmd: subcmd{
				Commands: []string{"mkdir"},
			},
			wantErr: true,
		},
		{
			name: "Success",
			app: App{
				spinner: loadSpinner(),
			},
			subcmd: subcmd{
				Commands: []string{"echo test"},
			},
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SubCommand{
				Name:     tt.subcmd.Name,
				Commands: tt.subcmd.Commands,
				Override: tt.subcmd.Override,
				Parallel: tt.subcmd.Parallel,
				Exclude:  tt.subcmd.Exclude,
				Files:    tt.subcmd.Files,
				Help:     tt.subcmd.Help,
			}
			app = &tt.app

			err := s.runCommands()

			if (err != nil) != tt.wantErr {
				t.Errorf("SubCommand.runCommands() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestSubcmd_getTemplatePath(t *testing.T) {
	type subcmd struct {
		Name     string
		Commands Commands
		Override bool
		Parallel bool
		Exclude  bool
		Files    Files
		Help     string
	}
	type args struct {
		file File
		name string
	}

	tests := []struct {
		name    string
		app     App
		args    args
		subcmd  subcmd
		want    string
		wantErr bool
	}{
		{
			name: "Template not found",
			app: App{
				gadgetName: "cpp",
				parser: Parser{
					settings: Settings{
						GadgetPath:   "",
						TemplatePath: "./testdata/templates",
					},
				},
				spinner: loadSpinner(),
			},
			args: args{
				file: File{
					Filepath: "main.java",
					Template: true,
				},
				name: "basic",
			},
			subcmd:  subcmd{},
			want:    "",
			wantErr: true,
		},
		{
			name: "Success",
			app: App{
				gadgetName: "cpp",
				parser: Parser{
					settings: Settings{
						GadgetPath:   "",
						TemplatePath: "./testdata/templates",
					},
				},
				spinner: loadSpinner(),
			},
			args: args{
				file: File{
					Filepath: "Makefile",
					Template: true,
				},
				name: "basic",
			},
			subcmd:  subcmd{},
			want:    "./testdata/templates/cpp/basic/Makefile",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			s := &SubCommand{
				Name:     tt.subcmd.Name,
				Commands: tt.subcmd.Commands,
				Override: tt.subcmd.Override,
				Parallel: tt.subcmd.Parallel,
				Exclude:  tt.subcmd.Exclude,
				Files:    tt.subcmd.Files,
				Help:     tt.subcmd.Help,
			}
			app = &tt.app

			path, err := s.getTemplatePath(tt.args.file, tt.args.name)

			if path != tt.want {
				t.Errorf("Gadget.runCommands() got = %v, want = %v", path, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Gadget.runCommands() error = %v, wantErr = %v", err, tt.wantErr)
			}

		})
	}
}
