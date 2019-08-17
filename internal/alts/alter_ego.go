package alts

type AlterEgo interface {
	GetRecover() int
	GetHandSize() int
}

type AlterEgoStats struct {
	recover  int
	handSize int
}

type AlterEgoParams struct {
	Recover  int
	HandSize int
}

func NewAlterEgoStats(p *AlterEgoParams) *AlterEgoStats {
	return &AlterEgoStats{
		p.Recover,
		p.HandSize,
	}
}

func (a *AlterEgoStats) GetRecover() int {
	return a.recover
}

func (a *AlterEgoStats) GetHandSize() int {
	return a.handSize
}
