package player

import (
	"Ebiten-Template-B1-22-23/tools"

	"github.com/hajimehoshi/ebiten/v2"
)

type Player struct {
	X, Y    float64
	Img     *ebiten.Image
	IsDrawn bool
	Speed   float64
}

func (player *Player) Update() error {
	if ebiten.IsKeyPressed(ebiten.KeyW) {
		player.Y -= 2.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyS) {
		player.Y += 2.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyA) {
		player.X -= 2.5
	}
	if ebiten.IsKeyPressed(ebiten.KeyD) {
		player.X += 2.5
	}
	player.OutOfScreen()
	return nil
}

func (p *Player) OutOfScreen() {
	if p.X > 710 {
		p.X = 0
	}
	if p.Y > 710 {
		p.Y = 0
	}
	if p.Y < 0 {
		p.Y = 710
	}
	if p.X < 0 {
		p.X = 710
	}
}

func (p *Player) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(p.X, p.Y)
	screen.DrawImage(p.Img, op)
}

func NewPlayer(path string, speed float64) *Player {
	return &Player{
		Img:   tools.LoadImg(path),
		Speed: speed,
	}
}
