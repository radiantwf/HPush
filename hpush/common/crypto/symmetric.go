package crypto

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
	"crypto/des"
	"errors"
)

const (
	AES = byte(iota)
	DES
)

const (
	NOPADDING = byte(iota)
	PKCS5PADDING
	PKCS7PADDING
)

const (
	CBC = byte(iota)
	ECB
	CTR
	OCF
	CFB
)

type CryptionOption struct {
	Algorithm byte
	Padding   byte
	Mode      byte
}

var DefaultCryptionOption = CryptionOption{
	Algorithm: DES,
	Padding:   PKCS5PADDING,
	Mode:      CBC,
}

type SymmetricCrypto struct{}

func (c *SymmetricCrypto) CipherPKCS5Padding(key []byte, option CryptionOption, origData []byte) (cipherData []byte, err error) {
	var block cipher.Block
	switch option.Algorithm {
	case DES:
		block, err = des.NewCipher(key)
	case AES:
		block, err = aes.NewCipher(key)
	default:
		err = errors.New("Cryption algorithm option error")
	}
	if err != nil {
		return
	}
	blockSize := block.BlockSize()

	switch option.Padding {
	case PKCS5PADDING:
		origData = c._PKCS5Padding(origData, blockSize)
	case PKCS7PADDING:
		origData = c._PKCS7Padding(origData, blockSize)
	case NOPADDING:
	default:
		err = errors.New("Cryption padding option error")
	}
	if err != nil {
		return
	}

	var blockMode cipher.BlockMode
	switch option.Mode {
	case CBC:
		blockMode = cipher.NewCBCEncrypter(block, key[:blockSize])
	case CFB:
		fallthrough
	case ECB:
		fallthrough
	case CTR:
		fallthrough
	case OCF:
		err = errors.New("Cryption mode option not support")
	default:
		err = errors.New("Cryption mode option error")
	}
	if err != nil {
		return
	}

	cipherData = make([]byte, len(origData))
	blockMode.CryptBlocks(cipherData, origData)
	return
}

func (c *SymmetricCrypto) DecipherPKCS5Padding(key []byte, option CryptionOption, cipherData []byte) (origData []byte, err error) {
	var block cipher.Block
	switch option.Algorithm {
	case DES:
		block, err = des.NewCipher(key)
	case AES:
		block, err = aes.NewCipher(key)
	default:
		err = errors.New("Cryption algorithm option error")
	}
	if err != nil {
		return
	}
	blockSize := block.BlockSize()

	var blockMode cipher.BlockMode
	switch option.Mode {
	case CBC:
		blockMode = cipher.NewCBCDecrypter(block, key[:blockSize])
	case CFB:
		fallthrough
	case ECB:
		fallthrough
	case CTR:
		fallthrough
	case OCF:
		err = errors.New("Cryption mode option not support")
	default:
		err = errors.New("Cryption mode option error")
	}
	if err != nil {
		return
	}

	origData = make([]byte, len(cipherData))
	blockMode.CryptBlocks(origData, cipherData)

	switch option.Padding {
	case PKCS5PADDING:
		origData = c._PKCS5UnPadding(origData)
	case PKCS7PADDING:
		origData = c._PKCS7UnPadding(origData)
	case NOPADDING:
	default:
		err = errors.New("Cryption padding option error")
	}
	return
}

func (c *SymmetricCrypto) _PKCS5Padding(cipherData []byte, blockSize int) []byte {
	return c._PKCS7Padding(cipherData, blockSize)
}

func (c *SymmetricCrypto) _PKCS5UnPadding(origData []byte) []byte {
	return c._PKCS7UnPadding(origData)
}

func (c *SymmetricCrypto) _PKCS7Padding(cipherData []byte, blockSize int) []byte {
	padding := blockSize - len(cipherData)%blockSize
	padData := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(cipherData, padData...)
}

func (c *SymmetricCrypto) _PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
