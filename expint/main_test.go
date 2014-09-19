package expint

import (
	"math"
	"testing"

	"github.com/go-math/support/assert"
)

func TestLoad(t *testing.T) {
	s, _ := Load(findFixture("002.json"))

	assert.Equal(s.Cores, uint32(2), t)
	assert.Equal(s.Nodes, uint32(4*2+12), t)

	assert.AlmostEqual(s.system.D, fixtureD, t)

	assert.AlmostEqual(abs(s.system.U), abs(fixtureU), t)
	assert.AlmostEqual(s.system.Λ, fixtureΛ, t)

	assert.AlmostEqual(s.system.E, fixtureE, t)
	assert.AlmostEqual(s.system.F, fixtureF, t)
}

func abs(A []float64) []float64 {
	B := make([]float64, len(A))

	for i := range B {
		B[i] = math.Abs(A[i])
	}

	return B
}
