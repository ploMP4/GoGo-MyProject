package pkg

import (
	"reflect"
	"testing"
)

func TestParser_parseSettings(t *testing.T) {
	type parser struct {
		settings Settings
		gadget   Gadget
		args     []string
	}
	p := parser{args: []string{"v"}}
	tests := []struct {
		name    string
		parser  parser
		wantErr bool
	}{
		{
			name:    "non existant settings.yaml file",
			parser:  p,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				args: tt.parser.args,
			}
			if err := p.parseSettings(); (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseSettings() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_parseGadget(t *testing.T) {
	type parser struct {
		settings Settings
		gadget   Gadget
		args     []string
	}
	type args struct {
		filename string
	}
	p := parser{
		settings: Settings{
			GadgetPath: "./testdata/gadgets/",
		},
		args: []string{"v"},
	}
	tests := []struct {
		name    string
		parser  parser
		args    args
		wantErr bool
	}{
		{
			name:   "",
			parser: p,
			args: args{
				filename: "cpp",
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
			if err := p.parseGadget(tt.args.filename); (err != nil) != tt.wantErr {
				t.Errorf("Parser.parseGadget() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}

func TestParser_getHelp(t *testing.T) {
	type parser struct {
		settings Settings
		gadget   Gadget
		args     []string
	}
	tests := []struct {
		name   string
		parser parser
		want   []string
	}{
		// {
		// 	name: "",
		// 	parser: parser{
		// 		settings: Settings{
		// 			GadgetPath: "./testdata/gadgets/",
		// 		},
		// 		args: []string{"v"},
		// 	},
		// 	want: []string{},
		// },
		{
			name: "gadget path not set",
			parser: parser{
				args: []string{"v"},
			},
			want: nil,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				settings: tt.parser.settings,
				gadget:   tt.parser.gadget,
				args:     tt.parser.args,
			}
			if got := p.getHelp(); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_getSubHelp(t *testing.T) {
	type parser struct {
		settings Settings
		gadget   Gadget
		args     []string
	}
	type args struct {
		filename string
	}
	p := parser{
		settings: Settings{
			GadgetPath: "./testdata/gadgets/",
		},
		args: []string{"v"},
	}
	tests := []struct {
		name    string
		parser  parser
		args    args
		want    []string
		wantErr bool
	}{
		// {
		// 	name:   "",
		// 	parser: p,
		// 	args: args{
		// 		filename: "cpp",
		// 	},
		// 	want:    []string{},
		// 	wantErr: false,
		// },
		{
			name:   "non existant gadget",
			parser: p,
			args: args{
				filename: "some_name",
			},
			want:    nil,
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := Parser{
				settings: tt.parser.settings,
				gadget:   tt.parser.gadget,
				args:     tt.parser.args,
			}
			got, err := p.getSubHelp(tt.args.filename)
			if (err != nil) != tt.wantErr {
				t.Errorf("Parser.getSubHelp() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.getSubHelp() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestParser_parseArgs(t *testing.T) {
	type parser struct {
		settings Settings
		gadget   Gadget
		args     []string
	}
	p := parser{
		settings: Settings{GadgetPath: "./testdata/gadgets/"},
		args:     []string{"v"},
	}
	tests := []struct {
		name   string
		parser parser
		want   MainCommmands
		want1  []SubCommand
		want2  []string
	}{
		{
			name:   "pass",
			parser: p,
			want:   [][]string{{"mkdir"}},
			want1: []SubCommand{
				{
					Name:     "vanilla",
					Command:  []string{},
					Override: false,
					Parallel: true,
					Exclude:  false,
					Files: map[string]File{
						"makefile": {
							Filepath: "Makefile",
							Template: true,
							Change:   FileChange{},
						},
						"src": {
							Filepath: "src/main.cpp",
							Template: true,
							Change:   FileChange{},
						},
					},
					Help: "",
				},
			},
			want2: []string{"bin", "include", "src"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			p := &Parser{
				settings: tt.parser.settings,
				args:     tt.parser.args,
			}
			p.parseGadget("cpp")

			got, _, got2, _ := p.parseArgs()
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Parser.parseArgs() got = %v, want %v", got, tt.want)
			}
			// if !reflect.DeepEqual(got1, tt.want1) {
			// 	t.Errorf("Parser.parseArgs() got1 = %v, want %v", got1, tt.want1)
			// }
			if !reflect.DeepEqual(got2, tt.want2) {
				t.Errorf("Parser.parseArgs() got2 = %v, want %v", got2, tt.want2)
			}
		})
	}
}
