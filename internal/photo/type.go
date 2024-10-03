package photo

import "image"

type Photo struct {
	Path string
}

type Changer struct {
	bg image.Image
	img image.Image
	savePath string
}

type CanvasOpts struct {
	sizeX int
	sizeY int
	bgColor string 
}

type Nimage struct {
	name string
	img image.Image
}