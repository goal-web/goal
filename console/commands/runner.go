package commands

import (
	"fmt"
	"github.com/qbhy/goal/contracts"
	"github.com/qbhy/goal/logs"
	"io/ioutil"
	"os"
)

type runner struct {
	Base
	app contracts.Application
}

func Runner(app contracts.Application) contracts.Command {
	return &runner{
		Base: BaseCommand("run {--port:http端口号=8000}", "启动 goal"),
		app:  app,
	}
}

func (this *runner) Handle(arguments contracts.CommandArguments) interface{} {
	path := this.app.Get("path").(string)
	pidPath := path + "/goal.pid"
	// 写入 pid 文件
	_ = ioutil.WriteFile(pidPath, []byte(fmt.Sprintf("%d", os.Getpid())), os.ModePerm)

	if errors := this.app.Start(); len(errors) > 0 {
		logs.WithField("errors", errors).Fatal("goal 启动异常!")
	} else {
		_ = os.Remove(pidPath)
		logs.WithInterface(nil).Info("goal 已关闭")
	}
	return nil
}
