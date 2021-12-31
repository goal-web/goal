package tests

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestHashing(t *testing.T) {
	// 初始化容器

	path, _ := os.Getwd()
	app := getApp(path)

	hashFactory := app.Get("hash").(contracts.HasherFactory)
	value := "goal hashing"

	bcryptHashedValue := hashFactory.Make(value, nil)
	fmt.Println("bcryptHashedValue:", bcryptHashedValue)
	assert.True(t, hashFactory.Check(value, bcryptHashedValue, nil))
	assert.True(t, len(bcryptHashedValue) > 10)

	md5HashedValue := hashFactory.Driver("md5").Make(value, nil)
	fmt.Println("md5HashedValue:", md5HashedValue)
	assert.True(t, hashFactory.Driver("md5").Check(value, md5HashedValue, nil))
}
