package aes

import (
	"bytes"
	"crypto/aes"
	"crypto/cipher"
)

/**
 *key must 16/24/32 length
 */
func AESEncrypt(origData, key []byte, pkcs func(text []byte, size int) []byte) ([]byte, error) {
	if pkcs == nil {
		pkcs = PKCS5Padding
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	origData = pkcs(origData, block.BlockSize())
	// origData = ZeroPadding(origData, block.BlockSize())
	blockMode := cipher.NewCBCEncrypter(block, key[0:16])
	crypted := make([]byte, len(origData))
	// 根据CryptBlocks方法的说明，如下方式初始化crypted也可以
	// crypted := origData
	blockMode.CryptBlocks(crypted, origData)
	return crypted, nil
}

func AESDecrypt(crypted, key []byte, unPkcs func(data []byte) []byte) ([]byte, error) {
	if unPkcs == nil {
		unPkcs = PKCS5UnPadding
	}

	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}
	blockMode := cipher.NewCBCDecrypter(block, key[0:16])
	origData := make([]byte, len(crypted))
	// origData := crypted
	blockMode.CryptBlocks(origData, crypted)
	origData = unPkcs(origData)
	// origData = ZeroUnPadding(origData)
	return origData, nil
}

//使用PKCS7进行填充，IOS也是7
func PKCS7Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS7UnPadding(origData []byte) []byte {
	length := len(origData)
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}

func PKCS5Padding(ciphertext []byte, blockSize int) []byte {
	padding := blockSize - len(ciphertext)%blockSize
	padtext := bytes.Repeat([]byte{byte(padding)}, padding)
	return append(ciphertext, padtext...)
}

func PKCS5UnPadding(origData []byte) []byte {
	length := len(origData)
	// 去掉最后一个字节 unpadding 次
	unpadding := int(origData[length-1])
	return origData[:(length - unpadding)]
}
