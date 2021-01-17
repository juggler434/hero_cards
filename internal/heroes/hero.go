package heroes

//Hero represents the hero side of an identity
type Hero interface {
	GetAttack() int
	GetThwart() int
	GetDefense() int
	GetHandSize() int
}

//HeroStats represents the stats for a base hero
type HeroStats struct {
	attack   int
	thwart   int
	defense  int
	handSize int
}

//HeroParams are the statistics needed to initialize a new hero
type HeroParams struct {
	Attack   int
	Thwart   int
	Defense  int
	HandSize int
}

//NewHeroStats is the initializer for Hero
func NewHeroStats(p *HeroParams) *HeroStats {
	return &HeroStats{
		p.Attack,
		p.Thwart,
		p.Defense,
		p.HandSize,
	}
}

//GetAttack is the reader for a heroes attack
func (h *HeroStats) GetAttack() int {
	return h.attack
}

//GetDefense is the reader for a heroes defense
func (h *HeroStats) GetDefense() int {
	return h.defense
}

//GetThwart is the reader for a heroes thwart
func (h *HeroStats) GetThwart() int {
	return h.thwart
}

//GetHandSize is the reader for a heroes hand size
func (h *HeroStats) GetHandSize() int {
	return h.handSize
}
