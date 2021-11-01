package filesystemt

import (
	"io/fs"
	"io/ioutil"
)

type file struct {
	fs.FileInfo
}

func (this *file) Get() string {
	contents, _ := ioutil.ReadFile(this.Name())
	return string(contents)
}
