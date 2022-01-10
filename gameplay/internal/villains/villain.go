package villains

import helpers "github.com/juggler434/hero_cards/pkg"

//Villain represents a base villain
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

//VillainStats are the base stats to make up a villain
type VillainStats struct {
	maxHitPoints     int
	currentHitPoints int
	stunned          bool
	tough            bool
	confused         bool
	attack           int
	scheme           int
}

//VillainParams are the parameters needed to initialize a new villain
type VillainParams struct {
	MaxHitPoints int
	Attack       int
	Scheme       int
}

//NewVillain initiates a new VillainStats
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

//IsStunned returns a villains stunned status
func (v *VillainStats) IsStunned() bool {
	return v.stunned
}

//ApplyStun sets a villains stunned status to true
func (v *VillainStats) ApplyStun() {
	v.stunned = true
}

//ClearStun sets a villains stunned status to false
func (v *VillainStats) ClearStun() {
	v.stunned = false
}

//IsConfused returns a villains confused status
func (v *VillainStats) IsConfused() bool {
	return v.confused
}

//ApplyConfused sets a villains confused status to true
func (v *VillainStats) ApplyConfused() {
	v.confused = true
}

//ClearConfused sets a villains confused status to false
func (v *VillainStats) ClearConfused() {
	v.confused = false
}

//IsTough returns a villains tough status
func (v *VillainStats) IsTough() bool {
	return v.tough
}

//ApplyTough sets a villains tough status to true
func (v *VillainStats) ApplyTough() {
	v.tough = true
}

//ClearTough sets a villains tough status to false
func (v *VillainStats) ClearTough() {
	v.tough = false
}

//ApplyDamage reduces a villains currentHitPoints by amount, to a minimum of 0
func (v *VillainStats) ApplyDamage(amount int) {
	v.currentHitPoints = helpers.Max(v.currentHitPoints-amount, 0)
}

//ApplyHealing increases a villains currentHitPoints by amount, to a maximim of MaxHitPoints
func (v *VillainStats) ApplyHealing(amount int) {
	v.currentHitPoints = helpers.Min(v.currentHitPoints+amount, v.maxHitPoints)
}

//GetMaxHitPoints returns maxHitPoints
func (v *VillainStats) GetMaxHitPoints() int {
	return v.maxHitPoints
}

//GetCurrentHitPoints returns currentHitPoints
func (v *VillainStats) GetCurrentHitPoints() int {
	return v.currentHitPoints
}

//GetAttack returns attack
func (v *VillainStats) GetAttack() int {
	return v.attack
}

//GetScheme returns scheme
func (v *VillainStats) GetScheme() int {
	return v.scheme
}
