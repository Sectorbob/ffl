package model

type PlayerStats struct {
	Passing        Passing
	Rushing        Rushing
	Receiving      Receiving
	Tackling       Tackling
	PassDefense    PassDefense
	RunDefense     RunDefense
	Fumbling       Fumbling
	KickoffReturns KickoffReturns
	PuntReturns    PuntReturns
}

// Passing is a statline for passing related stats.
type Passing struct {
	// Att is the number of passing attempts thrown.
	Att int

	// Cmpl is the number of passing completions thrown.
	Cmpl int

	// Yards is the number of gross passing yards thrown.
	Yards int

	// TD is the number of passing touchdowns thrown.
	TD int

	// Int is the number of interceptions thrown.
	Int int

	// Lng is the longest completion thrown (in yards).
	Lng int

	// Cmpl20Plus is the number of completions thrown for 20+ yards
	Cmpl20Plus int

	// Cmpl40Plus is the number of completions thrown for 40+ yards
	Cmpl40Plus int

	// Sacks is the number of times the passer was sacked
	Sacks int

	// SackYards is the total number of yards the passer was sacked for.
	SackYards int
}

// CompPct is the completion percentage. If no attempts have been made 0.0 will
// be returned.
func (p *Passing) CompPct() float32 {
	if p.Att == 0 {
		return 0.0
	}
	return float32(p.Cmpl) / float32(p.Att)
}

// YPA is the yards per passing attempt. If no attempts have been made 0.0 will
// be returned.
func (p *Passing) YPA() float32 {
	if p.Att == 0 {
		return 0
	}
	return float32(p.Yards) / float32(p.Att)
}

// TDPct is the percentage of passing touchdowns per passing attempt. If no
// attempts have been made 0.0 will be returned.
func (p *Passing) TDPct() float32 {
	if p.Att == 0 {
		return 0
	}
	return float32(p.TD) / float32(p.Att)
}

// IntPct is the percentage of interceptions thrown per passing attempt. If no
// attempts have been made 0.0 will be returned.
func (p *Passing) IntPct() float32 {
	if p.Att == 0 {
		return 0
	}
	return float32(p.TD) / float32(p.Att)
}

// Rushing is a statline for rushing related stats.
type Rushing struct {
	// Att is the number of rushing attempts.
	Att int

	// Yards is the number of rushing yards.
	Yards int

	// TD is the number of rushing touchdowns.
	TD int

	// Lng is the longest rushing attempt (in yards).
	Lng int

	// Cmpl20Plus is the number of 20+ yard rushing attempts.
	Cmpl20Plus int

	// Cmpl40Plus is the number of 40+ yard rushing attempts.
	Cmpl40Plus int

	// Fumbles is the numnber of times the ball was fumbled while rushing.
	Fumbles int
}

// YPA is the yards per rushing attempt. If no attempts have been made 0.0 will
// be returned.
func (p *Rushing) YPA() float32 {
	if p.Att == 0 {
		return 0
	}
	return float32(p.Yards) / float32(p.Att)
}

// Receiving is a statline for receiving related stats.
type Receiving struct {
	Tgt        int
	Rec        int
	Yards      int
	TD         int
	Lng        int
	Cmpl20Plus int
	Cmpl40Plus int
	Fumbles    int
}

// YPT is the yards per target. If no targets have been made, 0.0 will be
// returned.
func (p *Receiving) YPT() float32 {
	if p.Tgt == 0 {
		return 0
	}
	return float32(p.Yards) / float32(p.Tgt)
}

// YPR is the yards per reception. If no receptions have been made, 0.0 will be
// returned.
func (p *Receiving) YPR() float32 {
	if p.Tgt == 0 {
		return 0
	}
	return float32(p.Yards) / float32(p.Rec)
}

type Tackling struct {
	Solo           int
	Ast            int
	Sacks          float32
	SackYards      int
	Safeties       int
	TacklesForLoss int
}

type PassDefense struct {
	Int      int
	IntTD    int
	IntYards int
	IntLng   int

	// Defended is the number of passes defended.
	Defended int
}

type RunDefense struct {
	Stuffs     int
	StuffYards int
}

type Fumbling struct {
	Fumbles       int
	Lost          int
	Forced        int
	OwnRecovered  int
	OppRecovered  int
	RecoveryYards int
	TD            int
}

type KickoffReturns struct {
	Att         int
	Yards       int
	Lng         int
	TD          int
	Att20Plus   int
	Att40Plus   int
	FairCatches int
	Fumbles     int
}

type PuntReturns struct {
	Att         int
	Yards       int
	Lng         int
	TD          int
	Att20Plus   int
	Att40Plus   int
	FairCatches int
	Fumbles     int
}
