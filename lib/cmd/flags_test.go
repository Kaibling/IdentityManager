package cmd

import (
	"reflect"
	"testing"
)

func TestParseFlags(t *testing.T) {
	Flags = map[string]bool{"a": false, "b": false}
	type args struct {
		args []string
	}
	tests := []struct {
		name  string
		args  args
		want  []string
		want1 map[string]bool
	}{
		{
			name:  "Test 1: things",
			args:  args{args: []string{"one", "zwo", "three", "-a"}},
			want:  []string{"one", "zwo", "three"},
			want1: map[string]bool{"a": true, "b": false},
		},
		{
			name:  "Test 2: more things",
			args:  args{args: []string{"-a", "one", "zwo", "-b", "three"}},
			want:  []string{"one", "zwo", "three"},
			want1: map[string]bool{"a": true, "b": true},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := ParseFlags(tt.args.args); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("ParseFlags() = %v, want %v", got, tt.want)
			}
			if !reflect.DeepEqual(Flags, tt.want1) {
				t.Errorf("ParseFlags() = %v, want %v", Flags, tt.want1)
			}
		})
	}
}
