package main

import (
	"fmt"
	"image"
	"image/draw"
	"image/jpeg"
	"image/png"
	"log"
	"os"

	"github.com/nfnt/resize"
)

func main() {
	for _, i := range []int{1, 2, 3, 4} {
		path, err := AddFlag("avatar.jpeg", i)
		if err != nil {
			log.Println(err)
		}
		log.Println(path)
	}
}

func AddFlag(avatarPath string, num int) (scrName string, err error) {
	paddingX, paddingY := 20, 20 //图片边框长度

	scrName = fmt.Sprintf("./output/flag_avatar%d.png", num)
	flagPath := fmt.Sprintf("./src/flag%d.png", num)

	file, err := os.Create(scrName)
	if err != nil {
		return
	}
	defer file.Close()

	flagFile, err := os.Open(flagPath)
	if err != nil {
		return
	}
	defer flagFile.Close()

	flagImg, err := png.Decode(flagFile)
	if err != nil {
		return
	}

	avatarFile, err := os.Open(avatarPath)
	if err != nil {
		return
	}
	defer avatarFile.Close()

	avatarImg, err := jpeg.Decode(avatarFile)
	if err != nil {
		return
	}

	//resize avatar
	newImage := resize.Resize(uint(flagImg.Bounds().Dx()-20), uint(flagImg.Bounds().Dy()-20), avatarImg, resize.Lanczos3)

	resultPng := image.NewRGBA(image.Rect(0, 0, flagImg.Bounds().Dx(), flagImg.Bounds().Dy()))
	draw.Draw(resultPng, resultPng.Bounds(), newImage, newImage.Bounds().Min.Sub(image.Pt(paddingX, paddingY)), draw.Over)
	draw.Draw(resultPng, resultPng.Bounds(), flagImg, flagImg.Bounds().Min, draw.Over)

	err = png.Encode(file, resultPng)
	if err != nil {
		return
	}
	return
}
