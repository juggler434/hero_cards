package identity

import (
	"github.com/juggler434/marvelServer/internal/alts"
	"github.com/juggler434/marvelServer/internal/heroes"
	helpers "github.com/juggler434/marvelServer/pkg"
)

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

type IdentityParams struct {
	MaxHitPoints int
	Hero         heroes.Hero
	AlterEgo     alts.AlterEgo
}

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

func (b *BaseIdentity) IsExhausted() bool {
	return b.exhausted
}

func (b *BaseIdentity) Exhaust() bool {
	if b.IsExhausted() == true {
		return false
	}

	b.exhausted = true
	return true
}

func (b *BaseIdentity) Ready() {
	b.exhausted = false
}

func (b *BaseIdentity) IsStunned() bool {
	return b.stunned
}

func (b *BaseIdentity) ApplyStun() {
	b.stunned = true
}

func (b *BaseIdentity) ClearStun() {
	b.stunned = false
}

func (b *BaseIdentity) IsConfused() bool {
	return b.confused
}

func (b *BaseIdentity) ApplyConfused() {
	b.confused = true
}

func (b *BaseIdentity) ClearConfused() {
	b.confused = false
}

func (b *BaseIdentity) IsTough() bool {
	return b.tough
}

func (b *BaseIdentity) ApplyTough() {
	b.tough = true
}

func (b *BaseIdentity) ClearTough() {
	b.tough = false
}

func (b *BaseIdentity) IsHeroForm() bool {
	return b.isHero
}

func (b *BaseIdentity) ChangeForm() {
	b.isHero = !b.isHero
}

func (b *BaseIdentity) ApplyDamage(amount int) {
	b.currentHitPoints = helpers.Max(0, b.currentHitPoints-amount)
}

func (b *BaseIdentity) ApplyHealing(amount int) {
	b.currentHitPoints = helpers.Min(b.maxHitPoints, b.currentHitPoints+amount)
}

func (b *BaseIdentity) GetHandSize() int {
	if b.isHero {
		return b.hero.GetHandSize()
	}

	return b.alterEgo.GetHandSize()
}

func (b *BaseIdentity) GetMaxHitPoints() int {
	return b.maxHitPoints
}

func (b *BaseIdentity) GetCurrentHitPoints() int {
	return b.currentHitPoints
}

func (b *BaseIdentity) GetAttack() int {
	if b.isHero {
		return b.hero.GetAttack()
	}
	return 0
}

func (b *BaseIdentity) GetThwart() int {
	if b.isHero {
		return b.hero.GetThwart()
	}

	return 0
}

func (b *BaseIdentity) GetDefense() int {
	if b.isHero {
		return b.hero.GetDefense()
	}
	return 0
}

func (b *BaseIdentity) GetRecover() int {
	if b.isHero {
		return 0
	}

	return b.alterEgo.GetRecover()
}
