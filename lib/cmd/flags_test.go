package cmd

import "testing"

func TestParseFlags(t *testing.T) {
	type args struct {
		args []string
	}
	tests := []struct {
		name string
		args args
	}{
		{
			name: "Test 1: things",
			args: args{args: []string{"one", "zwo", "three", "-a"}},
		},
		{
			name: "Test 2: more things",
			args: args{args: []string{"-a", "one", "zwo", "-b", "three"}},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			ParseFlags(tt.args.args)
		})
	}
	t.Error()
}
