package internal

import (
	rl "github.com/gen2brain/raylib-go/raylib"
)

const (
	ScreenWidth  = 800
	ScreenHeight = 450
)

type Game struct {
	FimJogo   bool
	Pause     bool
	Jogador   Jogador
	Bola      Bola
	Brick     [LINHA_POR_BLOCO][BLOCOS_POR_LINHA]Bloco
	BlocoSize rl.Vector2
}

var AberturaSon rl.Sound
var Textura rl.Texture2D

type Plataforma struct {
}

func IniciaSomJogo(caminhoSom string) {

	rl.InitAudioDevice()

	AberturaSon = rl.LoadSound(caminhoSom)

	rl.PlaySound(AberturaSon)
}

func CarregarImagemFundoJogo() {
	image := rl.LoadImage("assets/imagens/tela_inicial_jogo.png") // Carrega a imagem na CPU (RAM)
	Textura = rl.LoadTextureFromImage(image)                      // Converte a imagem em texture, e depos salva na GPU (VRAM)

	rl.UnloadImage(image)

	Textura.Width = int32(800)
	Textura.Height = int32(450)
}

func TelaJogoInicial(g *Game) {
	IniciaSomJogo("assets/sons/som_jogo.ogg")

	for !rl.WindowShouldClose() {
		rl.BeginDrawing()

		rl.ClearBackground(rl.Black)

		// Plataforma jogador
		rl.DrawRectangle(int32(g.Jogador.Position.X-g.Jogador.Size.X/2), int32(g.Jogador.Position.Y-g.Jogador.Size.Y/2),
			int32(g.Jogador.Size.X), int32(g.Jogador.Size.Y), rl.White)

		rl.EndDrawing()
	}

	rl.UnloadTexture(Textura)
	rl.UnloadSound(AberturaSon)
	rl.CloseAudioDevice()
}

func TelaLoading() {
	contadorRb := uint8(0)

	IniciaSomJogo("assets/sons/abertura_som.ogg")
	CarregarImagemFundoJogo()

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

	rl.UnloadTexture(Textura)
	rl.UnloadSound(AberturaSon)
	rl.CloseAudioDevice()

}

func FecharJogo() {
	rl.CloseAudioDevice()
}

func InicializaIntesJogo(g *Game) {
	//Inicializa Posição mouse jogando e seta
	g.Jogador.Vida = JOGADOR_MAXIMA_VIDA
	g.Jogador.Position = rl.Vector2{}
	g.Jogador.Position.X = float32(ScreenWidth / 2)
	g.Jogador.Position.Y = float32(ScreenHeight/2) + 120

	g.Jogador.Size = rl.Vector2{}
	g.Jogador.Size.X = float32(ScreenWidth / 10)
	g.Jogador.Size.Y = 20

	// Inicializa Bola
	g.Bola.Position = rl.Vector2{}
	g.Bola.Position.X = float32(ScreenWidth / 2)
	g.Bola.Position.Y = float32(ScreenHeight*7/8 - 30)

	g.Bola.Velocidade = rl.Vector2{}
	g.Bola.Radius = 7
	g.Bola.Ativo = false
}
