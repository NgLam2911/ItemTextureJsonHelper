package main

import (
	"encoding/json"
	"fmt"
	"os"
	"strings"
)

type textures struct {
	TexturePath string `json:"textures"`
}

type itemTextures struct {
	ResourcePackName string              `json:"resource_pack_name"`
	TextureName      string              `json:"texture_name"`
	TextureData      map[string]textures `json:"texture_data"`
}

var itemTexturesData itemTextures
var textureData map[string]textures
var resourcePackName string
var textureName string
var textureDirectory string

func init() {
	textureData = make(map[string]textures)
}

func SetResourcePackName(name string) {
	resourcePackName = name
}

func SetTextureName(name string) {
	textureName = name
}

func SetTextureDirectory(dir string) {
	textureDirectory = dir
}

func getItemTextureDirectory() string {
	return textureDirectory + "/textures/items"
}

func scanFiles(subDir string) {
	var files []os.DirEntry
	var err error
	files, err = os.ReadDir(getItemTextureDirectory() + subDir)
	check(err)

	for _, file := range files {
		if file.IsDir() {
			scanFiles(subDir + "/" + file.Name())
			continue
		}
		if strings.HasSuffix(file.Name(), ".png") {
			fileName := file.Name()[:len(file.Name())-4]
			texturePath := "textures/items" + subDir + "/" + fileName
			fmt.Println("Added " + texturePath)
			textureData[fileName] = textures{
				texturePath,
			}
		}
	}
}

func makeJson() {
	scanFiles("")
	itemTexturesData = itemTextures{
		resourcePackName,
		textureName,
		textureData,
	}
	file, err := json.MarshalIndent(itemTexturesData, "", "	")
	check(err)
	err = os.WriteFile(textureDirectory+"/textures/item_texture.json", file, 0644)
	fmt.Println("Generated json file at " + textureDirectory + "/textures/item_texture.json")
}
