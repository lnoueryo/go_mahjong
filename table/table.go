package table

import (
	"fmt"
	"mahjong/table/tiles"
	"mahjong/player"
	"strings"
)

type Table struct {
	Game string
	Round int
	*tiles.Tiles
	Players []*player.Player
}

func NewTable(game string) *Table {
	round := 1
	t := tiles.NewTiles()
	return &Table{game, round, t, []*player.Player{}}
}

func (tb *Table) AddPlayer(player *player.Player) {
	tb.Players = append(tb.Players, player)
	tb.Tiles.SetPlayer(player)
	if len(tb.Players) == 4 {
		tb.Tiles.CreateWallTiles(tb.Players)
	}
}

func (t *Table) Print() {
	fmt.Printf("%s 対局 %s\n", strings.Repeat("=", 40), strings.Repeat("=", 40))
	fmt.Printf("%s\n", t.Game)
	fmt.Printf("%d局目\n", t.Round)
	t.Tiles.Print()
}