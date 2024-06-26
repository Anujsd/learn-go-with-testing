package interactions_test

import (
	"testing"

	"github.com/anujsd/go-specs-greet/domain/interactions"
	"github.com/anujsd/go-specs-greet/specifications"
)

func TestGreet(t *testing.T) {
	specifications.GreetSpecification(t, specifications.GreetAdapter(interactions.Greet))
}

func TestCurse(t *testing.T) {
	specifications.CurseSpecification(t, specifications.CurseAdapter(interactions.Curse))
}
