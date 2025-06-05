package main

import (
	"fmt"
	"math/rand"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	height     = 800
	width      = 800
	tileHeight = height / 4
	tileWidth  = height / 4
)

var (
	background = rl.NewColor(41, 46, 66, 1)
)

type Square struct {
	Position rl.Vector2
	Size     rl.Vector2
	Color    rl.Color
	Value int32
}


type Game struct {
	Squares []Square
}

func (s *Square) Draw() {
	rl.DrawRectangleV(s.Position, s.Size, s.Color)
	value := fmt.Sprint(s.Value)
	textPosition := func () (int32, int32) {
		X := s.Position.X + s.Size.X / 2
		Y := s.Position.Y + s.Size.Y / 2

		return int32(X), int32(Y)
	}
	x, y := textPosition()
	rl.DrawText(value, x, y, 10, rl.White)
}

func (g *Game) NewSquare(x int32, y int32) {
	square := Square{
		rl.NewVector2(float32(x), float32(y)),
		rl.NewVector2(float32(tileWidth), float32(tileHeight)),
		rl.Blue,
		2,
	}
	g.Squares = append(g.Squares, square)
}

func (g *Game) GenerateNewSquare() {
	for {
		x := randRange(0, 4) * tileWidth
		y := randRange(0 ,4) * tileHeight

		for i := 0; i < len(g.Squares); i++ {

			if x == int32(g.Squares[i].Position.X) || y == int32(g.Squares[i].Position.Y) {
				x = randRange(0, 4) * tileWidth
				y = randRange(0 ,4) * tileHeight
				i = 0
			}
		}
		g.NewSquare(x, y)
		return
	}
}

func main() {
	rl.InitWindow(800, 800, "2048")	
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	g := Game{}
	g.NewSquare(600, 400)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()
		
		g.Update()
		g.Draw()

		rl.ClearBackground(background)
		rl.EndDrawing()
	}

}

func (g *Game) Draw() {
		// Draw squares
		for i := 0; i < len(g.Squares); i++ {
			g.Squares[i].Draw()
		}

		// Draw grid
		for x := int32(0); x < int32(width); x += tileWidth {
			rl.DrawLine(x, 0, x, height, rl.White)
		}

		for y := int32(0); y < int32(height); y += tileHeight {
			rl.DrawLine(0, y, width, y, rl.White)
		}
}

func (g *Game) Update() {
	g.MovimentHandle()
}

func (g *Game) MovimentHandle() {
		if rl.IsKeyPressed(rl.KeyDown) {
			for i := 0; i < len(g.Squares); i++ {
				maxValue := tileHeight * 3
				g.Squares[i].Position.Y = float32(maxValue)
			}
			g.GenerateNewSquare()
		}

		if rl.IsKeyPressed(rl.KeyUp) {
			for i := 0; i < len(g.Squares); i++ {
				maxValue :=  0
				g.Squares[i].Position.Y = float32(maxValue)
			}
			g.GenerateNewSquare()
		}

		if rl.IsKeyPressed(rl.KeyLeft) {
			for i := 0; i < len(g.Squares); i++ {
				maxValue := 0
				g.Squares[i].Position.X = float32(maxValue)
			}
			g.GenerateNewSquare()
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			for i := 0; i < len(g.Squares); i++ {
				maxValue :=  tileWidth * 3
				g.Squares[i].Position.X = float32(maxValue)
			}
			g.GenerateNewSquare()
		}
}

func randRange(min, max int) int32 {
	return int32(rand.Intn(max-min) + min)
}
