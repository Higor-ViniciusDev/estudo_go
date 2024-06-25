package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 450
)

var AberturaSon rl.Sound
var Textura rl.Texture2D

type Plataforma struct {
}

func IniciaSomJogo() {

	rl.InitAudioDevice()

	AberturaSon = rl.LoadSound("assets/sons/abertura_som.ogg")

	rl.PlaySound(AberturaSon)
}

func CarregarImagemFundoJogo() {
	image := rl.LoadImage("assets/imagens/tela_inicial_jogo.png") // Carrega a imagem na CPU (RAM)
	Textura = rl.LoadTextureFromImage(image)                      // Converte a imagem em texture, e depos salva na GPU (VRAM)

	rl.UnloadImage(image)

	Textura.Width = int32(800)
	Textura.Height = int32(450)
}

func TelaJogoInicial() {
	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)
		rl.DrawText("Jogo Iniciado", 250, 100, 20, rl.NewColor(255, 255, 255, 255))

		rl.EndDrawing()
	}
}

func TelaLoading() {
	contadorRb := uint8(0)

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		if rl.IsKeyDown(rl.KeyEnter) {
			break
			// internal.CarrregarJogoInicial()
		}

		if contadorRb > 255 {
			contadorRb = 0
		}

		contadorRb += 6

		rl.DrawTexture(Textura, ScreenWidth/2-Textura.Width/2, 0, rl.White)

		rl.DrawText("Para Iniciar o Jogo tecle Enter", 250, 350, 20, rl.NewColor(contadorRb, contadorRb, contadorRb, contadorRb))

		rl.EndDrawing()
	}

}

func FecharJogo() {
	rl.CloseAudioDevice()
}
