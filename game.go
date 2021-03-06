package main

import (
	"image"
	"image/color"
	"image/png"
	"log"
	"os"

	"github.com/hajimehoshi/ebiten"
	colorful "github.com/lucasb-eyer/go-colorful"
	"github.com/nfnt/resize"
)

var brown, _ = colorful.Hex("#895C22")
var green, _ = colorful.Hex("#53D46B")

type islandDetails struct {
	sizeX   int
	sizeY   int
	coordsX int
	coordsY int
}

var x = islandDetails{sizeX: islandsHeight, sizeY: islandsWidth, coordsX: islandThreeX, coordsY: islandThreeY}
var y = islandDetails{sizeX: islandYLength, sizeY: islandsWidth, coordsX: islandTwoX, coordsY: islandTwoY}
var z = islandDetails{sizeX: islandXLength, sizeY: islandsWidth, coordsX: islandOneX, coordsY: islandOneY}

var arr = []islandDetails{x, y, z}

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

func (i *islandDetails) detectCollision() (b int) {
	if posX <= 0 {
		posX = 0
	} else if posX >= screenWidth-characterSize {
		posX = screenWidth - characterSize
	}
	if int(posX) >= i.coordsX && int(posY) <= i.coordsY {
		log.Printf("coordsX: %v coordsY: %v, posX: %v, posY: %v", i.coordsX, i.coordsY, posX, posY)
		onGround = true
		return i.coordsX - characterSize
		// } else if int(posX) <= i.sizeY && int(posY) < i.sizeY {
		//     onGround = true
		//     return islandTwoY - characterSize
	}
	onGround = true
	return lowerBound
}

func (i *islandDetails) drawEachLandmass(s *ebiten.Image) {
	mass, _ := ebiten.NewImage(i.sizeX, i.sizeY, ebiten.FilterNearest)
	mass.Fill(brown)
	foliage, _ := ebiten.NewImage(i.sizeX, i.sizeY/3, ebiten.FilterNearest)
	foliage.Fill(green)
	massOpts := &ebiten.DrawImageOptions{}
	massOpts.GeoM.Translate(float64(i.coordsX), float64(i.coordsY))
	foliageOpts := &ebiten.DrawImageOptions{}
	foliageOpts.GeoM.Translate(float64(i.coordsX), float64(i.coordsY-1))
	s.DrawImage(mass, massOpts)
	s.DrawImage(foliage, foliageOpts)
}

func drawLand(s *ebiten.Image) {
	if landmass == nil {
		landmass, _ = ebiten.NewImage(screenWidth, landHeight, ebiten.FilterNearest)
		grass, _ = ebiten.NewImage(screenWidth, grassHeight, ebiten.FilterNearest)
	}

	for _, v := range arr {
		v.drawEachLandmass(s)
	}

	grass.Fill(green)
	landmass.Fill(brown)
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
	oneOpts := &ebiten.DrawImageOptions{}
	twoOpts := &ebiten.DrawImageOptions{}
	one, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	two, err := ebiten.NewImageFromImage(img, ebiten.FilterNearest)
	if cloudX < screenWidth-20 || cloudX > 0 {
		cloudX += 0.5
	} else {
		cloudX = 0
	}
	oneOpts.GeoM.Translate(cloudX, 5)
	twoOpts.GeoM.Translate(cloudX/2, 100)
	s.DrawImage(one, oneOpts)
	s.DrawImage(two, twoOpts)
	logError(err)
}
