package main

import (
	"flag"
	"fmt"
	"github.com/goal-web/goal-cli/app/gen"
	"os"
	"path/filepath"
)

func main() {
	// 定义命令行参数
	out := flag.String("out", "app", "output")                   // 输出目录
	template := flag.String("template", "template.tmpl", "模板文件") // 模板文件
	dir := flag.String("dir", "pro", "指定扫描的 proto 文件目录")         // 要扫描的目录

	flag.Parse()

	// 检查是否指定了 proto 文件目录
	if *dir == "" {
		fmt.Println("请指定要扫描的 proto 文件目录, 使用 --dir 参数")
		os.Exit(1)
	}

	// 获取指定目录中的所有 proto 文件
	protoFiles, err := scanProtoFiles(*dir)
	if err != nil {
		fmt.Printf("扫描目录 %s 中的 proto 文件失败: %v\n", *dir, err)
		os.Exit(1)
	}

	// 遍历所有找到的 proto 文件，依次调用 gen.Pro()
	for _, protoFile := range protoFiles {
		fmt.Printf("正在处理 proto 文件: %s\n", protoFile)
		gen.Pro(protoFile, *template, *out)
	}
}

// scanProtoFiles 扫描指定目录，返回所有的 .proto 文件路径
func scanProtoFiles(root string) ([]string, error) {
	var protoFiles []string

	// 使用 filepath.Walk 遍历目录，查找所有 .proto 文件
	err := filepath.Walk(root, func(path string, info os.FileInfo, err error) error {
		if err != nil {
			return err
		}

		// 如果是 .proto 文件，则加入到列表中
		if !info.IsDir() && filepath.Ext(path) == ".proto" {
			protoFiles = append(protoFiles, path)
		}

		return nil
	})

	if err != nil {
		return nil, err
	}

	return protoFiles, nil
}
