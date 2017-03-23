package goxbm

import (
	"bytes"
	"image"
	"image/png"
	"io/ioutil"
	"os"
	"testing"
)

func TestEncodeAgainstKnownValue(t *testing.T) {
	png := getPngImage()
	existingXbm := getExistingXbmBytes()

	buf := bytes.NewBuffer([]byte{})
	err := Encode(buf, png)
	if err != nil {
		t.Fatalf("Error encoding xbm: %s", err)
	}
	if bytes.Compare(buf.Bytes(), existingXbm) != 0 {
		t.Error("Byte comparison failed")
	}
}

func getExistingXbmBytes() []byte {
	imgBytes, err := ioutil.ReadFile("test_image.xbm")
	if err != nil {
		panic(err)
	}
	return imgBytes
}

func getPngImage() image.Image {
	imgFile, err := os.Open("test_image.png")
	if err != nil {
		panic(err)
	}
	defer imgFile.Close()

	img, err := png.Decode(imgFile)
	if err != nil {
		panic(err)
	}
	return img
}
