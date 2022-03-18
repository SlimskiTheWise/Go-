package duck

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

var donald Duck
var john Person

func TestDuck_Quack(t *testing.T) {

	assert := assert.New(t)

	r := donald.Quack()

	assert.Equal("Quack", r)

}

func TestDuck_Feathers(t *testing.T) {

	assert := assert.New(t)

	r := donald.Feathers()

	assert.Equal("White Feathers", r)
}

func TestPerson_Quack(t *testing.T) {

	assert := assert.New(t)

	r := john.Quack()

	assert.Equal("A person mocking a duck", r)

}

func TestPerson_Feathers(t *testing.T) {

	assert := assert.New(t)

	r := john.Feathers()

	assert.Equal("A person gathering feathers", r)
	
}
