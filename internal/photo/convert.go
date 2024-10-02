package photo

import (
	"bytes"
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

func (p *Photo) EncodePhoto(img *image.Image) ([]byte, error) {
	buff := new(bytes.Buffer)
	err := jpeg.Encode(buff, *img, nil)
	return buff.Bytes(), err
}

// func (p *Photo) SavePhoto(imgStr *string) {
// 	decoded, err := base64.StdEncoding.DecodeString(*imgStr)
// 	imgFile, err := jpeg.Decode(bytes.NewReader(*imgByte))
// 	if err != nil {
// 		fmt.Println("Error while decoding in savePhoto")
// 		fmt.Println(err)
// 	}

// 	out, _ := os.Create(p.Path)
// 	defer out.Close()

// 	var opts jpeg.Options
// 	opts.Quality = 95

// 	err = jpeg.Encode(out, imgFile, &opts)
// 	if err != nil{
// 		fmt.Printf("Error while encoding the saved image")
// 	}
// }

func (p *Photo) SavePhoto(imgBytes *[]byte) {
	// decoded, err := base64.StdEncoding.DecodeString(string(*imgBytes))
	// if err != nil {
	// 	fmt.Println("Error while decoding base64 in savePhoto")
	// 	fmt.Println(err)
	// }

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
	if err != nil{
		fmt.Printf("Error while encoding the saved image")
	}
}