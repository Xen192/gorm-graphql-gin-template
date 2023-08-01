package encrypt

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"os"

	"github.com/sirupsen/logrus"
)

var bytes = []byte{79, 12, 46, 52, 96, 35, 23, 13, 42, 16, 49, 73, 75, 52, 63, 34}

func GetHash(text string) string {
	h := sha1.New()
	h.Write([]byte(text))
	return hex.EncodeToString(h.Sum(nil))
}
func Encode(b []byte) string {
	return base64.RawURLEncoding.EncodeToString(b)
}
func Decode(s string) []byte {
	data, err := base64.RawURLEncoding.DecodeString(s)
	if err != nil {
		logrus.Error(err)
	}
	return data
}

func Encrypt(text string) (string, error) {
	secret := os.Getenv("ENCRYPTION_SECRET")
	if secret == "" {
		return "", errors.New("Encryption Secret Not Found")
	}
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	plainText := []byte(text)
	cfb := cipher.NewCFBEncrypter(block, bytes)
	cipherText := make([]byte, len(plainText))
	cfb.XORKeyStream(cipherText, plainText)
	return Encode(cipherText), nil
}

func Decrypt(text string) (string, error) {
	secret := os.Getenv("ENCRYPTION_SECRET")
	if secret == "" {
		return "", errors.New("Encryption Secret Not Found")
	}
	block, err := aes.NewCipher([]byte(secret))
	if err != nil {
		return "", err
	}
	cipherText := Decode(text)
	cfb := cipher.NewCFBDecrypter(block, bytes)
	plainText := make([]byte, len(cipherText))
	cfb.XORKeyStream(plainText, cipherText)
	return string(plainText), nil
}
