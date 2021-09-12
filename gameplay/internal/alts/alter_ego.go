package alts

//AlterEgo represents the Alter Ego side of an identity
type AlterEgo interface {
	GetRecover() int
	GetHandSize() int
}

//AlterEgoStats are the stats associated with an Alter Ego
type AlterEgoStats struct {
	recover  int
	handSize int
}

//AlterEgoParams are the stats needed to initialize a new Alter Ego
type AlterEgoParams struct {
	Recover  int
	HandSize int
}

//NewAlterEgoStats is the initializer for AlterEgoStats
func NewAlterEgoStats(p *AlterEgoParams) *AlterEgoStats {
	return &AlterEgoStats{
		p.Recover,
		p.HandSize,
	}
}

//GetRecover is the reader for an alter egos recover stat
func (a *AlterEgoStats) GetRecover() int {
	return a.recover
}

//GetHandSize is the reader for an alter egos hand size
func (a *AlterEgoStats) GetHandSize() int {
	return a.handSize
}
