package internal

import (
	"os"
	"reflect"
	"syscall"
	"testing"

	"github.com/briandowns/spinner"
)

func TestRoot_validateInput(t *testing.T) {
	tests := []struct {
		name    string
		args    []string
		want    string
		want1   []string
		wantErr bool
	}{
		{
			name:    "no args",
			args:    []string{},
			want:    "",
			want1:   nil,
			wantErr: true,
		},
		{
			name:    "command and app name",
			args:    []string{"", "cpp", "my_app"},
			want:    "cpp",
			want1:   []string{"my_app"},
			wantErr: false,
		},
		{
			name:    "command app name and arguments",
			args:    []string{"", "cpp", "my_app", "v"},
			want:    "cpp",
			want1:   []string{"my_app", "v"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(stdout *os.File) {
				os.Stdout = stdout
			}(os.Stdout)
			os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

			os.Args = tt.args

			got, got1, err := validateInput()
			if (err != nil) != tt.wantErr {
				t.Errorf("validateInput() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("validateInput() got = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateInput() got2 = %v, want %v", got1, tt.want1)
			}
		})
	}
}

func TestApp_runMainCommands(t *testing.T) {
	type app struct {
		filename string
		appName  string
		parser   Parser
		spinner  *spinner.Spinner
	}
	type args struct {
		mainCommands MainCommmands
	}
	a := app{
		filename: "cpp",
		appName:  "my_app",
		parser:   Parser{},
		spinner:  loadSpinner(),
	}
	tests := []struct {
		name    string
		app     app
		args    args
		want    string
		wantErr bool
	}{
		{
			name: "run a command",
			app:  a,
			args: args{
				mainCommands: []string{"echo " + a.appName},
			},
			want:    "",
			wantErr: false,
		},
		{
			name: "command doesn't exist",
			app:  a,
			args: args{
				mainCommands: []string{"some_command"},
			},
			want:    "Unable to execute command: some_command",
			wantErr: true,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(stdout *os.File) {
				os.Stdout = stdout
			}(os.Stdout)
			os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

			app := &App{
				filename: tt.app.filename,
				appName:  tt.app.appName,
				parser:   tt.app.parser,
				spinner:  tt.app.spinner,
			}

			got, err := app.runMainCommands(tt.args.mainCommands)
			if (err != nil) != tt.wantErr {
				t.Errorf("App.runMainCommands() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if got != tt.want {
				t.Errorf("App.runMainCommands() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestApp_executeSubCommand(t *testing.T) {
	type app struct {
		filename string
		appName  string
		parser   Parser
		spinner  *spinner.Spinner
	}
	type args struct {
		command SubCommand
	}
	a := app{
		filename: "cpp",
		appName:  "my_app",
		parser:   Parser{},
		spinner:  loadSpinner(),
	}
	tests := []struct {
		name    string
		fields  app
		args    args
		wantErr bool
	}{
		{
			name:   "invalid subcommand passed",
			fields: a,
			args: args{
				command: SubCommand{
					Name:    "test",
					Command: "some_command",
				},
			},
			wantErr: true,
		},
		{
			name:   "run a subcommand",
			fields: a,
			args: args{
				command: SubCommand{
					Name:    "echo",
					Command: "echo test",
				},
			},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			defer func(stdout *os.File) {
				os.Stdout = stdout
			}(os.Stdout)
			os.Stdout = os.NewFile(uintptr(syscall.Stdin), os.DevNull)

			app := &App{
				filename: tt.fields.filename,
				appName:  tt.fields.appName,
				parser:   tt.fields.parser,
				spinner:  tt.fields.spinner,
			}
			if err := app.executeSubCommand(tt.args.command); (err != nil) != tt.wantErr {
				t.Errorf("App.executeSubCommand() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
