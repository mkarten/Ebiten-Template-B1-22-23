package tools

import (
	"image"
	"os"

	"github.com/hajimehoshi/ebiten/v2"
)

func LoadImg(s string) *ebiten.Image {
	file, err := os.Open(s)
	if err != nil {
		panic(err)
	}
	defer file.Close()
	img, _, err := image.Decode(file)
	if err != nil {
		panic(err)
	}
	return ebiten.NewImageFromImage(img)
}
