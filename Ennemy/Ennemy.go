package ennemy

import (
	"Ebiten-Template-B1-22-23/tools"
	"math/rand"
	"time"

	"github.com/hajimehoshi/ebiten/v2"
)

type Ennemy struct {
	X, Y    float64
	Img     *ebiten.Image
	IsDrawn bool
	Speed   float64
}

func (ennemy *Ennemy) Update() error {
	rand.Seed(time.Now().UnixNano())
	rand := rand.Intn(4)
	switch rand {
	case 0:
		ennemy.X += ennemy.Speed
	case 1:
		ennemy.X -= ennemy.Speed
	case 2:
		ennemy.Y += ennemy.Speed
	case 3:
		ennemy.Y -= ennemy.Speed
	}
	ennemy.OutOfScreen()
	return nil
}

func (ennemy *Ennemy) OutOfScreen() {
	if ennemy.X > 710 {
		ennemy.X = 0
	}
	if ennemy.Y > 710 {
		ennemy.Y = 0
	}
	if ennemy.Y < 0 {
		ennemy.Y = 710
	}
	if ennemy.X < 0 {
		ennemy.X = 710
	}
}

func (ennemy *Ennemy) Draw(screen *ebiten.Image) {
	op := &ebiten.DrawImageOptions{}
	op.GeoM.Translate(ennemy.X, ennemy.Y)
	screen.DrawImage(ennemy.Img, op)

}

func NewEnnemy(path string, speed float64) *Ennemy {
	return &Ennemy{
		X:     float64(rand.Intn(720)),
		Y:     float64(rand.Intn(720)),
		Img:   tools.LoadImg(path),
		Speed: speed,
	}
}
