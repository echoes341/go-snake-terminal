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
	food := foods[rand.Intn(4)]
	food.X = x
	food.Y = y
	return &food
}
