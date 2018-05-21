package redis

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

func TestInit(t *testing.T) {
	r, err := Init()
	assert.NoError(t, err)
	assert.NotEmpty(t, r)
}

func TestSET(t *testing.T) {
	name := [5]string{"河野拓", "勝山栄一", "上島良平", "福谷雄貴", "高畠斉"}
	for i := 0; i < 5; i++ {
		err := SET("name", name[i])
		assert.NoError(t, err)
	}
}

func TestGET(t *testing.T) {
	result, err := GET("name")
	if assert.NoError(t, err) {
		assert.Equal(t, result, "高畠斉")
	}
}

func TestEXISTS(t *testing.T) {
	result, err := EXISTS("name")
	if assert.NoError(t, err) {
		assert.True(t, result)
	}
}
