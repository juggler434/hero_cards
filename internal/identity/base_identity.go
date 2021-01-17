package identity

import (
	"github.com/juggler434/marvelServer/internal/alts"
	"github.com/juggler434/marvelServer/internal/heroes"
	helpers "github.com/juggler434/marvelServer/pkg"
)

//Identity represents a hero/alter ego combo
type Identity interface {
	IsExhausted() bool
	Exhaust() bool
	Ready()
	IsStunned() bool
	ApplyStun()
	ClearStun()
	IsConfused() bool
	ApplyConfused()
	ClearConfused()
	IsTough() bool
	ApplyTough()
	ClearTough()
	IsHeroForm() bool
	ChangeForm()
	ApplyDamage(amount int)
	ApplyHealing(amount int)
	GetHandSize() int
	GetMaxHitPoints() int
	GetCurrentHitPoints() int
	GetAttack() int
	GetThwart() int
	GetDefense() int
	GetRecover() int
}

//BaseIdentity is a generic identity
type BaseIdentity struct {
	maxHitPoints     int
	currentHitPoints int
	hero             heroes.Hero
	alterEgo         alts.AlterEgo
	exhausted        bool
	stunned          bool
	confused         bool
	tough            bool
	isHero           bool
}

//IdentityParams is the parameters needed to initialize a new Identity
type IdentityParams struct {
	MaxHitPoints int
	Hero         heroes.Hero
	AlterEgo     alts.AlterEgo
}

//NewBaseIdentity is the initializer for identity
func NewBaseIdenetity(p *IdentityParams) *BaseIdentity {
	return &BaseIdentity{
		maxHitPoints:     p.MaxHitPoints,
		currentHitPoints: p.MaxHitPoints,
		hero:             p.Hero,
		alterEgo:         p.AlterEgo,
		exhausted:        false,
		stunned:          false,
		confused:         false,
		tough:            false,
		isHero:           false,
	}
}

//IsExhausted gets an identities exhausted status
func (b *BaseIdentity) IsExhausted() bool {
	return b.exhausted
}

//Exhaust sets an identities exhausted status to true, returns false if hero is already exhausted
func (b *BaseIdentity) Exhaust() bool {
	if b.IsExhausted() == true {
		return false
	}

	b.exhausted = true
	return true
}

//Ready sets an identities exhausted status to false
func (b *BaseIdentity) Ready() {
	b.exhausted = false
}

//IsStunned returns true if an identity is stunned
func (b *BaseIdentity) IsStunned() bool {
	return b.stunned
}

//ApplyStun sets an identities stunned attribute to true
func (b *BaseIdentity) ApplyStun() {
	b.stunned = true
}

//ClearStun sets and identities stunned attribute to false
func (b *BaseIdentity) ClearStun() {
	b.stunned = false
}

//IsConfused returns true if an identity is confused
func (b *BaseIdentity) IsConfused() bool {
	return b.confused
}

//ApplyConfused sets an identities confused stat to true
func (b *BaseIdentity) ApplyConfused() {
	b.confused = true
}

//ClearConfused sets an identities confused stat to false
func (b *BaseIdentity) ClearConfused() {
	b.confused = false
}

//IsTough returns an identities tough status
func (b *BaseIdentity) IsTough() bool {
	return b.tough
}

//ApplyTough sets an identities tough status to true
func (b *BaseIdentity) ApplyTough() {
	b.tough = true
}

//ClearTough sets an identities tough status to false
func (b *BaseIdentity) ClearTough() {
	b.tough = false
}

//IsHeroForm return true if an identity is currenty on it's hero side
func (b *BaseIdentity) IsHeroForm() bool {
	return b.isHero
}

//ChangeForm switches isHero between true and false
func (b *BaseIdentity) ChangeForm() {
	b.isHero = !b.isHero
}

//ApplyDamage subtracts the amount from the identity's current hit points to a minimum of 0
func (b *BaseIdentity) ApplyDamage(amount int) {
	b.currentHitPoints = helpers.Max(0, b.currentHitPoints-amount)
}

//ApplyHealing raises an identity's current hit point by amount to a maximum of maxHitPoints
func (b *BaseIdentity) ApplyHealing(amount int) {
	b.currentHitPoints = helpers.Min(b.maxHitPoints, b.currentHitPoints+amount)
}

//GetHandSize returns the hand size for the identity's current form
func (b *BaseIdentity) GetHandSize() int {
	if b.isHero {
		return b.hero.GetHandSize()
	}

	return b.alterEgo.GetHandSize()
}

//GetMaxHitPoints returns the identity's maxHitPoints
func (b *BaseIdentity) GetMaxHitPoints() int {
	return b.maxHitPoints
}

//GetCurrentHitPoints returns an identity's currentHitPoints
func (b *BaseIdentity) GetCurrentHitPoints() int {
	return b.currentHitPoints
}

//GetAttack returns the identity's attack if in hero mode, or 0 if in alter ego
func (b *BaseIdentity) GetAttack() int {
	if b.isHero {
		return b.hero.GetAttack()
	}
	return 0
}

//GetThwart return the identity's thwart value if in hero mode, or 0 if in alter ego
func (b *BaseIdentity) GetThwart() int {
	if b.isHero {
		return b.hero.GetThwart()
	}

	return 0
}

//GetDefense returns the identity's defense value, or 0 if in alter ego
func (b *BaseIdentity) GetDefense() int {
	if b.isHero {
		return b.hero.GetDefense()
	}
	return 0
}

//GetRecover returns an identity's recover stat, or 0 if in hero mode
func (b *BaseIdentity) GetRecover() int {
	if b.isHero {
		return 0
	}

	return b.alterEgo.GetRecover()
}
