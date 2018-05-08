package Docker

import (
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
		out, err = exec.Command(c[0], c[1:]...).CombinedOutput()
	}

	if err != nil {
		return
	}
	Result = string(out)
	return
}

func Mk(id string, lang string) (ID string, err error) {
	switch lang {
	case "c":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	case "java":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	case "py":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	case "rb":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	default:
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
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

func Exec(name string, cmd string) (Result string, err error) {
	Result, err = cmdrun("docker exec -i " + name + " " + cmd)
	return
}
