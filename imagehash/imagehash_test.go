package main

import (
	"encoding/base64"
	"io/ioutil"
	"testing"
)

var testFile string = "resource/test_image.jpg"
var aHash uint64 = 9169063992688115967
var dHash uint64 = 17505769358845671865
var pHash uint64 = 10053688355855114876

func TestGetHashFromImageFile(t *testing.T) {
	hash := GetHashFromImageFile(testFile, "aHash")
	if hash != aHash {
		t.Errorf("aHash mismatch. Expected %v but got %v", aHash, hash)
	}
}

func TestInvalidAlgorithm(t *testing.T) {
	defer func() {
		if r := recover(); r == nil {
			t.Errorf("An error was expected as xHash is not a valid algorithm")
		}
	}()
	GetHashFromImageFile(testFile, "xHash")
}

func TestGetHashFromBase64Data(t *testing.T) {
	file, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Expected no error while reading file but got %v", err)
	}
	hash := GetHashFromBase64Data(base64.StdEncoding.EncodeToString(file), "dHash")
	if hash != dHash {
		t.Errorf("dHash mismatch. Expected %v but got %v", dHash, hash)
	}
}

func TestGetHash(t *testing.T) {
	file, err := ioutil.ReadFile(testFile)
	if err != nil {
		t.Errorf("Expected no error while reading file but got %v", err)
	}
	hash := GetHash(file, "pHash")
	if hash != pHash {
		t.Errorf("pHash mismatch. Expected %v but got %v", pHash, hash)
	}
}
