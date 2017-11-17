package symmetric

import (
	"crypto/aes"
	"crypto/cipher"
	"encoding/hex"

	"practice/go/util"

	"github.com/mozhata/merr"
)

// key length must be 16
func Encrypt(key, content string) (string, error) {
	if len(key) != aes.BlockSize {
		return "", merr.InvalidErr(nil, "suppose key's length is 16 but not. key: %s", key)
	}
	encrypted, err := aesEcrpt([]byte(key), []byte(content))
	if err != nil {
		return "", err
	}
	return string(encrypted), nil
}

// aesEcrpt key should fixed to 16
func aesEcrpt(key, content []byte) ([]byte, error) {
	util.Debug("len(key): %d, len(content): %d", len(key), len(content))
	var encrypted []byte
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, merr.InvalidErr(err, "key: %s", string(key))
	}
}

func Encrypt(text string, key []byte) (string, error) {
	var iv = key[:aes.BlockSize]
	encrypted := make([]byte, len(text))
	block, err := aes.NewCipher(key)
	if err != nil {
		return "", err
	}
	encrypter := cipher.NewCFBEncrypter(block, iv)
	encrypter.XORKeyStream(encrypted, []byte(text))
	return hex.EncodeToString(encrypted), nil
}
