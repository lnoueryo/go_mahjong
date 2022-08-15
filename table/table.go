package table

import (
	"fmt"
	"mahjong/table/wall"
	"strings"
)

type Table struct {
	Game string
	Round int
	wall.Wall
}

func NewTable(game string) *Table {
	round := 1
	w := wall.NewWall()
	return &Table{game, round, w}
}

func (t *Table) Print() {
	fmt.Printf("%s 対局 %s\n", strings.Repeat("=", 40), strings.Repeat("=", 40))
	fmt.Printf("%s\n", t.Game)
	fmt.Printf("%d局目\n", t.Round)
	t.Wall.Print()
}