package aes

import (
	"encoding/base64"
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
