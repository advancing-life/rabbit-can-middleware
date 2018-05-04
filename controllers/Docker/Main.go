package Docker

import (
	"fmt"
	"github.com/mattn/go-shellwords"
	"os/exec"
)

func cmdrun(input string) (Result string, err error) {
	var out []byte
	c, err := shellwords.Parse(input)
	if err != nil {
		return
	}
	switch len(c) {
	case 0: // 空の文字列が渡された場合
		return
	case 1: // コマンドのみを渡された場合
		err = exec.Command(c[0]).Run()
	default: // コマンド+オプションを渡された場合
		out, err = exec.Command(c[0], c[1:]...).Output()
	}

	if err != nil {
		fmt.Println("Command Exec Error.")
	}
	fmt.Printf("ls result: \n%s", string(out))
	Result = string(out)
	return
}

func MakeContainer(lang string) (ID string, err error) {
	switch lang {
	case "c":
		ID, err = cmdrun("echo 'select clang'")
		return
	case "java":
		ID, err = cmdrun("echo 'select java'")
		return
	case "python":
		ID, err = cmdrun("echo 'select python'")
		return
	case "rb":
		ID, err = cmdrun("echo 'select rb'")
		return
	default:
		ID, err = cmdrun("echo 'select non'")
		return
	}
}

/*
func DelContainer() {
}

func Run(cID string) {
}
*/
