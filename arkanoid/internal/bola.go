package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bola struct {
	position rl.Vector2
	speed    rl.Vector2
	radius   float32
	active   bool
}
