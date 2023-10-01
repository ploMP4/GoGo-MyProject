package internal

import (
	"os"
	"reflect"
	"testing"
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
			args:    []string{"", "cpp", "my_app", "basic"},
			want:    "cpp",
			want1:   []string{"my_app", "basic"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			os.Args = tt.args

			got, got1, err := validateInput()

			if got != tt.want {
				t.Errorf("validateInput() got = %v, want = %v", got, tt.want)
			}

			if !reflect.DeepEqual(got1, tt.want1) {
				t.Errorf("validateInput() got2 = %v, want = %v", got1, tt.want1)
			}

			if (err != nil) != tt.wantErr {
				t.Errorf("validateInput() error = %v, wantErr %v", err, tt.wantErr)
			}
		})
	}
}
