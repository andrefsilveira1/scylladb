package uuid

type Generator interface {
	Generate() string
	Parse(uuid string) error
}

func New() Generator {
	return &generator{}
}
