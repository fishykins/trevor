package models

type AlignmentType int

const (
	ChaoticEvilType AlignmentType = iota
	ChaoticNeutralType
	ChaoticGoodType
	NeutralEvilType
	NeutralType
	NeutralGoodType
	LawfulEvilType
	LawfulNeutralType
	LawfulGoodType
)

type Alignment struct {
	Authority float32 `json:"authority" bson:"authority"`
	Morality  float32 `json:"morality" bson:"morality"`
}

func (a *Alignment) IsChaotic() bool {
	return a.Authority < 0
}

func (a *Alignment) IsLawful() bool {
	return a.Authority > 0
}

func (a *Alignment) IsGood() bool {
	return a.Morality > 0
}

func (a *Alignment) IsEvil() bool {
	return a.Morality < 0
}

func (a *Alignment) AlignmentType() AlignmentType {
	if a.IsChaotic() && a.IsEvil() {
		return ChaoticEvilType
	}
	if a.IsChaotic() && a.Morality == 0 {
		return ChaoticNeutralType
	}
	if a.IsChaotic() && a.IsGood() {
		return ChaoticGoodType
	}
	if a.IsLawful() && a.IsEvil() {
		return LawfulEvilType
	}
	if a.IsLawful() && a.IsGood() {
		return LawfulGoodType
	}
	if a.IsLawful() && a.Morality == 0 {
		return LawfulNeutralType
	}
	if a.IsGood() {
		return NeutralGoodType
	}
	if a.IsEvil() {
		return NeutralEvilType
	}
	return NeutralType
}

func ChaoticEvil() Alignment {
	return Alignment{
		Authority: -1,
		Morality:  -1,
	}
}

func ChaoticNeutral() Alignment {
	return Alignment{
		Authority: -1,
		Morality:  0,
	}
}

func ChaoticGood() Alignment {
	return Alignment{
		Authority: -1,
		Morality:  1,
	}
}

func NeutralEvil() Alignment {
	return Alignment{
		Authority: 0,
		Morality:  -1,
	}
}

func Neutral() Alignment {
	return Alignment{
		Authority: 0,
		Morality:  -1,
	}
}

func NeutralGood() Alignment {
	return Alignment{
		Authority: 0,
		Morality:  1,
	}
}

func LawfulEvil() Alignment {
	return Alignment{
		Authority: 1,
		Morality:  -1,
	}
}

func LawfulNeutral() Alignment {
	return Alignment{
		Authority: 1,
		Morality:  0,
	}
}

func LawfulGood() Alignment {
	return Alignment{
		Authority: 1,
		Morality:  1,
	}
}
