package tests

import (
	"fmt"
	"github.com/qbhy/goal/console/commands"
	"os"
	"testing"
)

func TestParseSignature(t *testing.T) {
	fmt.Println(commands.ParseSignature("mail:send {user:用户ID} {title:标题?} {body:内容=hello world} {--queue}"))
}

func TestParseCommandLineArgs(t *testing.T) {
	fmt.Println(os.Args)
}
