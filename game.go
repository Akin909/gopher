package main

import (
	"image"
	"image/color"
	"image/png"
	"os"

	"github.com/hajimehoshi/ebiten"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
)

func drawCharacter(s *ebiten.Image) {
	//If square does not already exist initialise it - this way a new image is not created each time
	if character == nil {
		character, _ = ebiten.NewImage(characterSize, characterSize, ebiten.FilterNearest)
	}
	character.Fill(color.NRGBA{0xff, 0x00, 0x00, 0xff})
	opts := &ebiten.DrawImageOptions{}

	opts.GeoM.Translate(posX, posY)

	s.DrawImage(character, opts)
}

func drawLand(s *ebiten.Image) {
	if landmass == nil {
		landmass, _ = ebiten.NewImage(screenWidth, landHeight, ebiten.FilterNearest)
		grass, _ = ebiten.NewImage(screenWidth, grassHeight, ebiten.FilterNearest)
	}
	brown, err := colorful.Hex("#895C22")
	logError(err)
	green, err := colorful.Hex("#53D46B")
	logError(err)

	landmass.Fill(brown)
	grass.Fill(green)

	opts := &ebiten.DrawImageOptions{}
	opts.GeoM.Translate(0, screenHeight-landHeight)
	s.DrawImage(landmass, opts)

	grassOpts := &ebiten.DrawImageOptions{}
	grassOpts.GeoM.Translate(0, lowerBound+characterSize)
	s.DrawImage(grass, grassOpts)
}

func getImage() (i image.Image) {
	file, err := os.Open("assets/cloud.png")
	defer file.Close()
	logError(err)

	img, err := png.Decode(file)
	logError(err)

	resized := resize.Resize(50, 0, img, resize.Lanczos3)
	return resized
}

func drawClouds(s *ebiten.Image) {
	img := getImage()
	opts := &ebiten.DrawImageOptions{}
	ebimg, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)

	s.DrawImage(ebimg, opts)
	logError(err)
}