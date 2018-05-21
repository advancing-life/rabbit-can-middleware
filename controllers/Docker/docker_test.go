package docker

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestMk(t *testing.T) {
	result, err := Mk("test", "test")
	if err != nil {
		t.Error(err)
	}
	assert.NotEmpty(t, result)
}

func TestExec(t *testing.T) {
	// result, exit_status, err := Exec("test", "date")
	result := make(chan ExecutionCommand)
	execmd := ExecutionCommand{ContainerID: "test", Command: "ls -la"}

	go Exec(result, execmd, "ls -al")

	for i := range result {
		assert.NotEmpty(t, i)
		// assert.Equal(exit_status, "0")
		t.Log(i)
	}
}

func TestRm(t *testing.T) {
	err := Rm("test")
	assert.NoError(t, err)
}
