package table

import (
	"fmt"
	"mahjong/table/tiles"
	"strings"
)

type Table struct {
	Game string
	Round int
	*tiles.Tiles
}

func NewTable(game string) *Table {
	round := 1
	t := tiles.NewTiles()
	return &Table{game, round, t}
}

func (t *Table) Print() {
	fmt.Printf("%s 対局 %s\n", strings.Repeat("=", 40), strings.Repeat("=", 40))
	fmt.Printf("%s\n", t.Game)
	fmt.Printf("%d局目\n", t.Round)
	t.Tiles.Print()
}