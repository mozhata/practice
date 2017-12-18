package symmetric

import (
	"crypto/aes"
	"crypto/cipher"

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
	var encrypted []byte = make([]byte, 0, len(content))
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, merr.InvalidErr(err, "key: %s", string(key))
	}
	/*
	   func NewCBCDecrypter(b Block, iv []byte) BlockMode
	   func NewCBCEncrypter(b Block, iv []byte) BlockMode
	*/
	encrypter := cipher.NewCBCEncrypter(block, key)
	util.Debug("block size: %d", encrypter.BlockSize())
	encrypter.CryptBlocks(encrypted, content)
	util.Debug("encrypted: %d\n%s\ncontent: %d\n%s", len(encrypted), encrypted, len(content), content)
	encrypter.CryptBlocks(content, content)
	util.Debug("content: %d\n%s", len(content), content)
	return encrypted, nil
}
