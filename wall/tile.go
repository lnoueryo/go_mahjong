package wall

import "fmt"

type Tile struct {
	ID      int
	Name    string
	Variety string
}

func NewTile(id int, name string, variety string) Tile {
	return Tile{id, name, variety}
}

func (t Tile) Print() {
	fmt.Printf("ID: %d 名前: %s 種類: %s ", t.ID, t.Name, t.Variety)
}