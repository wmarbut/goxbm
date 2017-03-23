package goxbm

import (
	"bytes"
	"image"
	"image/png"
	"testing"
)

func TestEncodeAgainstKnownValue(t *testing.T) {
	png := getPngImage1()
	existingXbm := getExistingXbmBytes1()

	buf := bytes.NewBuffer([]byte{})
	err := Encode(buf, png)
	if err != nil {
		t.Fatalf("Error encoding xbm: %s", err)
	}
	if bytes.Compare(buf.Bytes(), existingXbm) != 0 {
		t.Error("Byte comparison failed on image 1")
	}

	png = getPngImage2()
	existingXbm = getExistingXbmBytes2()
	buf = bytes.NewBuffer([]byte{})
	err = Encode(buf, png)

	if err != nil {
		t.Fatalf("Error encoding xbm: %s", err)
	}
	if bytes.Compare(buf.Bytes(), existingXbm) != 0 {
		t.Error("Byte comparison failed on image 2")
	}
}

func getExistingXbmBytes1() []byte {
	imgBytes, err := Asset("test_image.xbm")
	if err != nil {
		panic(err)
	}
	return imgBytes
}

func getExistingXbmBytes2() []byte {
	imgBytes, err := Asset("test_image2.xbm")
	if err != nil {
		panic(err)
	}
	return imgBytes
}

func getPngImage1() image.Image {
	imgBytes, err := Asset("test_image.png")
	if err != nil {
		panic(err)
	}
	imgReader := bytes.NewReader(imgBytes)

	img, err := png.Decode(imgReader)
	if err != nil {
		panic(err)
	}
	return img
}

func getPngImage2() image.Image {
	imgBytes, err := Asset("test_image2.png")
	if err != nil {
		panic(err)
	}
	imgReader := bytes.NewReader(imgBytes)

	img, err := png.Decode(imgReader)
	if err != nil {
		panic(err)
	}
	return img
}
