package tests

import (
	"github.com/qbhy/goal/contracts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCacheFactory(t *testing.T) {
	path, _ := os.Getwd()
	app := getApp(path)

	cacheFactory := app.Get("cache").(contracts.CacheFactory)

	err := cacheFactory.Store().Forever("a", "testing")
	assert.Nil(t, err, err)
	assert.True(t, cacheFactory.Store().Get("a") == "testing")
}
