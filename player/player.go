package player

type Player struct {
	Name string
	Type string
}

func NewPlayer(name, pType string) *Player {
	return &Player{
		Name: name,
		Type: pType,
	}
}