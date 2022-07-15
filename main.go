package main

import (
	"fmt"
)

func check(e error) {
	if e != nil {
		panic(e)
	}
}

func main() {
	var resourcePackName string
	var textureName string
	var textureDirectory string
	fmt.Print("Resource Pack Name: ")
	_, err := fmt.Scan(&resourcePackName)
	check(err)
	SetResourcePackName(resourcePackName)
	fmt.Print("Texture Name: ")
	_, err = fmt.Scan(&textureName)
	check(err)
	SetTextureName(textureName)
	fmt.Print("Texture Directory: ")
	_, err = fmt.Scan(&textureDirectory)
	check(err)
	SetTextureDirectory(textureDirectory)
	fmt.Println("Making item_texture.json...")
	makeJson()
}
