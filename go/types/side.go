package types

type Side int

const (
	None Side = iota
	Radiant
	Dire
)

func (mr Side) String() string {
	switch mr {
	case Radiant:
		return "Radiant"
	case Dire:
		return "Dire"
	default:
		return "None"
	}
}
