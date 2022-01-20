package tests

import (
	"fmt"
	"github.com/goal-web/supports/utils"
	"os"
	"path/filepath"
	"testing"
)

func TestCurrentPath(t *testing.T) {
	path, _ := os.Getwd()
	fmt.Println("当前目录: " + path)

	filesPath, _ := filepath.Glob(path + "/../*.env")

	for _, file := range filesPath {
		fields, _ := utils.LoadEnv(file, "=")
		fmt.Println(fields)
	}
}
