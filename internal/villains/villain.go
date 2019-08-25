package villains

import helpers "github.com/juggler434/marvelServer/pkg"

type Villain interface {
	IsStunned() bool
	ApplyStun()
	ClearStun()
	IsTough() bool
	ApplyTough()
	ClearTough()
	IsConfused()
	ApplyConfused()
	ClearConfused()
	ApplyDamage(amount int)
	ApplyHealing(amount int)
	GetMaxHitPoints() int
	GetCurrentHitPoints() int
	GetAttack() int
	GetScheme() int
}

type VillainStats struct {
	maxHitPoints     int
	currentHitPoints int
	stunned          bool
	tough            bool
	confused         bool
	attack           int
	scheme           int
}

type VillainParams struct {
	MaxHitPoints int
	Attack       int
	Scheme       int
}

func NewVillain(p *VillainParams) *VillainStats {
	return &VillainStats{
		maxHitPoints:     p.MaxHitPoints,
		currentHitPoints: p.MaxHitPoints,
		stunned:          false,
		tough:            false,
		confused:         false,
		attack:           p.Attack,
		scheme:           p.Scheme,
	}
}

func (v *VillainStats) IsStunned() bool {
	return v.stunned
}

func (v *VillainStats) ApplyStun() {
	v.stunned = true
}

func (v *VillainStats) ClearStun() {
	v.stunned = false
}

func (v *VillainStats) IsConfused() bool {
	return v.confused
}

func (v *VillainStats) ApplyConfused() {
	v.confused = true
}

func (v *VillainStats) ClearConfused() {
	v.confused = false
}

func (v *VillainStats) IsTough() bool {
	return v.tough
}

func (v *VillainStats) ApplyTough() {
	v.tough = true
}

func (v *VillainStats) ClearTough() {
	v.tough = false
}

func (v *VillainStats) ApplyDamage(amount int) {
	v.currentHitPoints = helpers.Max(v.currentHitPoints-amount, 0)
}

func (v *VillainStats) ApplyHealing(amount int) {
	v.currentHitPoints = helpers.Min(v.currentHitPoints+amount, v.maxHitPoints)
}

func (v *VillainStats) GetMaxHitPoints() int {
	return v.maxHitPoints
}

func (v *VillainStats) GetCurrentHitPoints() int {
	return v.currentHitPoints
}

func (v *VillainStats) GetAttack() int {
	return v.attack
}

func (v *VillainStats) GetScheme() int {
	return v.scheme
}
