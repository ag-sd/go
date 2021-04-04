package main

import "C"

import (
	"bytes"
	"encoding/base64"
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"io/ioutil"
	"log"
)

func handle(hash interface{}, err error) uint64 {
	if err != nil {
		panic(err)
	}
	switch value := hash.(type) {
	case *goimagehash.ImageHash:
		return value.GetHash()
	case uint64:
		return value
	default:
		panic(fmt.Sprintf("Unexpected hash type received %T for %v", value, value))
	}
}

//export GetHash
func GetHash(imageBytes []byte, algorithm string) uint64 {

	img, format, readError := image.Decode(bytes.NewReader(imageBytes))
	if readError != nil {
		return handle(0, readError)
	}
	log.Println(fmt.Sprintf("Image is in %v format", format))

	if algorithm == "aHash" {
		return handle(goimagehash.AverageHash(img))
	} else if algorithm == "dHash" {
		return handle(goimagehash.DifferenceHash(img))
	} else if algorithm == "pHash" {
		return handle(goimagehash.PerceptionHash(img))
	}
	return handle(0, errors.New("unsupported hash type. Available hashes are aHash, dHash and pHash"))
}

//export GetHashFromBase64Data
func GetHashFromBase64Data(base64ImageData string, algorithm string) uint64 {
	imageBytes, err := base64.StdEncoding.DecodeString(base64ImageData)
	if err != nil {
		return handle(0, err)
	}
	return GetHash(imageBytes, algorithm)
}

//export GetHashFromImageFile
func GetHashFromImageFile(imageFile string, algorithm string) uint64 {
	file, err := ioutil.ReadFile(imageFile)
	if err != nil {
		return handle(0, err)
	}
	return GetHash(file, algorithm)
}

func main() {}
