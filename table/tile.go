package table

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

type Wall struct {
	WallTiles []Tile
	DeadTiles []Tile
}

func NewWall() Wall {
	id := 1
	tiles := make([]Tile, 0, TOTALTILESNUM)
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
					tiles = append(tiles, tile)
					id += 1
				}
			}
		}
	}
	interfaceTiles := ShuffleSlice(tiles)
	shuffledTiles := interfaceTiles.([]Tile)
	wallTiles := shuffledTiles[:WALLTILESNUM-1]
	deadTiles := shuffledTiles[WALLTILESNUM:]
	return Wall{wallTiles, deadTiles}
}

func (w *Wall) Print() {
	fmt.Printf("%s ツモ山 %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	for _, wt := range w.WallTiles {
		wt.Print()
	}
	fmt.Printf("\n%s 壁牌 %s\n", strings.Repeat("=", 25), strings.Repeat("=", 25))
	for _, dt := range w.DeadTiles {
		dt.Print()
	}
}

func ShuffleSlice(slice interface{}) interface{} {
	rand.Seed(time.Now().UnixNano())
    if v, ok := slice.([]Tile); ok {
        rand.Shuffle(len(v), func(i, j int) { v[i], v[j] = v[j], v[i] })
        return v
    }
	return slice
}