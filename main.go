package main

import (
	"fmt"
	"math/rand"
	"sort"

	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	height     = 800
	width      = 800
	tileHeight = height / 4
	tileWidth  = height / 4
	minimumX   = 0
	minimumY   = 0
	maximumX   = tileWidth
	maximumY   = tileHeight
)

var (
	background = rl.NewColor(41, 46, 66, 1)
)

type Square struct {
	Position rl.Vector2
	Size     rl.Vector2
	Color    rl.Color
	Value    int32
	Id       int32
}

// g.Squares[y * altura + x]
type Game struct {
	Squares []Square
}

func (s *Square) Draw() {
	rl.DrawRectangleV(s.Position, s.Size, s.Color)
	value := fmt.Sprintf("%d (%d)", s.Value, s.Id)
	textPosition := func() (int32, int32) {
		X := s.Position.X + s.Size.X/2
		Y := s.Position.Y + s.Size.Y/2

		return int32(X), int32(Y)
	}
	x, y := textPosition()
	rl.DrawText(value, x, y, 20, rl.White)
}

func (g *Game) NewSquare(x int32, y int32, value int32, id int32) {
	square := Square{
		rl.NewVector2(float32(x), float32(y)),
		rl.NewVector2(float32(tileWidth), float32(tileHeight)),
		rl.Blue,
		value,
		id,
	}
	g.Squares = append(g.Squares, square)
}

func (g *Game) GenerateNewSquare() {
	for {
		x := randRange(0, 4) * tileWidth
		y := randRange(0, 4) * tileHeight

		for i := 0; i < len(g.Squares); i++ {

			if x == int32(g.Squares[i].Position.X) || y == int32(g.Squares[i].Position.Y) {
				x = randRange(0, 4) * tileWidth
				y = randRange(0, 4) * tileHeight
				i = 0
			}
		}
		g.NewSquare(x, y, 2, int32(rand.Intn(5)))
		return
	}
}

func main() {
	rl.InitWindow(width, height, "2048")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	g := Game{}

	g.NewSquare(0, 200, 2, 40)
	g.NewSquare(200, 200, 4, 60)
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
			coordinate := g.MovimentHandleCollision(&g.Squares[i], float32(maxValue), "y", "down")
			g.Squares[i].Position.Y = coordinate
		}
	}

	if rl.IsKeyPressed(rl.KeyUp) {
		for i := 0; i < len(g.Squares); i++ {
			maxValue := 0
			coordinate := g.MovimentHandleCollision(&g.Squares[i], float32(maxValue), "y", "up")
			g.Squares[i].Position.Y = coordinate
		}
	}

	if rl.IsKeyPressed(rl.KeyLeft) {
		for i := 0; i < len(g.Squares); i++ {
			maxValue := 0
			coordinate := g.MovimentHandleCollision(&g.Squares[i], float32(maxValue), "x", "left")
			g.Squares[i].Position.X = coordinate
		}
	}

	if rl.IsKeyPressed(rl.KeyRight) {
		for i := 0; i < len(g.Squares); i++ {
			maxValue := tileWidth * 3
			coordinate := g.MovimentHandleCollision(&g.Squares[i], float32(maxValue), "x", "right")
			g.Squares[i].Position.X = coordinate
		}
	}

	if rl.IsKeyPressed(rl.KeyUp) ||
		rl.IsKeyPressed(rl.KeyDown) ||
		rl.IsKeyPressed(rl.KeyLeft) ||
		rl.IsKeyPressed(rl.KeyRight) {
		// g.GenerateNewSquare()
	}
}

func (g *Game) MovimentHandleCollision(s *Square, coordinateValue float32, coordinateName string, direction string) float32 {
	switch coordinateName {
	case "y":
		if direction == "down" {
			sort.Slice(g.Squares, func(i, j int) bool {
				return g.Squares[j].Position.Y < g.Squares[i].Position.Y
			})
		} else {
			sort.Slice(g.Squares, func(i, j int) bool {
				return g.Squares[j].Position.Y > g.Squares[i].Position.Y
			})
		}

		for i := 0; i < len(g.Squares); i++ {
			// ja tem alguem e esse alguem nao sou eu
			if g.Squares[i].Position.Y == coordinateValue &&
				s.Position.X == g.Squares[i].Position.X {
				// esse alguem tem o mesmo valor que eu
				if s.Value == g.Squares[i].Value {
					s.Value *= 2
					g.Squares = append(g.Squares[:i], g.Squares[i+1:]...)
					return coordinateValue
				} else {
					if direction == "down" {
						coordinateValue -= tileWidth
					} else {
						coordinateValue += tileWidth
					}
					i = 0
				}
			} else {
				return coordinateValue
			}
		}
	case "x":
		fmt.Printf("\nsquare position: %f \ncoordinate value: %f", s.Position.X, coordinateValue)
		if s.Position.X == coordinateValue && (coordinateValue == maximumX || coordinateValue == minimumX) {
			return coordinateValue
		}

		if direction == "right" {
			sort.Slice(g.Squares, func(i, j int) bool {
				return g.Squares[j].Position.X < g.Squares[i].Position.X
			})
		} else {
			sort.Slice(g.Squares, func(i, j int) bool {
				return g.Squares[j].Position.X > g.Squares[i].Position.X
			})
		}

		for i := 0; i < len(g.Squares); i++ {

			// ja tem alguem e esse alguem nao sou eu
			if g.Squares[i].Position.X == coordinateValue && s.Position.Y == g.Squares[i].Position.Y {
				fmt.Println("Tem alguem e nao sou eu")

				// esse alguem tem o mesmo valor que eu
				if s.Value == g.Squares[i].Value {
					fmt.Println("O valor eh o mesmo")
					s.Value *= 2
					g.Squares = append(g.Squares[:i], g.Squares[i+1:]...)

					return coordinateValue
				} else {
					if direction == "right" {
						coordinateValue -= tileWidth
					} else {
						coordinateValue += tileWidth
					}
					i = 0
				}
			} else {
				return coordinateValue
			}
		}
	default:
		return 0
	}
	return 0
}

func randRange(min, max int) int32 {
	return int32(rand.Intn(max-min) + min)
}
