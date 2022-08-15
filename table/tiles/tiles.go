package tiles

import (
	"fmt"
	"math/rand"
	"strings"
	"time"
)

const (
	TOTALTILESNUM = 136
	WALLTILESNUM = 122
	DEADTILESNUM = 14
)

var (
	VARIETIES = []string{"萬子", "筒子", "索子", "字牌"}
	SUITS     = []string{"1", "2", "3", "4", "5", "6", "7", "8", "9"}
	HONORS    = []string{"東", "南", "西", "北", "白", "發", "中"}
)


type Tiles struct {
	LeftTile		int
	WallTiles		[]*Tile
	DeadTiles		[]*Tile
	HandTiles		map[int][]*Tile
	DiscardedTiles	map[int][]*Tile
}

func NewTiles() *Tiles {
	id := 1
	tiles := make([]*Tile, 0, TOTALTILESNUM)
	for _, v := range VARIETIES {
		if v == "字牌" {
			for _, h := range HONORS {
				for k := 0; k < 4; k++ {
					tile := NewTile(id, h, v)
					tiles = append(tiles, tile)
					id += 1
				}
			}
		} else {
			for _, s := range SUITS {
				for k := 0; k < 4; k++ {
					tile := NewTile(id, s, v)
					if k == 0 && s == "5" {
						tile.Bonus = 1
					}
					tiles = append(tiles, tile)
					id += 1
				}
			}
		}
	}
	interfaceTiles := ShuffleSlice(tiles)
	shuffledTiles := interfaceTiles.([]*Tile)
	wallTiles := shuffledTiles[:WALLTILESNUM-1]
	deadTiles := shuffledTiles[WALLTILESNUM:]
	return &Tiles{len(wallTiles), wallTiles, deadTiles}
}

func (ts *Tiles) DealTiles() []*Tile {
	handTiles := ts.WallTiles[:13]
	ts.WallTiles = ts.WallTiles[13:]
	return handTiles
}

func (ts *Tiles) DrawTile() *Tile {
	ts.LeftTile -= 1
	drawedTile := ts.WallTiles[0]
	ts.WallTiles = ts.WallTiles[1:]
	return drawedTile
}

func (ts *Tiles) Print() {
	fmt.Printf("%s ツモ山 %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	for _, wt := range ts.WallTiles {
		wt.Print()
	}
	fmt.Printf("\n%s 壁牌 %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	for _, dt := range ts.DeadTiles {
		dt.Print()
	}
}

func ShuffleSlice(slice interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
    if v, ok := slice.([]*Tile); ok {
        rand.Shuffle(len(v), func(i, j int) { v[i], v[j] = v[j], v[i] })
        return v
    }
	return slice
}