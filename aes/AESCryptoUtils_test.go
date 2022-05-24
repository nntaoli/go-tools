package aes

import (
	"crypto/sha256"
	"encoding/base64"
	"encoding/hex"
	"testing"
)

var key = []byte("5CqIX1xi&7oNw6th") // 16

func Test_DESEncrypt(t *testing.T) {
	data := []byte("hello")
	secretText, err := AESEncrypt(data, key, PKCS5Padding)

	if err != nil {
		t.Error(err)
	}
	t.Log(secretText)
	//output vkYbld7ogr5nNi/bRp6XMA==
	//assert.Equal(t, base64.StdEncoding.EncodeToString(secretText), "vkYbld7ogr5nNi/bRp6XMA==")

}

func Test_DESDecrypt(t *testing.T) {
	t.Log("des decrypt...")
	data, _ := base64.StdEncoding.DecodeString("TdtrZC+GOoZxZIh4aZZ3csUpasongOi/zc+QJvxjRF77yjVeC8mZ2CkaqUXTkJWG")

	cleartext, err := AESDecrypt(data, key, nil)
	if err != nil {
		t.Error(err)
	}

	t.Log(string(cleartext))
	//assert.Equal(t, string(cleartext), "hello")
}

func Test_DesDec(t *testing.T) {
	data, err := base64.StdEncoding.DecodeString("xJvlRhR5owcis83V1iI/vdhgD/Y2HtVCwLnVhVS5yoNWeRbq//P1C8X6ErdTC681OH5f6MXy5baay1hzv06zGRRzTzLHObzYANZnUkwCYAY=")
	t.Log(err)
	t.Log(len(data))
	sha := sha256.New()
	sha.Write([]byte("x^UV4xG{W90n/hello-pwd"))
	clearTextData, err := AESDecrypt(data, sha.Sum(nil), PKCS5UnPadding)
	t.Log(err)
	t.Log(string(clearTextData))
}

func TestAESGCMEncrypt(t *testing.T) {
	data := "hello world"
	ciphertext, nonce, err := AESGCMEncrypt([]byte(data), key)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(base64.StdEncoding.EncodeToString(ciphertext))
	t.Log(hex.EncodeToString(nonce))
}

func TestAESGCMDecrypt(t *testing.T) {
	data, _ := base64.StdEncoding.DecodeString("CBbPUPs0UKeXVLW6XZHDSUlbs4UUO/VGTBBf")
	nonce, _ := hex.DecodeString("00121c12bca716afd75b25b2")
	plaintext, err := AESGCMDecrypt(data, key, nonce)
	if err != nil {
		t.Fatal(err)
	}
	t.Log(string(plaintext))
}
