package main

import (
	"image/color"
	_ "image/png"
	"log"

	"Ebiten-Template-B1-22-23/ennemy"
	"Ebiten-Template-B1-22-23/player"

	"github.com/hajimehoshi/ebiten/v2"
	"github.com/hajimehoshi/ebiten/v2/ebitenutil"
	"github.com/hajimehoshi/ebiten/v2/inpututil"
)

const (
	screenWidth  = 280
	screenHeight = 280
	resolWidth   = 720
	resolHeight  = 720
)

type Game struct {
	Player *player.Player
	Ennemy []*ennemy.Ennemy
	Paused bool
}

func (game *Game) Update() error {
	if inpututil.IsKeyJustPressed(ebiten.KeyEscape) {
		game.Paused = !game.Paused
	}
	if !game.Paused {
		game.Player.Update()
		for _, ennemy := range game.Ennemy {
			ennemy.Update()
		}
	}

	return nil
}

func (game *Game) Draw(screen *ebiten.Image) {
	screen.Fill(color.Black)

	game.Player.Draw(screen)
	for _, ennemy := range game.Ennemy {
		ennemy.Draw(screen)
	}
	if game.Paused {
		scale := 5.0
		Pause := ebiten.NewImage(100, 100)
		ebitenutil.DebugPrint(Pause, "Paused")
		op := &ebiten.DrawImageOptions{}
		op.GeoM.Translate((270)/scale, 320/scale)
		op.GeoM.Scale(scale, scale)
		screen.DrawImage(Pause, op)

	}
}

func (game *Game) Layout(outsideWidth, outsideHeight int) (screenWidth, screenHeight int) {
	return resolWidth, resolHeight
}

func main() {
	game := new(Game)
	NbEnemys := 10
	game.Player = player.NewPlayer("img/Player.png", 2.5)
	for i := 0; i < NbEnemys; i++ {
		game.Ennemy = append(game.Ennemy, ennemy.NewEnnemy("img/Ennemy.png", 5))
	}
	ebiten.SetCursorMode(ebiten.CursorModeVisible)
	ebiten.SetWindowSize(screenWidth*2, screenHeight*2)
	ebiten.SetWindowTitle("Cours ebiten")
	ebiten.SetTPS(60)
	ebiten.SetWindowResizingMode(ebiten.WindowResizingModeEnabled)
	if err := ebiten.RunGame(game); err != nil {
		log.Fatal(err)
	}
}
