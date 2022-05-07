package config

import (
	"reflect"
	"testing"
)

func Test_readConfigFile(t *testing.T) {
	tests := []struct {
		name    string
		want    *Config
		wantErr bool
	}{
		{
			name:    "Test 1: things",
			want:    &Config{Email: "@change.me", DBFilePath: ".im.csv"},
			wantErr: false,
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			got, err := readConfigFile()
			if (err != nil) != tt.wantErr {
				t.Errorf("readConfigFile() error = %v, wantErr %v", err, tt.wantErr)
				return
			}
			if !reflect.DeepEqual(got, tt.want) {
				t.Errorf("readConfigFile() = %v, want %v", got, tt.want)
			}
		})
	}
}
