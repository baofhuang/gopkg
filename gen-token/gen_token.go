package main

import (
	"bytes"
	"crypto/md5"
	"encoding/hex"
	"flag"
	"fmt"
	"time"
)

// GenerateToken 生成token,申请时间（精确到纳秒）+ 使用者
func GenerateToken(tokenType, owner string) string {
	if owner == "" {
		return ""
	}
	applyTime := time.Now().UTC().UnixNano()
	return GetStrMd5(tokenType, owner, fmt.Sprint(applyTime))
}

// GetStrMd5 md5加密
func GetStrMd5(str ...string) string {
	if len(str) == 0 {
		return ""
	} else if len(str) == 1 {
		return getStrMd5(str[0])
	} else {
		var buffer bytes.Buffer
		for i := range str {
			buffer.WriteString(str[i])
		}
		return GetByteMd5(buffer.Bytes())
	}
}

// getStrMd5 md5加密
func getStrMd5(str string) string {
	h := md5.New()
	h.Write([]byte(str))
	return hex.EncodeToString(h.Sum(nil))
}

// GetByteMd5 md5加密
func GetByteMd5(b []byte) string {
	h := md5.New()
	h.Write(b)
	return hex.EncodeToString(h.Sum(nil))
}

func main() {
	owner := flag.String("owner", "default", "please input owner")
	tokenType := flag.String("type", "public", "please input token type")
	flag.Parse()
	fmt.Println("token:", GenerateToken(*tokenType, *owner))
}
