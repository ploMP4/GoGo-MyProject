package internal

import "testing"

func TestGadget_runCommands(t *testing.T) {
	type gadget struct {
		Commands    Commands
		Chdir       bool
		Dirs        []string
		Files       Files
		SubCommands SubCommands
		Help        string
	}

	tests := []struct {
		name    string
		app     App
		gadget  gadget
		want    string
		wantErr bool
	}{
		{
			name: "No gadget commands",
			app: App{
				spinner: loadSpinner(),
			},
			gadget: gadget{
				Commands: nil,
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "Unable to execute command",
			app: App{
				spinner: loadSpinner(),
			},
			gadget: gadget{
				Commands: []string{"mkdir"},
			},
			want:    "Unable to execute command: mkdir",
			wantErr: true,
		},
		{
			name: "Success",
			app: App{
				spinner: loadSpinner(),
			},
			gadget: gadget{
				Commands: []string{"echo test"},
			},
			want:    "",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gadget{
				Commands:    tt.gadget.Commands,
				Chdir:       tt.gadget.Chdir,
				Dirs:        tt.gadget.Dirs,
				Files:       tt.gadget.Files,
				SubCommands: tt.gadget.SubCommands,
				Help:        tt.gadget.Help,
			}
			app = &tt.app

			msg, err := g.runCommands()

			if msg != tt.want {
				t.Errorf("Gadget.runCommands() got = %v, want = %v", msg, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Gadget.runCommands() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}

func TestGadget_getTemplatePath(t *testing.T) {
	type gadget struct {
		Commands    Commands
		Chdir       bool
		Dirs        []string
		Files       Files
		SubCommands SubCommands
		Help        string
	}
	type args struct {
		file File
	}

	tests := []struct {
		name    string
		app     App
		args    args
		gadget  gadget
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
			},
			gadget:  gadget{},
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
					Filepath: "main.cpp",
					Template: true,
				},
			},
			gadget:  gadget{},
			want:    "./testdata/templates/cpp/main.cpp",
			wantErr: false,
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			g := &Gadget{
				Commands:    tt.gadget.Commands,
				Chdir:       tt.gadget.Chdir,
				Dirs:        tt.gadget.Dirs,
				Files:       tt.gadget.Files,
				SubCommands: tt.gadget.SubCommands,
				Help:        tt.gadget.Help,
			}
			app = &tt.app

			path, err := g.getTemplatePath(tt.args.file)

			if path != tt.want {
				t.Errorf("Gadget.runCommands() got = %v, want = %v", path, tt.want)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("Gadget.runCommands() error = %v, wantErr = %v", err, tt.wantErr)
			}
		})
	}
}
