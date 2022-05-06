package generator

import "testing"

func Test_randomPassword(t *testing.T) {
	tests := []struct {
		name string
		want string
	}{
		{
			name: "Test 1: things",
			want: "dffd",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := randomPassword(); got != tt.want {
				t.Errorf("randomPassword() = %v, want %v", got, tt.want)
			}
		})
	}
}
