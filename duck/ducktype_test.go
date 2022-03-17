package duck

import "testing"

var donald Duck
var john Person

func TestDuck_Quack(t *testing.T) {
	r := donald.Quack()

	if r != "Quack" {
		t.Errorf("error %s", r)
	}

}

func TestDuck_Feathers(t *testing.T) {
	r := donald.Feathers()

	if r != "White Feathers" {
		t.Errorf("error %s", r)
	}
}

func TestPerson_Quack(t *testing.T) {
	r := john.Quack()

	if r != "A person mocking a duck" {
		t.Errorf("error %s", r)
	}
}

func TestPerson_Feathers(t *testing.T) {
	r := john.Feathers()

	if r != "A person gathering feathers" {
		t.Errorf("error %s", r)
	}
}
