package main

import (
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
			if x == int32(g.Squares[i].Position.X) ||
			y == int32(g.Squares[i].Position.Y) {

				g.GenerateNewSquare()
			} else {

				g.NewSquare(x, y)
				return
			}
		}
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

		for i := 0; i < len(g.Squares); i++ {
			g.Squares[i].Draw()
		}

		for x := int32(0); x < int32(width); x += tileWidth {
			rl.DrawLine(x, 0, x, height, rl.White)
		}

		for y := int32(0); y < int32(height); y += tileHeight {
			rl.DrawLine(0, y, width, y, rl.White)
		}

		if rl.IsKeyPressed(rl.KeyUp) {
			newSquarePosX := randRange(0, 4) * tileWidth
			newSquarePosY := randRange(0, 4) * tileHeight

			for i := 0; i < len(g.Squares); i++ {
				maxValue := 0
				// TODO: checar a colisao de todos
				g.Squares[i].Position.Y = float32(maxValue)
			}

			// TODO: provavelmente ainda terao coordenadas duplicadas
			for i := 0; i < len(g.Squares); i++ {
				if(newSquarePosY == int32(g.Squares[i].Position.Y)) {
					newSquarePosY = randRange(0, 4) * tileWidth
				}
			}
			
			g.NewSquare(int32(newSquarePosX), int32(newSquarePosY))
		}

		if rl.IsKeyPressed(rl.KeyDown) {
			for i := 0; i < len(g.Squares); i++ {
				maxValue := tileHeight * 3
				g.Squares[i].Position.Y = float32(maxValue)
			}
			
			g.GenerateNewSquare()
		}

		if rl.IsKeyPressed(rl.KeyRight) {
			newSquarePosX := randRange(0, 4) * tileWidth
			newSquarePosY := randRange(0, 4) * tileHeight

			for i := 0; i < len(g.Squares); i++ {
				maxValue := tileWidth * 3
				// TODO: checar a colisao de todos
				g.Squares[i].Position.X = float32(maxValue)
			}

			// TODO: provavelmente ainda terao coordenadas duplicadas
			for i := 0; i < len(g.Squares); i++ {
				if(newSquarePosY == int32(g.Squares[i].Position.Y)) {
					newSquarePosY = randRange(0, 4) * tileWidth
				}
			}
			
			g.NewSquare(int32(newSquarePosX), int32(newSquarePosY))
		}

		if rl.IsKeyPressed(rl.KeyLeft) {
			newSquarePosX := randRange(0, 4) * tileWidth
			newSquarePosY := randRange(0, 4) * tileHeight

			for i := 0; i < len(g.Squares); i++ {
				maxValue := 0
				// TODO: checar a colisao de todos
				g.Squares[i].Position.X = float32(maxValue)
			}

			// TODO: provavelmente ainda terao coordenadas duplicadas
			for i := 0; i < len(g.Squares); i++ {
				if(newSquarePosY == int32(g.Squares[i].Position.Y)) {
					newSquarePosY = randRange(0, 4) * tileWidth
				}
			}
			
			g.NewSquare(int32(newSquarePosX), int32(newSquarePosY))
		}
		rl.ClearBackground(background)
		rl.EndDrawing()
	}

}

func randRange(min, max int) int32 {
	return int32(rand.Intn(max-min) + min)
}
