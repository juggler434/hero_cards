package identity_test

import (
	"testing"

	"github.com/juggler434/marvelServer/internal/alts"
	"github.com/juggler434/marvelServer/internal/heroes"
	"github.com/juggler434/marvelServer/internal/identity"
)

const (
	TestAttack          = 2
	TestDefense         = 3
	TestThwart          = 4
	TestRecover         = 5
	TestMaxHitPoints    = 10
	TestHeroHandSize    = 4
	TestAltEgoHandSize  = 6
	TestDamage          = 4
	TestOverkillDamage  = 1000
	TestHealing         = 5
	TestOverkillHealing = 2000
)

func getTestHero() *heroes.HeroStats {
	p := heroes.HeroParams{
		Attack:   TestAttack,
		Thwart:   TestThwart,
		Defense:  TestDefense,
		HandSize: TestHeroHandSize,
	}
	return heroes.NewHeroStats(&p)
}

func getTestAlterEgo() *alts.AlterEgoStats {
	p := alts.AlterEgoParams{
		Recover:  TestRecover,
		HandSize: TestAltEgoHandSize,
	}

	return alts.NewAlterEgoStats(&p)
}

func getTestBaseIdentity() *identity.BaseIdentity {
	var testHero heroes.Hero
	testHero = getTestHero()

	var testAlterEgo alts.AlterEgo
	testAlterEgo = getTestAlterEgo()

	p := identity.NewBaseIdentityParams{
		MaxHitPoints: TestMaxHitPoints,
		Hero:         testHero,
		AlterEgo:     testAlterEgo,
	}

	return identity.NewBaseIdentity(&p)
}

func TestBaseIdentity_ConfusedGetterSetters(t *testing.T) {
	i := getTestBaseIdentity()
	if i.IsConfused() {
		t.Error("IsConfused should have returned false, instead returned true")
	}
	i.ApplyConfused()
	if !i.IsConfused() {
		t.Error("ApplyConfused should have set confused to true, but it was false")
	}
	i.ClearConfused()
	if i.IsConfused() {
		t.Error("ClearConfused should have set confused to false, but it was true")
	}
}

func TestBaseIdentity_StunnedGettersSetters(t *testing.T) {
	i := getTestBaseIdentity()
	if i.IsStunned() {
		t.Errorf("IsStunned should have return false, but it returned true")
	}

	i.ApplyStun()
	if !i.IsStunned() {
		t.Error("ApplyStun should have set stun to true, but it was false")
	}

	i.ClearStun()
	if i.IsStunned() {
		t.Error("ClearStun should have set stun to false, but it was true")
	}
}

func TestBaseIdentity_ToughGettersSetters(t *testing.T) {
	i := getTestBaseIdentity()
	if i.IsTough() {
		t.Error("IsTough should have return false, but it returned true")
	}

	i.ApplyTough()
	if !i.IsTough() {
		t.Error("Apply tough should have set tough to true, but it was false")
	}

	i.ClearTough()
	if i.IsTough() {
		t.Error("Clear tough should have set tough to false, but it was true")
	}
}

func TestBaseIdentity_ExhaustGettersSetters(t *testing.T) {
	i := getTestBaseIdentity()
	if i.IsExhausted() {
		t.Error("IsExhausted should have returned false, but retunned true")
	}

	r := i.Exhaust()
	if !i.IsExhausted() {
		t.Error("Exhaust should have set exhausted to true, but it was false")
	}
	if !r {
		t.Error("Calling Exhaust when exhaust was set to false should have returned true, it returned false")
	}

	sr := i.Exhaust()
	if sr {
		t.Error("Calling exhaust when exhaust was set to true should have returned false, it returned true")
	}

	i.Ready()
	if i.IsExhausted() {
		t.Errorf("Ready should have set exhausted to false, it was true")
	}
}

func TestBaseIdentity_HeroFormGettersSetters(t *testing.T) {
	i := getTestBaseIdentity()
	if i.IsHeroForm() {
		t.Errorf("IsHeroForm should have returned false, it returned true")
	}

	i.ChangeForm()
	if !i.IsHeroForm() {
		t.Errorf("Change Form should have set set isHero to true, but it was false")
	}

	i.ChangeForm()
	if i.IsHeroForm() {
		t.Errorf("Change form should have changed isHero to false, but it was true")
	}
}

