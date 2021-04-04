package main

import "C"

import (
	"bytes"
	"errors"
	"fmt"
	"github.com/corona10/goimagehash"
	"image"
	_ "image/gif"
	_ "image/jpeg"
	_ "image/png"
	"log"
)
//https://medium.com/learning-the-go-programming-language/calling-go-functions-from-other-languages-4c7d8bcc69bf

func handle(hash *goimagehash.ImageHash, err error) (uint64, error)  {
	if err != nil {
		return 0, err
	}
	return hash.GetHash(), nil
}

//export GetHash
func GetHash(imageBytes []byte, algorithm string) (uint64, error)  {

	img, format, readError := image.Decode(bytes.NewReader(imageBytes))
	if readError != nil {
		return 0, readError
	}
	log.Println(fmt.Sprintf("Image is in %v format", format))

	if algorithm == "aHash" {
		return handle(goimagehash.AverageHash(img))
	} else if algorithm == "dHash" {
		return handle(goimagehash.DifferenceHash(img))
	} else if algorithm == "pHash" {
		return handle(goimagehash.PerceptionHash(img))
	}
	return 0, errors.New("unsupported hash type. Available hashes are aHash, dHash and pHash")
}

func main() {}