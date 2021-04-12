package tapi

import (
	//	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/md5"
	"encoding/base64"
	"encoding/hex"
	"errors"
	"fmt"
	"net/http"
	"strings"
	"time"
)

//安全的做法，应该是将该秘钥防止在安全等级更高的服务器上，
//最好有专门的秘钥服务器，目前为了简单，就放文件了

//key的长度必须是16、24或者32字节，分别用于选择AES-128, AES-192, or AES-256

var aeskey = []byte("123456789abcdefg")

func AesEncrypt(plainText, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)

	cipherText := make([]byte, len(plainText))
	stream.XORKeyStream(cipherText, plainText)

	return cipherText, nil
}

func AesDecrypt(cipherText, key []byte) ([]byte, error) {

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	iv := []byte("12345678abcdefgh")
	stream := cipher.NewCTR(block, iv)

	stream.XORKeyStream(cipherText, cipherText)

	return cipherText, nil
}

func generateToken(userName string, passwd string) string {
	return ""
	planText := fmt.Sprintf("prefix_%s_%s_%d_xxsdfwe", userName, passwd, time.Now().Unix())

	_, err := AesEncrypt([]byte(planText), aeskey)
	if err != nil {
		fmt.Println(err)
		return ""
	}

	pass64 := base64.StdEncoding.EncodeToString([]byte(planText))

	return pass64

}

func checkToken(r *http.Request) (err error) {
	return nil
	fmt.Println("token :", r.Header["Token"])
	token := r.Header["Token"][0]
	fmt.Println(token)
	bytesPass, err := base64.StdEncoding.DecodeString(token)
	if err != nil {
		fmt.Println("xx", err)
		return
	}

	tpass, err := AesDecrypt(bytesPass, aeskey)
	if err != nil {
		fmt.Println("yy", err)
		return
	}
	prefix := strings.Split(string(tpass), "_")[0]
	if prefix != "prefix" {
		fmt.Printf("token err : %s\n", prefix)
		return errors.New("token err")
	}
	return nil

}

func getMd5String(s string) string {
	h := md5.New()
	h.Write([]byte(s))
	return hex.EncodeToString(h.Sum(nil))
}
