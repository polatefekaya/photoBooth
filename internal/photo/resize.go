package photo

import (
	"fmt"
	"image"
	"image/draw"
	"math"

	xdraw "golang.org/x/image/draw"
)

const float64EqualityThreshold float64 = 1e-2
 
func almostEqual(a, b float64) bool {
    return math.Abs(a - b) <= float64EqualityThreshold
}

func (c *Changer) Resize(mode string, useCanvas bool, fgXfactor float64, cOpts *CanvasOpts){
	//mode is can be "height" or "width"
	cBounds := image.Rect(0,0, cOpts.sizeX, cOpts.sizeY)
	sizes := orderSizes(mode, c.img, c.bg)

	var newBounds image.Rectangle
	if useCanvas {
		newBounds = cBounds
	} else {
		newBounds = sizes["big"].img.Bounds()
	}

	sizes["small"] = resizeImage(sizes["small"], mode, newBounds)
	//resize the other one if the resize reference is canvas
	if useCanvas {
		sizes["big"] = resizeImage(sizes["big"], mode, newBounds)
	}


	c.img = searchNimage("foreground", sizes)
	c.bg = searchNimage("background", sizes)

	if !almostEqual(fgXfactor , 1) {
		
		c.img = resizeImage(
			Nimage{"foreground", c.img},
			mode,
			image.Rect(0,0, int(float64(newBounds.Dx()) * fgXfactor) , int(float64(newBounds.Dy()) * fgXfactor))).img
	}
}

func searchNimage(val string, dic map[string]Nimage) (image.Image){
	for _, v := range dic {
		if v.name == val{
			return v.img
		}
	}
	fmt.Println("Value not found")
	errMsg := fmt.Sprintf("No named image found with given value: %s", val)
	panic(errMsg)
}

func orderSizes(mode string, img, bg image.Image) map[string]Nimage{
	imgSize := img.Bounds()
	bgSize := bg.Bounds()

	sizeMap := map[string]Nimage{"small": Nimage{}, "big": Nimage{}}
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
		sizeMap["small"] = Nimage{"foreground",img}
		sizeMap["big"] = Nimage{"background", bg}
	} else {
		sizeMap["small"] = Nimage{"background", bg}
		sizeMap["big"] = Nimage{"foreground",img}
	}

	return sizeMap
}

func resizeImage(nimg Nimage, mode string, newBounds image.Rectangle) Nimage {
	//resBounds := image.Rect(0,0, newWidth, newHeight)
	resBounds := calcBounds(mode, nimg.img.Bounds(), newBounds)
	resImg := image.NewRGBA(resBounds)

	xdraw.ApproxBiLinear.Scale(resImg, resBounds,nimg.img, nimg.img.Bounds(), draw.Over, nil)
	
	return Nimage{nimg.name, resImg}
}

func calcBounds(mode string, oldSizes image.Rectangle, newBounds image.Rectangle) image.Rectangle {
	var xFactor float64

	if mode == "height"{
		xFactor = float64(newBounds.Dy()) / float64(oldSizes.Dy())
		newWidth := int(float64(oldSizes.Dx()) * xFactor)
		return image.Rect(0,0, newWidth, newBounds.Dy())
	}

	xFactor = float64(newBounds.Dx()) / float64(oldSizes.Dx())
	newHeight := int(float64(oldSizes.Dy()) * xFactor)

	return image.Rect(0,0, newHeight, newBounds.Dx())
}