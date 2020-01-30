package utils

import (
	"bytes"
	"reflect"
	"testing"
)

func TestAsk4confirm(t *testing.T) {
	type args struct {
		confirmMessage string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		// TODO: Add test cases.
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Ask4confirm(tt.args.confirmMessage); got != tt.want {
				t.Errorf("Ask4confirm() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestExists(t *testing.T) {
	type args struct {
		name string
	}
	tests := []struct {
		name string
		args args
		want bool
	}{
		{
			name: "check if the directory/file exists",
			args: args{
				name: "main.go",
			},
			want: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Exists(tt.args.name); got != tt.want {
				t.Errorf("Exists() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestRemoveDuplicatesString(t *testing.T) {
	type args struct {
		elements []string
	}
	tests := []struct {
		name string
		args args
		want []string
	}{
		{
			name: "delete duplicate string slice",
			args: args{
				elements: []string{"test", "test"},
			},
			want: []string{"test"},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := RemoveDuplicatesString(tt.args.elements); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("RemoveDuplicatesString() = %v, want %v", got, tt.want)
			}
		})
	}
}

func TestPrintCmd(t *testing.T) {
	type args struct {
		cmd     string
		verbose bool
	}
	tests := []struct {
		name  string
		args  args
		wantW string
	}{
		{
			name: "cmd output verbose",
			args: args{
				cmd:     "test output",
				verbose: true,
			},
			wantW: "test output\n",
		},
		{
			name: "cmd output no verbose",
			args: args{
				cmd:     "test output > /dev/null",
				verbose: true,
			},
			wantW: "test output > /dev/null\n",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			PrintCmd(w, tt.args.cmd, tt.args.verbose)
			if gotW := w.String(); gotW != tt.wantW {
				t.Errorf("PrintCmd() = %v, want %v", gotW, tt.wantW)
			}
		})
	}
}
