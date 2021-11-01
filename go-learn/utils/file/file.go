package main

import (
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"path"
	"strings"
)

func main() {
	str := "hello.txt"
	ext := GetFileExt(str)
	fmt.Println(ext)
	fmt.Println(GetFileName(str))
	fmt.Println(EncodeMD5(GetFileName(str)))

}

func GetFileName(str string) string {
	ext := GetFileExt(str)
	fileName := strings.TrimSuffix(str, ext)
	return fileName
}

func GetFileExt(str string) string {
	return path.Ext(str)
}

func EncodeMD5(value string) string {
	m := md5.New()
	m.Write([]byte(value))
	return hex.EncodeToString(m.Sum(nil))
}
