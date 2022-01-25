package tests

import (
	"fmt"
	"github.com/goal-web/console/commands"
	"github.com/goal-web/console/inputs"
	"github.com/goal-web/console/scheduling"
	"github.com/golang-module/carbon/v2"
	"github.com/gorhill/cronexpr"
	"github.com/stretchr/testify/assert"
	"os"
	"testing"
	"time"
)

func TestParseSignature(t *testing.T) {
	cmd, args := commands.ParseSignature("mail:send {user:用户ID} {title:标题?} {body:内容=hello world} {--queue:队列}")
	assert.True(t, cmd == "mail:send")
	assert.True(t, len(args) == 4)

	for _, arg := range args {
		switch arg.Name {
		case "user":
			assert.True(t, arg.Description == "用户ID")
			assert.True(t, arg.Type == commands.RequiredArg)
		case "title":
			assert.True(t, arg.Description == "标题")
			assert.True(t, arg.Type == commands.OptionalArg)
		case "body":
			assert.True(t, arg.Description == "内容")
			assert.True(t, arg.Type == commands.OptionalArg)
			assert.True(t, arg.Default == "hello world")
		case "queue":
			assert.True(t, arg.Description == "队列")
			assert.True(t, arg.Type == commands.Option)
		}
	}

}
func TestCommandArguments(t *testing.T) {
	input := inputs.StringArray([]string{"mail:send", "1", "标题", "body", "--queue=redis"})
	args := input.GetArguments()
	assert.True(t, input.GetCommand() == "mail:send")

	assert.True(t, args.GetArg(0) == "1")
	assert.True(t, args.GetArg(1) == "标题")
	assert.True(t, args.GetArg(2) == "body")
	assert.True(t, args.StringOption("queue", "") == "redis")
}

func TestParseCommandLineArgs(t *testing.T) {
	fmt.Println(os.Args)
}

type TestCommand struct {
	commands.Base
	t *testing.T
}

func NewTestCommand(t *testing.T) TestCommand {
	return TestCommand{
		Base: commands.BaseCommand("tests {name} {avatar?} {--queue=redis} {--daemon} {--v}", "测试一下命令行"),
		t:    t,
	}
}

func (this *TestCommand) Handle() interface{} {
	assert.True(this.t, this.GetString("name") == "名称")
	assert.True(this.t, this.GetString("avatar") == "头像")
	assert.True(this.t, this.GetString("queue") == "队列")
	assert.True(this.t, this.GetBool("daemon") == true)
	assert.True(this.t, this.GetBool("v") == false)
	return nil
}

func TestCommandCase(t *testing.T) {
	command := NewTestCommand(t)
	array := inputs.StringArray([]string{"tests", "名称", "头像", "--queue=队列", "--daemon"})
	args := array.GetArguments()
	assert.True(t, command.InjectArguments(args) == nil)
	command.Handle()
}

func TestCronParse(t *testing.T) {
	event := scheduling.NewEvent(nil, func() {}, "")

	event.EveryMinute().EverySecond()

	nextTime := cronexpr.MustParse(event.Expression()).Next(time.Now())

	fmt.Println(event.Expression(), nextTime, carbon.Now().DiffInSeconds(carbon.Time2Carbon(nextTime)))
}
