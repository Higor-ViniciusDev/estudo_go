package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	JOGADOR_MAXIMA_VIDA = 5
	LINHA_POR_BLOCO     = 5
	BLOCOS_POR_LINHA    = 20
)

type Jogador struct {
	position rl.Vector2
	size     rl.Vector2
	life     int
}
