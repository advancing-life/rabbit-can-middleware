package docker

import (
	"bufio"
	"bytes"
	"fmt"
	"github.com/mattn/go-shellwords"
	"os/exec"
)

// ExecutionCommand ...
type ExecutionCommand struct {
	ContainerID string `json:"container_id"`
	Command     string `json:"command"`
	Result      string `json:"result"`
	ExitStatus  string `json:"exit_status"`
}

// cmdrun ...
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

// Mk ...
func Mk(id, lang string) (ID string, err error) {
	switch lang {
	case "c":
		ID, err = cmdrun("docker run --name " + id + " -itd jpnlavender/clang_on_debian /bin/bash")
		return
	case "java":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	case "py":
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	case "rb":
		ID, err = cmdrun("docker run --name " + id + " -itd jpnlavender/ruby_on_debian /bin/bash")
		return
	default:
		ID, err = cmdrun("docker run --name " + id + " -itd debian:jessie /bin/bash")
		return
	}
}

// Rm ...
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

// Exec ...
func Exec(exech chan ExecutionCommand, execmd ExecutionCommand, name string) {
	// go cmdrun(exits, "docker exec -i "+name+" echo $?")

	go func() {
		defer close(exech)
		c, err := shellwords.Parse("docker exec -i " + name + " " + execmd.Command)
		var ec bytes.Buffer

		if err != nil {
			fmt.Print(err)
		}
		ecmd := exec.Command(c[0], c[1:]...)
		ecmd.Stdout = &ec
		stdout, err := ecmd.StdoutPipe()

		if err != nil {
			fmt.Print(err)
		}

		ecmd.Start()

		scanner := bufio.NewScanner(stdout)
		for scanner.Scan() {
			fmt.Printf("%q\n", ec.String())
			execmd.Result = scanner.Text()
			exech <- execmd
		}
		ecmd.Wait()
	}()
}
