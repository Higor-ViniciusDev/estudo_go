package main

import (
	"estudo_go/arkanoid/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(internal.ScreenWidth, internal.ScreenHeight, "Arkanoid")
	defer rl.CloseWindow()

	rl.SetTargetFPS(60)
	g := internal.Game{}

	internal.TelaLoading()

	internal.InicializaIntesJogo(&g)

	internal.TelaJogoInicial(&g)

	rl.CloseWindow()
}
