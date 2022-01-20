package tests

import (
	"fmt"
	"github.com/goal-web/encryption"
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestAESEncryptor(t *testing.T) {
	encryptor := encryption.AES("123456781234567812345678")
	encrypted := encryptor.Encode("goal")

	fmt.Printf("加密后的数据：%s\n", encrypted)
	assert.True(t, encrypted != "")

	decrypted, err := encryptor.Decode(encrypted)
	assert.Nil(t, err)
	assert.True(t, decrypted == "goal")
	fmt.Printf("解密后的数据：%s\n", decrypted)
}
