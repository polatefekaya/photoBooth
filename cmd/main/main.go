package main

import (
	"fmt"
	"github.com/polatefekaya/photoBooth/internal/photo"
	"github.com/polatefekaya/photoBooth/internal/messaging"
)

func main()  {
	consume()
}
func consume(){
	messaging.StartConnection()
}

func test() {
	photoBytes := photo.GetPhoto()
	fmt.Printf("%d", cap(photoBytes))

	pht := photo.Photo{
		Path: "./resources/image.JPG",
	}

	img, imgExt, err := pht.DecodePhoto()
	if err != nil {
		fmt.Printf("error while decoding")
		fmt.Println(err.Error())
	}
	fmt.Printf("%s", imgExt)
	fmt.Println(img)
}