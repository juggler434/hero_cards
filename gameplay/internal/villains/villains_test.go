package villains_test

import (
	"testing"

	"github.com/juggler434/hero_cards/gameplay/internal/villains"
)

const (
	TestAttack          = 2
	TestScheme          = 1
	TestHitPoints       = 14
	TestDamage          = 5
	TestOverkillDamage  = 100
	TestHealing         = 3
	TestOverkillHealing = 100
)

func createTestVillain() *villains.VillainStats {
	p := villains.VillainParams{
		MaxHitPoints: TestHitPoints,
		Attack:       TestAttack,
		Scheme:       TestScheme,
	}

	return villains.NewVillain(&p)
}

func TestVillainStats_HitPointGettersSetters(t *testing.T) {
	v := createTestVillain()

	mhp := v.GetMaxHitPoints()
	if mhp != TestHitPoints {
		t.Errorf("GetMaxHitPoints should return a villains max hp.  Expected: %d, got %d", TestHitPoints, mhp)
	}

	chp := v.GetCurrentHitPoints()
	if chp != TestHitPoints {
		t.Errorf("GetCurrentHitPoints should return the villains current hit points.  Expected %d, got %d", TestHitPoints, chp)
	}

	v.ApplyDamage(TestDamage)
	chp = v.GetCurrentHitPoints()
	if chp != mhp-TestDamage {
		t.Errorf("ApplyDamage should reduce the villains current hit points by the damage amount.  Expected: %d, got: %d", mhp-TestDamage, chp)
	}

	mhp = v.GetMaxHitPoints()
	if mhp != TestHitPoints {
		t.Errorf("ApplyDamage should not change max hit points.  Expected: %d, got %d", TestHitPoints, mhp)
	}

	v.ApplyDamage(TestOverkillDamage)
	chp = v.GetCurrentHitPoints()
	if chp != 0 {
		t.Errorf("ApplyDamage with damage over current hit points should set currenHitPoints to 0.  Expected: 0, got: %d", chp)
	}

	v.ApplyHealing(TestHealing)
	ehp := chp + TestHealing
	chp = v.GetCurrentHitPoints()
	if chp != ehp {
		t.Errorf("ApplyHealing should raise currentHitPoints by the healing amount.  Expected: %d, got: %d", ehp, chp)
	}

	v.ApplyHealing(TestOverkillHealing)
	chp = v.GetCurrentHitPoints()
	if chp != v.GetMaxHitPoints() {
		t.Errorf("ApplyHealing should not raise a villains currentHealth above their maxHealth.  Expected: %d, got: %d", v.GetMaxHitPoints(), chp)
	}
}

func TestVillainStats_GetAttack(t *testing.T) {
	v := createTestVillain()
	a := v.GetAttack()
	if a != TestAttack {
		t.Errorf("GetAttack should return a villains attack.  Expected: %d, got: %d", TestAttack, a)
	}
}

func TestVillainStats_GetScheme(t *testing.T) {
	v := createTestVillain()
	s := v.GetScheme()
	if s != TestScheme {
		t.Errorf("GetScheme should return a villains scheme.  Expected: %d, got: %d", TestScheme, s)
	}
}

func TestVillainStats_StunGetterSetter(t *testing.T) {
	v := createTestVillain()
	if v.IsStunned() {
		t.Error("IsStunned should return if the villain is stunned.  Expected: false, got: true")
	}

	v.ApplyStun()
	if !v.IsStunned() {
		t.Errorf("ApplyStun should set isStunned to true. Expected: true, got: false")
	}

	v.ClearStun()
	if v.IsStunned() {
		t.Errorf("ClearStun should set isStunned to false.  Expected: false, got: true")
	}
}

func TestVillainStats_ConfusedGetterSetter(t *testing.T) {
	v := createTestVillain()
	if v.IsConfused() {
		t.Error("IsStunned should return if the villain is stunned.  Expected: false, got: true")
	}

	v.ApplyConfused()
	if !v.IsConfused() {
		t.Errorf("ApplyStun should set isStunned to true. Expected: true, got: false")
	}

	v.ClearConfused()
	if v.IsConfused() {
		t.Errorf("ClearStun should set isStunned to false.  Expected: false, got: true")
	}
}

func TestVillainStats_ToughGetterSetter(t *testing.T) {
	v := createTestVillain()
	if v.IsTough() {
		t.Errorf("IsTough should return a villains tough status. Expected: false, got: true")
	}

	v.ApplyTough()
	if !v.IsTough() {
		t.Errorf("ApplyTough should set tough to true.  Expected true, got: false")
	}

	v.ClearTough()
	if v.IsTough() {
		t.Errorf("ClearTough should set tough to false.  Expected: false, got: true")
	}
}
