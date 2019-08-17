package heroes_test

import (
	"testing"

	"github.com/juggler434/marvelServer/internal/heroes"
)

const (
	TestAttackValue  = 3
	TestThwartValue  = 2
	TestDefenseValue = 1
	TestHandSize     = 5
)

func createTestHero() *heroes.HeroStats {
	p := heroes.HeroParams{
		Attack:   TestAttackValue,
		Thwart:   TestThwartValue,
		Defense:  TestDefenseValue,
		HandSize: TestHandSize,
	}
	return heroes.NewHeroStats(&p)
}

func TestHeroStats_GetAttack(t *testing.T) {
	h := createTestHero()
	av := h.GetAttack()
	if av != TestAttackValue {
		t.Errorf("GetAttack should return the hero's attack value: expected %d, got %d", TestAttackValue, av)
	}
}

func TestHeroStats_GetThwart(t *testing.T) {
	h := createTestHero()
	tv := h.GetThwart()
	if tv != TestThwartValue {
		t.Errorf("GetThwart should return the hero's thwart value: expected %d, got %d", TestThwartValue, tv)
	}
}

func TestHeroStats_GetDefence(t *testing.T) {
	h := createTestHero()
	dv := h.GetDefense()
	if dv != TestDefenseValue {
		t.Errorf("GetDefence should return a hero's defense value: expected %d, got %d", TestDefenseValue, dv)
	}
}

func TestHeroStats_GetHandSize(t *testing.T) {
	h := createTestHero()
	hs := h.GetHandSize()
	if hs != TestHandSize {
		t.Errorf("GetHandSize should return a hero's hand size: expected %d, got %d", TestHandSize, hs)
	}
}
