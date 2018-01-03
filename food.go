package main

import "math/rand"

type food struct {
	Coord
	points int
	symbol rune
}

//foods available, chosen randomly
var foods = [...]food{
	{Coord{}, 10, '♥'},
	{Coord{}, 10, '♠'},
	{Coord{}, 10, '♣'},
	{Coord{}, 10, '♦'},
}

func newFood(x, y int) *food {
	pos := rand.Intn(4)
	food := foods[pos]
	food.X = x
	food.Y = y
	return &food
}
