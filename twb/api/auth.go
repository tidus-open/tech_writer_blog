package tapi

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/hex"
	"fmt"
	"time"
)

//安全的做法，应该是将该秘钥防止在安全等级更高的服务器上，
//最好有专门的秘钥服务器，目前为了简单，就放文件了
var twbKey = []byte("qwer823589x")

func aesEncrypt(plainText, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)

	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText
}

func aesDecrypt(cipherText, key []byte) []byte {

	block, err := aes.NewCipher(key)
	if err != nil {
		panic(err)
	}

	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return cipherText
}

func generateToken(userName string, passwd string) string {
	planText := fmt.Sprintf("%s %s %d xxsdfwe", userName, passwd, time.Now().Unix())
	fmt.Println(planText)
	return planText

}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
