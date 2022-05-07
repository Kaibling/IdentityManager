package encryption

import (
	"reflect"
	"testing"
)

// b := encryption.Encrypt([]byte("key"), []byte("super gehein"))
// s := encryption.EncodeB64(b)
// fmt.Println(s)
// b, err := encryption.DecodeB64(s)
// if err != nil {
// 	fmt.Println(err.Error())
// }
// p := encryption.Decrypt([]byte("key"), b)
// fmt.Println(string(p))
// func TestEncrypt(t *testing.T) {
// 	generateSalt = func() []byte {
// 		return []byte{163, 164, 178, 70, 146, 133, 15, 232, 52, 165, 190, 193, 139, 93, 77, 255}
// 	}
// 	test1Data := []byte{163, 164, 178, 70, 146, 133, 15, 232, 52, 165, 190, 193, 139, 93, 77, 255, 70, 96, 12, 168, 17, 30, 226, 135, 112, 111, 200, 94, 157, 239, 4, 235, 203, 152, 113, 224, 179, 14, 180, 225, 14, 104, 174, 217, 130, 137, 101, 3, 202, 77, 43, 64, 181, 201, 131, 52, 130, 203, 105, 15, 250, 176, 10, 224}

// 	type args struct {
// 		passphrase []byte
// 		msg        []byte
// 	}
// 	tests := []struct {
// 		name string
// 		args args
// 		want []byte
// 	}{
// 		{
// 			name: "Test 1: thingy",
// 			args: args{passphrase: []byte("Test"), msg: []byte("tst Data")},
// 			want: test1Data,
// 		},
// 	}
// 	for _, tt := range tests {
// 		t.Run(tt.name, func(t *testing.T) {
// 			if got := Encrypt(tt.args.passphrase, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
// 				t.Errorf("Encrypt() = %v, want %v", got, tt.want)
// 			}
// 		})
// 	}
// }

func TestDecrypt(t *testing.T) {

	test1Data := []byte{163, 164, 178, 70, 146, 133, 15, 232, 52, 165, 190, 193, 139, 93, 77, 255, 70, 96, 12, 168, 17, 30, 226, 135, 112, 111, 200, 94, 157, 239, 4, 235, 203, 152, 113, 224, 179, 14, 180, 225, 14, 104, 174, 217, 130, 137, 101, 3, 202, 77, 43, 64, 181, 201, 131, 52, 130, 203, 105, 15, 250, 176, 10, 224}
	type args struct {
		passphrase []byte
		msg        []byte
	}
	tests := []struct {
		name string
		args args
		want []byte
	}{
		{
			name: "Test 1: things",
			args: args{passphrase: []byte("Test"), msg: test1Data},
			want: []byte("tst Data"),
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			if got := Decrypt(tt.args.passphrase, tt.args.msg); !reflect.DeepEqual(got, tt.want) {
				t.Errorf("Decrypt() = %v, want %v", got, tt.want)
			}
		})
	}
}
