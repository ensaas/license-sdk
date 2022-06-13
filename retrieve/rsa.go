package retrieve

import (
	"crypto"
	"crypto/rsa"
	"crypto/sha256"
	"crypto/x509"
	"encoding/pem"
	"errors"
	"log"
)

const (
	_RSA_PUBLIC_KEY = `-----BEGIN PUBLIC KEY-----
MIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQC74Y32RFKZcpPjTYLrU0SoaNW7
MJBQ8l5sxRzjPNg+goJy+T5WiPZ1OQsH3OHqBDZGYnhEZZAcxFjH9E7T0riMmzMz
K2AZgyEu1Vp+jWtKP7E5K4NQjujeze4VUoO9Z524EAPab7il2qIZ9rMrhymnl/Si
EOtaW+YR5xQX0seJLwIDAQAB
-----END PUBLIC KEY-----`
)

//public key validate
func RsaVerySignWithSha256(data, signData []byte) (bool, error) {
	block, _ := pem.Decode([]byte(_RSA_PUBLIC_KEY))
	if block == nil {
		log.Println(errors.New("public key error"))
		return false, errors.New("public key error")
	}

	pubKey, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		log.Println("Parse key err:", err)
		return false, err
	}

	hashed := sha256.Sum256(data)
	err = rsa.VerifyPKCS1v15(pubKey.(*rsa.PublicKey), crypto.SHA256, hashed[:], signData)
	if err != nil {
		log.Println("Validate key err:", err)
		return false, err
	}
	return true, nil
}
