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
		fmt.Printf("\x1b[31m%s\x1b[0m", err)
	}
	fmt.Printf("\x1b[35mresult:\x1b[0m \n\x1b[31m%s\x1b[0m", string(out))
	Result = string(out)
	return
}

func Mk(id string, lang string) (ID string, err error) {
	switch lang {
	case "c":
		ID, err = cmdrun("echo 'select clang'")
		return
	case "java":
		ID, err = cmdrun("echo 'select java'")
		return
	case "py":
		ID, err = cmdrun("echo 'select py'")
		return
	case "rb":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	default:
		ID, err = cmdrun("echo 'select non'")
		return
	}
}

func Rm(ID string) (err error) {
	_, err = cmdrun("docker stop " + ID)
	if err != nil {
		return
	}
	_, err = cmdrun("docker rm " + ID)
	if err != nil {
		return
	}
	return
}

func Exec(ID string, cmd string) (err error) {
	_, err = cmdrun("docker exec -it " + cmd)
	return
}
