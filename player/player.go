package player



type Player struct {
	ID int
	Name string
}

func NewPlayer(id int, name string) *Player {
	p := new(Player)
	p.ID = id
	p.Name = name
	return p
}