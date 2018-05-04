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

func MakeContainer() (ID string, err error) {
	ID, err = cmdrun("ls -a -l")
	return
}

/*
func DelContainer() {
}

func Run(cID string) {
}
*/
