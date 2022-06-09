package encryptor

import (
	"encoding/base64"
	"fmt"
	"testing"
)

func TestDefaultEncryptor_Decrypt(t *testing.T) {
	s := "123456781234567812345678"
	fmt.Println(base64.StdEncoding.EncodeToString([]byte(s)))
}