func TestBaseIdentity_HitPointGetterSetters(t *testing.T) {
	i := getTestBaseIdentity()
	mhp := i.GetMaxHitPoints()
	chp := i.GetCurrentHitPoints()

	if chp != mhp {
		t.Errorf("Current Hit Points should equal max hit points: expected: %d, got: %d", mhp, chp)
	}

	i.ApplyDamage(TestDamage)
	chp = i.GetCurrentHitPoints()
	if chp != mhp-TestDamage {
		t.Errorf("ApplyDamage should have reduced current hip points by %d, expected: %d, got %d", TestDamage, mhp-TestDamage, chp)
	}

	i.ApplyDamage(TestOverkillDamage)
	chp = i.GetCurrentHitPoints()
	if chp != 0 {
		t.Errorf("Applying more damage than a hero's current hit points should set hit points to 0.  Expected: 0, got %d", chp)
	}

	i.ApplyHealing(TestHealing)
	ehp := chp + TestHealing
	chp = i.GetCurrentHitPoints()
	if chp != ehp {
		t.Errorf("ApplyHealing should have raised currentHitPoints by %d.  Expected %d, got: %d", TestHealing, ehp, chp)
	}

	i.ApplyHealing(TestOverkillHealing)
	chp = i.GetCurrentHitPoints()
	if chp != mhp {
		t.Errorf("Applying more healing than the hero's max hit points should set current hit points to max.  Expected: %d, got %d", mhp, chp)
	}
}

func TestBaseIdentity_GetHandSize(t *testing.T) {
	i := getTestBaseIdentity()

	aehs := i.GetHandSize()
	if aehs != TestAltEgoHandSize {
		t.Errorf("While isHero is false, hand size should equal alter ego hand size.  Expected: %d, got %d", TestAltEgoHandSize, aehs)
	}

	i.ChangeForm()
	hhs := i.GetHandSize()
	if hhs != TestHeroHandSize {
		t.Errorf("while isHero is true, hand size should equal hero hand size.  Expceted: %d, got %d", TestHeroHandSize, hhs)
	}
}

func TestBaseIdentity_GetAttack(t *testing.T) {
	i := getTestBaseIdentity()

	a := i.GetAttack()
	if a != 0 {
		t.Errorf("When isHero is false, GetAttack should return 0.  Expected 0, got: %d", a)
	}

	i.ChangeForm()
	a = i.GetAttack()
	if a != TestAttack {
		t.Errorf("When isHero is true, GetAttack should return the hero's attack.  Expected %d, got %d", TestAttack, a)
	}
}

func TestBaseIdentity_GetThwart(t *testing.T) {
	i := getTestBaseIdentity()

	th := i.GetThwart()
	if th != 0 {
		t.Errorf("When isHero is false, GetThwart should return 0.  Expected 0, got: %d", th)
	}

	i.ChangeForm()
	th = i.GetThwart()
	if th != TestThwart {
		t.Errorf("When isHero is true, GetThwart should return the hero's thwart.  Expected %d, got %d", TestThwart, th)
	}
}

func TestBaseIdentity_GetDefense(t *testing.T) {
	i := getTestBaseIdentity()

	d := i.GetDefense()
	if d != 0 {
		t.Errorf("When isHero is false, GetDefense should return 0.  Expected 0, got: %d", d)
	}

	i.ChangeForm()
	d = i.GetDefense()
	if d != TestDefense {
		t.Errorf("When isHero is true, GetDefense should return the hero's defense.  Expected %d, got %d", TestDefense, d)
	}
}

func TestBaseIdentity_GetRecover(t *testing.T) {
	i := getTestBaseIdentity()

	i.ChangeForm()
	r := i.GetRecover()
	if r != 0 {
		t.Errorf("When isHero is true, GetRecover should return 0.  Expected 0, got: %d", r)
	}

	i.ChangeForm()
	r = i.GetRecover()
	if r != TestRecover {
		t.Errorf("When isHero is false, GetRecover should return the alter ego's recover.  Expected %d, got %d", TestRecover, r)
	}
}
