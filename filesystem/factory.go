package filesystem

import (
	"bufio"
	"fmt"
	"github.com/goal-web/contracts"
	"github.com/goal-web/supports/utils"
	"github.com/qbhy/goal/logs"
	"io/fs"
	"time"
)

type Factory struct {
	config  contracts.Config
	disks   map[string]contracts.FileSystem
	drivers map[string]contracts.FileSystemProvider
}

func (this Factory) getConfig(name string) contracts.Fields {
	return this.config.GetFields(
		utils.IfString(name == "default", "filesystem", fmt.Sprintf("filesystem.disks.%s", name)),
	)
}

func (this *Factory) Disk(name string) contracts.FileSystem {
	if disk, existsStore := this.disks[name]; existsStore {
		return disk
	}

	this.disks[name] = this.get(name)

	return this.disks[name]
}

func (this *Factory) Extend(driver string, provider contracts.FileSystemProvider) {
	this.drivers[driver] = provider
}

func (this *Factory) get(name string) contracts.FileSystem {
	config := this.getConfig(name)
	drive := utils.GetStringField(config, "driver", "local")
	driveProvider, existsProvider := this.drivers[drive]
	if !existsProvider {
		logs.WithFields(nil).Fatal(fmt.Sprintf("不支持的文件系统驱动：%s", drive))
	}
	return driveProvider(config)
}

func (this *Factory) Name() string {
	return this.Disk("default").Name()
}

func (this *Factory) Exists(path string) bool {
	return this.Disk("default").Exists(path)
}

func (this *Factory) Get(path string) (string, error) {
	return this.Disk("default").Get(path)
}

func (this *Factory) ReadStream(path string) (*bufio.Reader, error) {
	return this.Disk("default").ReadStream(path)
}

func (this *Factory) Put(path, contents string) error {
	return this.Disk("default").Put(path, contents)
}

func (this *Factory) WriteStream(path string, contents string) error {
	return this.Disk("default").WriteStream(path, contents)
}

func (this *Factory) GetVisibility(path string) contracts.FileVisibility {
	return this.Disk("default").GetVisibility(path)
}

func (this *Factory) SetVisibility(path string, perm fs.FileMode) error {
	return this.Disk("default").SetVisibility(path, perm)
}

func (this *Factory) Prepend(path, contents string) error {
	return this.Disk("default").Prepend(path, contents)
}

func (this *Factory) Append(path, contents string) error {
	return this.Disk("default").Append(path, contents)
}

func (this *Factory) Delete(path string) error {
	return this.Disk("default").Delete(path)
}

func (this *Factory) Copy(from, to string) error {
	return this.Disk("default").Copy(from, to)
}

func (this *Factory) Move(from, to string) error {
	return this.Disk("default").Move(from, to)
}

func (this *Factory) Size(path string) (int64, error) {
	return this.Disk("default").Size(path)
}

func (this *Factory) LastModified(path string) (time.Time, error) {
	return this.Disk("default").LastModified(path)
}

func (this *Factory) Files(directory string) []contracts.File {
	return this.Disk("default").Files(directory)
}

func (this *Factory) AllFiles(directory string) []contracts.File {
	return this.Disk("default").AllFiles(directory)
}

func (this *Factory) Directories(directory string) []string {
	return this.Disk("default").Directories(directory)
}

func (this *Factory) AllDirectories(directory string) []string {
	return this.Disk("default").AllDirectories(directory)
}

func (this *Factory) MakeDirectory(path string) error {
	return this.Disk("default").MakeDirectory(path)
}

func (this *Factory) DeleteDirectory(directory string) error {
	return this.Disk("default").DeleteDirectory(directory)
}
