package specifications

type GreetAdapter func(name string) string
type CurseAdapter func(name string) string

func (g GreetAdapter) Greet(name string) (string, error) {
	return g(name), nil
}

func (g CurseAdapter) Curse(name string) (string, error) {
	return g(name), nil
}
