package generator

import "testing"

func TestRandomPassword(t *testing.T) {
	passwordlength = 12
	got := RandomPassword()
	if len(got) != 12 {
		t.Error()
	}

}

// func TestrandomDate(t *testing.T) {
// 	date := randomDate()
// 	reqexformat // todo
// }
func TestGenerateUserName(t *testing.T) {
	username := generateUserName("test", "last", "efwefwe92")
	if username == "" {
		t.Error()
	}
}
