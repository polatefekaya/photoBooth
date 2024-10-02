package photo

import (
	"bytes"
	"fmt"
	"image"
	"image/jpeg"
	"image/png"
	"os"
)

func GetPhoto() []byte {
	return make([]byte, 10)
}

func (p *Photo) DecodePhoto() (image.Image, string, error) {
	imgFile, err := os.Open(p.Path)
	if err != nil {
		fmt.Printf("Error while opening io")
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	return img, "jpg", err
}

func (p *Photo) DecodePng() (image.Image, string, error) {
	imgFile, err := os.Open(p.Path)
	if err != nil {
		fmt.Printf("Error while opening io")
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	return img, "png", err
}

func (p *Photo) EncodePhoto(img *image.Image) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := jpeg.Encode(buff, *img, nil)
	return buff.Bytes(), err
}

func (p *Photo) SavePng(imgBytes *[]byte) {
	imgFile, err := png.Decode(bytes.NewReader(*imgBytes))
	if err != nil {
		fmt.Println("Error while decoding in SavePng")
		fmt.Println(err)
	}
	out, _ := os.Create(p.Path)
	defer out.Close()

	err = png.Encode(out, imgFile)
	if err != nil {
		fmt.Printf("Error while encoding the saved image in SavePng\n")
	}
	pht := Photo{
		Path: "./resources/gen/image.png",
	}
	pht.ReplaceBackground()
}

func (p *Photo) SaveJpeg(imgBytes *[]byte) {
	imgFile, err := jpeg.Decode(bytes.NewReader(*imgBytes))
	if err != nil {
		fmt.Println("Error while decoding in savePhoto")
		fmt.Println(err)
	}

	out, _ := os.Create(p.Path)
	defer out.Close()

	var opts jpeg.Options
	opts.Quality = 95

	err = jpeg.Encode(out, imgFile, &opts)
	if err != nil {
		fmt.Printf("Error while encoding the saved image\n")
	}
	pht := Photo{
		Path: "./resources/gen/image.png",
	}
	pht.ReplaceBackground()
}

func (p *Photo) ReplaceBackground() {
	img, _, err := p.DecodePng()
	if err != nil {
		fmt.Printf("Error while Decoding photo in ReplaceBackground\n")
	}

	bgPht := Photo{
		Path: "./resources/background.jpg",
	}

	bg, _, err := bgPht.DecodePhoto()
	if err != nil {
		fmt.Printf("Error while decoding background image in ReplaceBackground\n")
	}

	bgSize := bg.Bounds().Size()
	imgSize := img.Bounds().Size()
	
	fmt.Printf("Generated Image Size x: %d y: %d\n", imgSize.X, imgSize.Y)
	fmt.Printf("Background Image Size x: %d y: %d\n", bgSize.X, bgSize.Y)

	cOpts := CanvasOpts{
		sizeX: 5120,
		sizeY: 3840,
		bgColor: "no",
	}

	changer := Changer{
		bg: bg,
		img: img,
		savePath: "./resources/gen/generated.png",
	}
	changer.BgChanger(&cOpts)
}
