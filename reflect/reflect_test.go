package reflect

import (
	"github.com/stretchr/testify/assert"
	"reflect"
	"testing"
)

func TestReflectT(t *testing.T) {
	assert := assert.New(t)

	var num int = 1

	r := ReflectT(num)

	assert.Equal(reflect.TypeOf(num), r)
}

func TestGetValueOfPointer(t *testing.T) {
	assert := assert.New(t)

	var a *int = new(int)
	*a = 1

	g := GetValueOfPointer(a)

	assert.Equal(int64(1), g)
}
