package main

import (
	"crypto/md5"
	"crypto/sha1"
	"encoding/base64"
	"encoding/hex"
	"fmt"
)

func main() {
	pwd := setPassword(123456, "41dada-3")
}

// 生成密码
func setPassword(str string, salt string) string {
	publicKey := "FR4ehHBBbda414141smNSq0zLV"
	return md5Str(sha1Str(str+salt) + md5Str(publicKey))
}

// md5加密
func md5Str(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// sha1加密
func sha1Str(str string) string {
	w := sha1.New()
	w.Write([]byte(str))
	return hex.EncodeToString(w.Sum(nil))
}

func Base64Encode(str string) string {
	return base64.StdEncoding.EncodeToString([]byte(str))
}
