package specifications

import (
	"testing"

	"github.com/alecthomas/assert/v2"
)

type Greeter interface {
	Greet(name string) (string, error)
}

func GreetSpecification(t testing.TB, greeter Greeter) {
	got, err := greeter.Greet("Anuj")
	assert.NoError(t, err)
	assert.Equal(t, got, "Hello, Anuj")
}

type MeanGreeter interface {
	Curse(name string) (string, error)
}

func CurseSpecification(t *testing.T, meany MeanGreeter) {
	got, err := meany.Curse("Anuj")
	assert.NoError(t, err)
	assert.Equal(t, got, "Go to hell, Anuj!")
}
