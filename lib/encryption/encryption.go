package encryption

import (
	cryptorand "crypto/rand"

	b64 "encoding/base64"

	"golang.org/x/crypto/argon2"
	"golang.org/x/crypto/chacha20poly1305"
)

var keySize uint32 = 32
var time uint32 = 3
var memory uint32 = 32 * 1024
var threads uint8 = 4
var saltLength = 16

func EncodeB64(msg []byte) string {
	return b64.StdEncoding.EncodeToString(msg)
}
func DecodeB64(msg string) ([]byte, error) {
	return b64.StdEncoding.DecodeString(msg)
}

var generateSalt = func() []byte {
	salt := make([]byte, saltLength)
	if _, err := cryptorand.Read(salt); err != nil {
		panic(err)
	}
	return salt
}

func Encrypt(passphrase, msg []byte) []byte {

	salt := generateSalt()
	key := argon2.Key(passphrase, salt, time, memory, threads, keySize)
	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}

	// Encryption.
	var encryptedMsg []byte
	// Select a random nonce, and leave capacity for the ciphertext.
	nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())
	if _, err := cryptorand.Read(nonce); err != nil {
		panic(err)
	}
	// Encrypt the message and append the ciphertext to the nonce.
	encryptedMsg = aead.Seal(nonce, nonce, msg, nil)
	if err != nil {
		panic(err)
	}
	return append(salt, encryptedMsg...)
}

func Decrypt(passphrase, msg []byte) []byte {

	salt := msg[:saltLength]
	msg = msg[saltLength:]
	key := argon2.Key(passphrase, salt, time, memory, threads, keySize)

	aead, err := chacha20poly1305.NewX(key)
	if err != nil {
		panic(err)
	}
	if len(msg) < aead.NonceSize() {
		panic("ciphertext too short")
	}
	// Split nonce and ciphertext.
	nonce, ciphertext := msg[:aead.NonceSize()], msg[aead.NonceSize():]
	// Decrypt the message and check it wasn't tampered with.
	plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
	if err != nil {
		panic(err)
	}
	return plaintext
}

// func Encrypt(passphrase, msg []byte) {

// 	salt := make([]byte, keySize)
// 	if _, err := cryptorand.Read(salt); err != nil {
// 		panic(err)
// 	}
// 	key := argon2.Key(passphrase, salt, time, memory, threads, keySize)
// 	aead, err := chacha20poly1305.NewX(key)
// 	if err != nil {
// 		panic(err)
// 	}

// 	// Encryption.
// 	var encryptedMsg []byte
// 	{
// 		// Select a random nonce, and leave capacity for the ciphertext.
// 		nonce := make([]byte, aead.NonceSize(), aead.NonceSize()+len(msg)+aead.Overhead())
// 		if _, err := cryptorand.Read(nonce); err != nil {
// 			panic(err)
// 		}
// 		// Encrypt the message and append the ciphertext to the nonce.
// 		encryptedMsg = aead.Seal(nonce, nonce, msg, nil)
// 	}

// 	// Decryption.
// 	{
// 		if len(encryptedMsg) < aead.NonceSize() {
// 			panic("ciphertext too short")
// 		}
// 		// Split nonce and ciphertext.
// 		nonce, ciphertext := encryptedMsg[:aead.NonceSize()], encryptedMsg[aead.NonceSize():]
// 		// Decrypt the message and check it wasn't tampered with.
// 		plaintext, err := aead.Open(nil, nonce, ciphertext, nil)
// 		if err != nil {
// 			panic(err)
// 		}

// 		fmt.Printf("%s\n", plaintext)
// 	}
// }
