package encryptor

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"fmt"
)

const (
	default_salt = "6HnyoiEyXdToZDi74e&FUjdhMJ9tkG5J"
)

type Encryptor interface {
	// Encrypt encrypt data
	Encrypt(data interface{}) (string, error)
	// Decrypt decrypt data
	Decrypt(cipher string) (interface{}, error)
}

type defaultEncryptor struct {
	salt []byte
}

func New(salt string) (*defaultEncryptor, error) {
	if len(salt) == 0 {
		salt = default_salt
	} else if len(salt) != 32 {
		return nil, fmt.Errorf("aes salt must be 32 byte")
	}
	base64Str := base64.StdEncoding.EncodeToString([]byte(salt))
	return &defaultEncryptor{salt: []byte(base64Str)}, nil
}

func (d *defaultEncryptor) Encrypt(data interface{}) (string, error) {
	dataBytes, ok := data.([]byte)
	if !ok {
		return "", fmt.Errorf("data be encrypted not bytes")
	}

	block, _ := aes.NewCipher(d.salt)
	blockSize := block.BlockSize()

	data = _PKCS7Padding(dataBytes, blockSize)
	blockMode := cipher.NewCBCEncrypter(block, d.salt[:blockSize])
	cryted := make([]byte, len(dataBytes))
	blockMode.CryptBlocks(cryted, dataBytes)

	return base64.StdEncoding.EncodeToString(cryted), nil
}

func (d *defaultEncryptor) Decrypt(ciphered string) (interface{}, error) {
	crytedByte, _ := base64.StdEncoding.DecodeString(ciphered)

	block, _ := aes.NewCipher(d.salt)
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, d.salt[:blockSize])

	orig := make([]byte, len(crytedByte))
	blockMode.CryptBlocks(orig, crytedByte)
	return _PKCS7UnPadding(orig), nil
}

func _PKCS7Padding(ciphertext []byte, blocksize int) []byte {
	padding := blocksize - len(ciphertext)%blocksize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func _PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
