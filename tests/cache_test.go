package tests

import (
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
)

func TestCacheFactory(t *testing.T) {
	path, _ := os.Getwd()
	app := getApp(path)

	cacheFactory := app.Get("cache").(contracts.CacheFactory)

	fmt.Println(cacheFactory.Store())
	err := cacheFactory.Store().Forever("a", "testing")
	assert.Nil(t, err, err)
	fmt.Println(cacheFactory.Store().Get("a"))
	assert.True(t, cacheFactory.Store().Get("a") == "testing")
}
