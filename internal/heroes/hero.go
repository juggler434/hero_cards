package heroes

type Hero interface {
	GetAttack() int
	GetThwart() int
	GetDefense() int
	GetHandSize() int
}

type HeroStats struct {
	attack   int
	thwart   int
	defense  int
	handSize int
}

type HeroParams struct {
	Attack   int
	Thwart   int
	Defense  int
	HandSize int
}

func NewHeroStats(p *HeroParams) *HeroStats {
	return &HeroStats{
		p.Attack,
		p.Thwart,
		p.Defense,
		p.HandSize,
	}
}

func (h *HeroStats) GetAttack() int {
	return h.attack
}

func (h *HeroStats) GetDefense() int {
	return h.defense
}

func (h *HeroStats) GetThwart() int {
	return h.thwart
}

func (h *HeroStats) GetHandSize() int {
	return h.handSize
}
