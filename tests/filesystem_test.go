package tests

import (
	"fmt"
	"github.com/qbhy/goal/filesystemt"
	"io/fs"
	"testing"
)

func TestLocalFileSystemDriver(t *testing.T) {
	local := filesystemt.NewLocalFileSystem("/Users/qbhy/project/go/goal/storages/", fs.ModePerm)
	fmt.Println(local.MakeDirectory("logs"))
}
