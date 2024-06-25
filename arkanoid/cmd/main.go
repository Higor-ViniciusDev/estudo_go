package main

import (
	"estudo_go/arkanoid/internal"

	rl "github.com/gen2brain/raylib-go/raylib"
)

func main() {
	rl.InitWindow(internal.ScreenWidth, internal.ScreenHeight, "Arkanoid")
	defer rl.CloseWindow()

	internal.IniciaSomJogo()
	internal.CarregarImagemFundoJogo()

	rl.SetTargetFPS(60)

	internal.TelaLoading()
	internal.TelaJogoInicial()

	rl.UnloadTexture(internal.Textura)
	rl.UnloadSound(internal.AberturaSon)

	rl.CloseAudioDevice()

	rl.CloseWindow()
}
