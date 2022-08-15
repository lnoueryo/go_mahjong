package tiles

import "fmt"

type Tile struct {
	ID      int
	Name    string
	Variety	string
	Bonus	int
}

func NewTile(id int, name string, variety string) *Tile {
	t := new(Tile)
	t.ID = id
	t.Name = name
	t.Variety = variety
	t.Bonus = 0
	return t
}

func (t *Tile) Print() {
	fmt.Printf("ID: %d 名前: %s 種類: %s ", t.ID, t.Name, t.Variety)
}