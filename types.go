// Copyright (c) 2024 Michael D Henderson. All rights reserved.

package sb

// Agent is a player or NPC in the game.
// Agents receive "commissions" from the Family for their work.
type Agent struct {
	Name          string
	GoldCoins     int
	VictoryPoints int
	Favors        int
}

// Family is a group of NPCs that sponsor the arts and employ the actors.
type Family struct {
	Name   string
	Crest  string
	Avatar string // an image of the family crest
}

type FAMILY int

const (
	FAMILY_NONE FAMILY = iota
	FAMILY_COOPER
	FAMILY_WALKER
	FAMILY_PAYNE
	FAMILY_HUGHES
	FAMILY_FLETCHER
	FAMILY_NASH
)

type Actor struct {
	Name           string // must be first name followed by last name to make the sort work
	Rank           int    // from 1 through 9
	Benefactor     FAMILY
	Avatar         string
	Page, Row, Col int
}

// Compare returns the tie-fighter result (-1,0,1) of two actors
// by comparing their ranks. Ties are resolved by comparing the names.
// Results:
//
//	-1 --- a has a lower rank than b
//	 0 --- a and b are the same actor
//	 1 --- a has a higher rank than b
func (a Actor) Compare(b Actor) int {
	if a.Rank < b.Rank {
		return -1
	} else if a.Rank == b.Rank {
		if a.Name < b.Name {
			return -1
		} else if a.Name == b.Name {
			return 0
		}
	}
	return 1
}

// Play is a theatrical production. Agents compete to place their actors in a Play.
// Doing so earns them commissions in the form of gold coins, favors, and victory points.
//
// A Play is considered finished when the sum of the ranks of the actors is
// greater than or equal to the theatrical value.
//
// If the play has a Benefactor, the Play is worth an additional 1 Victory Point to the final score.
//
// I believe that the initial PointValue is also the number of favors to assign to the Play.
type Play struct {
	Name            string
	TheatricalValue int
	Category        Category
	PointValue      int
	Benefactor      FAMILY       // the Family that sponsored the Play
	Favors          []FavorToken // randomly assigned to the Play when it is created
	Actors          []Actor      // Actors are assigned to the Play by Agents
	Page, Row, Col  int
}

type Category int

const (
	COMEDY Category = iota
	HISTORY
	TRAGEDY
)

// IsFinished returns true if the play has accrued enough points to be
// considered finished. The play is considered finished when the sum of
// the ranks of the actors is greater than or equal to the theatrical value.
func (p Play) IsFinished() bool {
	accruedValue := 0
	for _, actor := range p.Actors {
		accruedValue += actor.Rank
	}
	return accruedValue >= p.TheatricalValue
}

type FavorToken struct{}
