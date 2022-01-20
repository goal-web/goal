package tests

import (
	"github.com/goal-web/filesystem"
	"github.com/goal-web/supports/utils"
	"github.com/stretchr/testify/assert"
	"io/fs"
	"os"
	"testing"
	"time"
)

func TestLocalFileSystemDriver(t *testing.T) {
	path, err := os.Getwd()
	assert.Nil(t, err, err)
	diskPath := path + "/" + utils.Md5(time.Now().Format("2006-01-02 15:04:05.999999999 -0700 MST"))
	disk := filesystem.NewLocalFileSystem("testing", diskPath, fs.ModePerm)

	checkFileContent := func(path, contents string) {
		fileContents, err := disk.Get(path)
		assert.Nil(t, err)
		assert.True(t, fileContents == contents, fileContents)
	}

	// 测试创建文件夹
	assert.Nil(t, disk.MakeDirectory("logs"))

	logPath := "logs/goal.log"
	// 测试写文件
	assert.Nil(t, disk.Put(logPath, "滴滴滴"))

	// 测试检查文件是否存在
	assert.True(t, disk.Exists(logPath))
	assert.False(t, disk.Exists(logPath+"xxx"))

	// 测试获取文件内容
	checkFileContent(logPath, "滴滴滴")

	reader, readErr := disk.ReadStream(logPath)
	assert.Nil(t, readErr)
	line, b, readLineErr := reader.ReadLine()
	assert.Nil(t, readLineErr)
	assert.True(t, !b && string(line) == "滴滴滴")

	// 测试写入文件，缓存方式，已存在的文件
	assert.Nil(t, disk.WriteStream(logPath, "啦啦啦"))
	// 不存在的文件
	assert.Nil(t, disk.WriteStream(logPath+"x", "啦啦啦"))
	checkFileContent(logPath, "啦啦啦")

	// 测试插入内容到文件头部
	assert.Nil(t, disk.Prepend(logPath, "goal"))
	assert.Nil(t, disk.Prepend(logPath+"prepend", "goal"))
	// 验证文件内容
	checkFileContent(logPath, "goal啦啦啦")
	checkFileContent(logPath+"prepend", "goal")

	// 测试追加内容到文件尾部
	assert.Nil(t, disk.Append(logPath, "append"))
	assert.Nil(t, disk.Append(logPath+"append", "append"))
	// 验证文件内容
	checkFileContent(logPath, "goal啦啦啦append")
	checkFileContent(logPath+"append", "append")

	// 删除文件
	assert.Nil(t, disk.Delete(logPath+"prepend"))
	assert.Nil(t, disk.Delete(logPath+"append"))
	// 验证文件是否被删除
	assert.False(t, disk.Exists(logPath+"prepend"))
	assert.False(t, disk.Exists(logPath+"append"))

	// 测试复制文件
	assert.Nil(t, disk.Copy(logPath, logPath+".copy"))
	assert.True(t, disk.Exists(logPath+".copy"))
	checkFileContent(logPath+".copy", "goal啦啦啦append")

	// 测试移动文件
	assert.Nil(t, disk.Move(logPath+".copy", logPath+".move"))
	assert.True(t, disk.Exists(logPath+".move"))
	assert.False(t, disk.Exists(logPath+".copy"))
	checkFileContent(logPath+".move", "goal啦啦啦append")

	// 测试获取文件大小
	logSize, sizeErr := disk.Size(logPath)
	assert.Nil(t, sizeErr)
	assert.True(t, logSize == 19)

	// 测试获取文件大小
	_, lastModifiedErr := disk.LastModified(logPath)
	assert.Nil(t, lastModifiedErr)

	// 获取目录下所有文件
	files := disk.Files("logs")
	assert.True(t, len(files) == 3)
	assert.Nil(t, disk.MakeDirectory("logs/subdir"))
	putErr := disk.Put("logs/subdir/goal.log", "goal")
	assert.Nil(t, putErr, putErr)
	assert.True(t, len(disk.Files("logs")) == 3)    // 不包含子目录的文件数
	assert.True(t, len(disk.AllFiles("logs")) == 4) // 包含子目录的文件数

	// 获取所有目录
	assert.Nil(t, disk.MakeDirectory("logs/subdir/subsubdir"))
	assert.True(t, len(disk.Directories("logs")) == 1)    // 不包含子目录的目录数
	assert.True(t, len(disk.AllDirectories("logs")) == 2) // 包含子目录的目录数

	// 测试文件可见性
	assert.True(t, disk.GetVisibility(logPath) == filesystem.VISIBLE)
	// 测试设置文件权限
	assert.Nil(t, disk.SetVisibility(logPath, fs.FileMode(0400)))
	assert.True(t, disk.GetVisibility(logPath) == filesystem.INVISIBLE)

	// 删除logs文件夹
	assert.Nil(t, disk.DeleteDirectory("logs"))
	//测试删除所有文件
	assert.Nil(t, disk.DeleteDirectory(""))
}
