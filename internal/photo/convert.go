package photo

import (
	"fmt"
	"image"
	"image/jpeg"
	"os"
)

func GetPhoto() []byte {
	return make([]byte, 10)
}

func (p *Photo) DecodePhoto() (image.Image, string ,error) {
	imgFile, err := os.Open(p.Path)
	if(err != nil){
		fmt.Printf("Error while opening io")
	}
	defer imgFile.Close()

	img, err := jpeg.Decode(imgFile)
	return img, "jpg", err
}