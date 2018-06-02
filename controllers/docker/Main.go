package docker

import (
	"bufio"
	"github.com/mattn/go-shellwords"
	"os/exec"
)

// ExecutionCommand ...
type ExecutionCommand struct {
	ContainerID string `json:"container_id"`
	Command     string `json:"command"`
	Result      string `json:"result"`
	ErrResult   error  `json:"error_result"`
	// ExitStatus  string `json:"exit_status"`
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
	go func() {
		defer close(exech)
		c, err := shellwords.Parse("docker exec -i " + name + " sh -c '" + execmd.Command + "'")

		if err != nil {
			execmd.ErrResult = err
			exech <- execmd
		}

		cmd := exec.Command(c[0], c[1:]...)

		stdout, err := cmd.StdoutPipe()
		if err != nil {
			execmd.ErrResult = err
			exech <- execmd
		}
		stderr, err := cmd.StderrPipe()
		if err != nil {
			execmd.ErrResult = err
			exech <- execmd
		}

		err = cmd.Start()
		if err != nil {
			execmd.ErrResult = err
			exech <- execmd
		}

		streamReader := func(scanner *bufio.Scanner, outputChan chan string, doneChan chan bool) {
			defer close(outputChan)
			defer close(doneChan)
			for scanner.Scan() {
				outputChan <- scanner.Text()
			}
			doneChan <- true
		}

		stdoutScanner := bufio.NewScanner(stdout)
		stdoutOutputChan := make(chan string)
		stdoutDoneChan := make(chan bool)
		stderrScanner := bufio.NewScanner(stderr)
		stderrOutputChan := make(chan string)
		stderrDoneChan := make(chan bool)
		go streamReader(stdoutScanner, stdoutOutputChan, stdoutDoneChan)
		go streamReader(stderrScanner, stderrOutputChan, stderrDoneChan)

		stillGoing := true
		for stillGoing {
			select {
			case <-stdoutDoneChan:
				stillGoing = false
			case line := <-stdoutOutputChan:
				execmd.Result = line
				exech <- execmd
			case line := <-stderrOutputChan:
				execmd.Result = line
				exech <- execmd
			}
		}

		ret := cmd.Wait()
		if ret != nil {
			execmd.ErrResult = err
			exech <- execmd
		}

	}()
}
