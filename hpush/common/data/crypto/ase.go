package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

const ()

func AesCipher(key []byte, origData []byte) (cipherData []byte, err error) {
	block, err1 := aes.NewCipher(key)
	if err1 != nil {
		err = err1
		return
	}
	blockSize := block.BlockSize()
	origData = PKCS5Padding(origData, blockSize)
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[:blockSize])
	cipherData = make([]byte, len(origData))
	blockMode.CryptBlocks(cipherData, origData)
	return
}

func AesDecipher(key []byte, cipherData []byte) (origData []byte, err error) {
	block, err1 := aes.NewCipher(key)
	if err1 != nil {
		err = err1
		return
	}
	blockSize := block.BlockSize()
	blockMode := cipher.NewCBCDecrypter(block, key[:blockSize])
	origData = make([]byte, len(cipherData))
	blockMode.CryptBlocks(origData, cipherData)
	origData = PKCS5UnPadding(origData)
	return
}

func PKCS5Padding(cipherData []byte, blockSize int) []byte {
	padding := blockSize - len(cipherData)%blockSize
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherData, padData...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
