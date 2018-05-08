package Docker

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
	result, exit_status, err := Exec("test", "date")
	if assert.NoError(t, err) {
		assert.NotEmpty(t, result)
		assert.Equal(exit_status, "0")
	}
}

func TestRm(t *testing.T) {
	err := Rm("test")
	assert.NoError(t, err)
}
