package util

import (
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"fmt"
	"github.com/forgoer/openssl"
	"math/big"
)

var iv = []byte("0102030405060708")
var presetKey = []byte("0CoJUm6Qyw8W8jud")
var linuxapiKey = []byte("rFgB&h#%2?^eDg:Q")
var stdChars = []byte("ABCDEFGHIJKLMNOPQRSTUVWXYZabcdefghijklmnopqrstuvwxyz0123456789")
var publicKey = []byte("-----BEGIN PUBLIC KEY-----\nMIGfMA0GCSqGSIb3DQEBAQUAA4GNADCBiQKBgQDgtQn2JZ34ZC28NWYpAUd98iZ37BUrX/aKzmFbt7clFSs6sXqHauqKWqdtLkF2KexO40H1YTX8z2lSgBBOAxLsvaklV8k4cBFK9snQXE9/DDaFt6Rr7iVZMldczhC0JNgTz+SHXT6CBHuX3e9SdB1Ua44oncaTWz7OBGLbCiK45wIDAQAB\n-----END PUBLIC KEY-----")
var eapiKey = []byte("e82ckenh8dichen8")

func aesEncryptCBC(buffer []byte, key []byte, ivv []byte) []byte {
	dst, _ := openssl.AesCBCEncrypt(buffer, key, ivv, openssl.PKCS7_PADDING)
	return dst
	// base64 解码
	//fmt.Println(base64.StdEncoding.EncodeToString(dst))

	// 解密
	//dst, _ = openssl.AesCBCDecrypt(dst, presetKey, iv, openssl.PKCS7_PADDING)
	//fmt.Println(string(dst)) // 123456
}

func aesEncryptECB(buffer []byte, key []byte) []byte {
	dst, _ := openssl.AesECBEncrypt(buffer, key, openssl.PKCS7_PADDING)
	return dst
	//fmt.Println(base64.StdEncoding.EncodeToString(dst))  // yXVUkR45PFz0UfpbDB8/ew==
	// hex 解码
	//fmt.Println(hex.EncodeToString(dst))

	//解密
	//dst, _ = openssl.AesECBDecrypt(dst, linuxapiKey, openssl.PKCS7_PADDING)
	//fmt.Println(string(dst)) // 123456
}

func NewLen16Rand() ([]byte, []byte) {
	randByte := make([]byte, 16)
	randByteReverse := make([]byte, 16)
	for i := 0; i < 16; i++ {
		result, _ := rand.Int(rand.Reader, big.NewInt(62))
		randByte[i] = stdChars[result.Int64()]
		randByteReverse[15-i] = stdChars[result.Int64()]
	}
	return randByte, randByteReverse
}

func aesEncrypt(buffer []byte, mod string, key []byte, ivv []byte) []byte {
	if mod == "cbc" {
		return aesEncryptCBC(buffer, key, ivv)
	} else if mod == "ecb" {
		return aesEncryptECB(buffer, key)
	}
	return nil
}

func rsaEncrypt(buffer []byte, key []byte) []byte {
	buffers := make([]byte, 128-16, 128)
	buffers = append(buffers, buffer...)
	block, _ := pem.Decode(key)
	if block == nil {
		return nil
	}
	pubInterface, err := x509.ParsePKIXPublicKey(block.Bytes)
	if err != nil {
		fmt.Println(err.Error())
	}
	// 类型断言
	pub := pubInterface.(*rsa.PublicKey)

	// 加密 因为网易采用的是无padding加密故直接进行计算
	c := new(big.Int).SetBytes([]byte(buffers))
	encryptedBytes := c.Exp(c, big.NewInt(int64(pub.E)), pub.N).Bytes()
	return encryptedBytes
	////加密
	//a,err:=rsa.EncryptPKCS1v15(rand.Reader, pub, buffers)
	//if err!=nil{
	//	fmt.Println(err.Error())
	//}
	//return a
}

type WeapiType struct {
	Params    string `json:"params"`
	EncSecKey string `json:"encSecKey"`
}

func Weapi(text interface{}) *WeapiType{
	return nil
}
