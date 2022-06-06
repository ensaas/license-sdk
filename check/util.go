package check

import (
	"crypto/md5"
	"encoding/hex"
	"io"
	"math"
	"strings"
)

var num2char = "0123456789abcdefghijklmnopqrstuvwxyz"

func Md5SumString(input string) (string, error) {
	buffer := strings.NewReader(input)
	return Md5Sum(buffer)
}

func Md5Sum(reader io.Reader) (string, error) {
	var returnMD5String string
	hash := md5.New()
	if _, err := io.Copy(hash, reader); err != nil {
		return returnMD5String, err
	}
	hashInBytes := hash.Sum(nil)[:16]
	returnMD5String = hex.EncodeToString(hashInBytes)
	return returnMD5String, nil
}

func BHex2Num(str string, n int) int {
	str = strings.ToLower(str)
	v := 0.0
	length := len(str)
	for i := 0; i < length; i++ {
		s := string(str[i])
		index := strings.Index(num2char, s)
		v += float64(index) * math.Pow(float64(n), float64(length-1-i)) // 倒序
	}
	return int(v)
}

func Num2BHex(number int, n int) string {
	if number < 36 {
		return num2char[number : number+1]
	}
	var merchant, remainder = number / n, number % n
	base36Encode := num2char[remainder : remainder+1]
	for {
		if merchant != 0 {
			merchant, remainder = merchant/n, merchant%n
			base36Encode = num2char[remainder:remainder+1] + base36Encode
		} else {
			return base36Encode
		}
	}
}

func Lpad(str string, totallen int, char byte) string {
	length := len(str)
	if totallen < length {
		return str
	}
	result := str
	for i := 0; i < totallen-length; i++ {
		result = string(char) + result
	}
	return result
}

func Sub(a, b int) int {
	return a - b
}

func Add(a, b int) int {
	return a + b
}
