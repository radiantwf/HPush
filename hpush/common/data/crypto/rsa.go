package crypto

import (
	"bytes"
	"crypto/rand"
	"crypto/rsa"
	"crypto/x509"
	"encoding/pem"
	"errors"
)

const ()

func RsaCipher(publickey string, origData []byte) (cipherData []byte, err error) {
	block, _ := pem.Decode([]byte(publickey))
	if block == nil {
		err = errors.New("public key error")
		return
	}
	pubInterface, err1 := x509.ParsePKIXPublicKey(block.Bytes)
	if err1 != nil {
		err = err1
		return
	}
	pub := pubInterface.(*rsa.PublicKey)
	cipherData, err = rsa.EncryptPKCS1v15(rand.Reader, pub, origData)
	return
}

func RsaDecipher(privatekey string, cipherData []byte) (origData []byte, err error) {
	block, _ := pem.Decode([]byte(privatekey))
	if block == nil {
		err = errors.New("private key error")
		return
	}
	priv, err1 := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err1 != nil {
		err = err1
		return
	}
	origData, err = rsa.DecryptPKCS1v15(rand.Reader, priv, cipherData)
	return
}

func GenRsaPrivateKey() (privatekey string, err error) {
	var buf bytes.Buffer
	// 生成私钥文件
	privateKey, err1 := rsa.GenerateKey(rand.Reader, 2048)
	if err1 != nil {
		err = err1
		return
	}
	derStream := x509.MarshalPKCS1PrivateKey(privateKey)
	block := &pem.Block{
		Type:  "PRIVATE KEY",
		Bytes: derStream,
	}
	err = pem.Encode(&buf, block)
	if err != nil {
		return
	}
	privatekey = buf.String()
	return
}

func GenRsaPublicKey(privatekey string) (publickey string, err error) {
	var buf bytes.Buffer
	block, _ := pem.Decode([]byte(privatekey))
	if block == nil {
		err = errors.New("private key error! ")
		return
	}
	priv, err1 := x509.ParsePKCS1PrivateKey(block.Bytes)
	if err1 != nil {
		err = err1
		return
	}
	// 生成公钥文件
	publicKey := &priv.PublicKey
	derPkix, err1 := x509.MarshalPKIXPublicKey(publicKey)
	if err1 != nil {
		err = err1
		return
	}
	block = &pem.Block{
		Type:  "PUBLIC KEY",
		Bytes: derPkix,
	}
	err = pem.Encode(&buf, block)
	if err != nil {
		return
	}
	publickey = buf.String()
	return
}
