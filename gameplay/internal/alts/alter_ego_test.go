package alts_test

import (
	"testing"

	"github.com/juggler434/marvelServer/gameplay/internal/alts"
)

const (
	TestRecover  = 4
	TestHandSize = 6
)

func createTestAlt() *alts.AlterEgoStats {
	p := &alts.AlterEgoParams{
		Recover:  TestRecover,
		HandSize: TestHandSize,
	}

	return alts.NewAlterEgoStats(p)
}

func TestAlterEgoStats_GetRecover(t *testing.T) {
	ae := createTestAlt()
	rv := ae.GetRecover()
	if rv != TestRecover {
		t.Errorf("GetRecover should return an alter ego's recover stat, expected: %d, got: %d", TestRecover, rv)
	}
}

func TestAlterEgoStats_GetHandSize(t *testing.T) {
	ae := createTestAlt()
	hs := ae.GetHandSize()
	if hs != TestHandSize {
		t.Errorf("GetHandSize should return and alter ego's hand size, expected: %d, got: %d", TestHandSize, hs)
	}
}
