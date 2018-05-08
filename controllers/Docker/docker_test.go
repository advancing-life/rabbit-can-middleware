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
	result, err := Exec("test", "date")
	assert.NotEmpty(t, result)
	assert.NoError(t, err)
}

func TestRm(t *testing.T) {
	err := Rm("test")
	assert.NoError(t, err)
}
