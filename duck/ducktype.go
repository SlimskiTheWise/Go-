package duck

type Duck struct {
}

func (d Duck) Quack() string {
	return "Quack"
}

func (d Duck) Feathers() string {
	return "White Feathers"
}

type Person struct {
}

func (p Person) Quack() string {
	return "A person mocking a duck"
}

func (p Person) Feathers() string {
	return "A person gathering feathers"
}

type Quacker interface {
	quack()
	feathers()
}
