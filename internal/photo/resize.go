package photo

import(
	"image"
	"image/draw"
	xdraw "golang.org/x/image/draw"
)

func (c *Changer) Resize(){
	sizes := findSizes("height", c.img, c.bg)

	var newHeight, newWidth int

	newHeight = sizes["big"].Bounds().Dy()
	newWidth = sizes["big"].Bounds().Dx()

	sizes["small"] = resizeImage(sizes["small"], newHeight, newWidth)

	
}

func findSizes(mode string,img, bg image.Image) map[string]image.Image{
	imgSize := img.Bounds()
	bgSize := bg.Bounds()

	var sizeMap map[string]image.Image
	var cond bool 

	switch mode {
		case "height":
			cond = imgSize.Dy() < bgSize.Dy()
		case "width":
			cond = imgSize.Dx() < bgSize.Dx()
		default:
			cond = false
	}

	if cond {
		sizeMap["small"] = img
		sizeMap["big"] = bg
	} else {
		sizeMap["small"] = bg
		sizeMap["big"] = img
	}

	return sizeMap
}

func resizeImage(img image.Image, newHeight, newWidth int) image.Image {
	resBounds := image.Rect(0,0, newWidth, newHeight)
	resImg := image.NewRGBA(resBounds)

	xdraw.ApproxBiLinear.Scale(resImg, resBounds,img, img.Bounds(), draw.Over, nil)
	
	return resImg
}
