package filesystem

import (
	"io/fs"
	"io/ioutil"
)

type file struct {
	fs.FileInfo
	disk string
}

func (this *file) Get() string {
	contents, _ := ioutil.ReadFile(this.Name())
	return string(contents)
}

func (this *file) Disk() string {
	return this.disk
}
