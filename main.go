package main

import (
	// "fmt"
	"mahjong/player"
	"mahjong/table"
)

func main() {
	table := table.NewTable("東風戦")
	player1 := &player.Player{1, "lnoueryo"}
	player2 := &player.Player{2, "lnoueryo"}
	player3 := &player.Player{3, "lnoueryo"}
	player4 := &player.Player{4, "lnoueryo"}
	table.AddPlayer(player1)
	table.AddPlayer(player2)
	table.AddPlayer(player3)
	table.AddPlayer(player4)
	table.Print()
}