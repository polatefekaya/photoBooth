package photo

import (
	"image"
	"image/draw"
	"image/png"
	"log"
	"os"
	xdraw "golang.org/x/image/draw"
)

var canvas *image.RGBA = nil

func (c *Changer) BgChanger(cOpts *CanvasOpts) {
	c.createCanvas(cOpts)
	c.placeBg()
	c.placeImg()
	c.saveImg()
}

func (c *Changer) createCanvas(co *CanvasOpts) {
	canvas = image.NewRGBA(image.Rect(0, 0, co.sizeX, co.sizeY))

}

func bgOffsetX(bgX int) int {
	offset := (bgX -canvas.Rect.Dx()) / 2
	return offset
}

func (c *Changer) placeBg() {
	if canvas == nil {
		log.Panicf("Canvas is nil, maybe calling createCanvas before this method works")
	}
	bgSize := c.bg.Bounds()
	draw.Draw(
		canvas,
		canvas.Bounds(),
		c.bg,
		image.Point{
			X: bgOffsetX(bgSize.Dx()), 
			Y: 0},
		draw.Src)

	// m := image.NewRGBA(image.Rect(0, 0, 640, 480))
	// blue := color.RGBA{0, 0, 255, 255}
	// draw.Draw(m, m.Bounds(), &image.Uniform{blue}, image.ZP, draw.Src)
}

func imgOffsetX(imgX int) int {
	mp := canvas.Rect.Dx() / 2
	mp -= imgX / 2
	return int(mp)
}
func imgOffsetY(imgY int) int {
	bp := canvas.Rect.Dy()
	bp -= imgY 
	return int(bp)
}

func (c *Changer) placeImg() {
	if canvas == nil {
		log.Panicf("Canvas is nil, maybe calling createCanvas before this method works")
	}
	//bgSize := c.bg.Bounds()
	imgSize := c.img.Bounds()

	draw.Draw(
		canvas,
		c.img.Bounds().Add(
			image.Point{
				X: imgOffsetX(imgSize.Dx()),
				Y: imgOffsetY(imgSize.Dy())}),
			c.img, image.Point{},
			draw.Over)

}

func (c *Changer) saveImg() {
	if canvas == nil {
		log.Panicf("Canvas is nil, maybe calling createCanvas before this method works")
	}

	output, err := os.Create(c.savePath)
	if err != nil {
		log.Panicf("Error occured while creating a output file")
	}
	defer output.Close()

	err = png.Encode(output, canvas)
	if err != nil {
		log.Panicf("Error while encoding png file for output")
	}
}