package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

type Bola struct {
	Position   rl.Vector2
	Velocidade rl.Vector2
	Radius     float32
	Ativo      bool
}
