package aes

import (
	"crypto/aes"
	"crypto/cipher"
	"crypto/rand"
	"io"
)

// AESGCMEncrypt
// 请保存好nonce，解密需要
func AESGCMEncrypt(origData, key []byte) (cipherData, nonce []byte, err error) {
	block, err := aes.NewCipher(key)
	if err != nil {
		return nil, nil, err
	}
	// origData = ZeroPadding(origData, block.BlockSize())
	gcm, err := cipher.NewGCM(block)
	if err != nil {
		return nil, nil, err
	}

	nonce = make([]byte, 12)
	_, err = io.ReadFull(rand.Reader, nonce)
	if err != nil {
		return nil, nil, err
	}

	cipherData = gcm.Seal(nil, nonce, origData, nil)

	return cipherData, nonce, nil
}

func AESGCMDecrypt(cipherData, key, nonce []byte) (plaintextData []byte, err error) {
	aesBlock, err := aes.NewCipher(key)
	if err != nil {
		return nil, err
	}

	gcm, err := cipher.NewGCM(aesBlock)
	if err != nil {
		return nil, err
	}

	plaintextData, err = gcm.Open(nil, nonce, cipherData, nil)
	return
}
