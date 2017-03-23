// Copyright 2017 Whit Marbut. All rights reserved.
// License information may be found in the LICENSE file.

package goxbm

import (
	"bytes"
	"fmt"
	"image"
	"io"
)

// Encoder configures encoding XBM images.
// More information about XBM may be found on Wikipedia:
// https://en.wikipedia.org/wiki/X_BitMap
type Encoder struct {
	CompressionLevel CompressionLevel
}

type CompressionLevel int

const (
	DefaultCompression CompressionLevel = 0
)

const MAX_COL uint32 = 65535
const BYTE_SIZE = 8

var PIX_VALUES []byte

func init() {
	PIX_VALUES = []byte{1, 2, 4, 8, 16, 32, 64, 128}
}

func Encode(w io.Writer, m image.Image) error {
	var encoder Encoder
	return encoder.Encode(w, m)
}

func (enc *Encoder) Encode(w io.Writer, img image.Image) error {
	width := img.Bounds().Dx()
	height := img.Bounds().Dy()
	fullSize := width * height

	strBuf := bytes.NewBufferString("#define xbm_width ") //TODO optimize to know exact size of buffer ahead of time
	fmt.Fprintf(strBuf, "%d\n", width)
	fmt.Fprintf(strBuf, "#define xbm_height %d\nstatic unsigned char xbm_bits[] = {", width)

	var i = 0
	for y := 0; y < height; y++ { /*Step through each line of the image vertically*/
		var curByte byte             /*We start a new byte on a new line, even if we haven't used all the bits of the old byte */
		for x := 0; x < width; x++ { /*Step through each pixel on the current line*/
			i++
			bitPos := (x % BYTE_SIZE)
			if bitPos == 0 && x != 0 {
				fmt.Fprintf(strBuf, formatByte(curByte))
				if i != fullSize {
					fmt.Fprintf(strBuf, ", ")
				}
				curByte = 0
			}
			r, g, b, a := img.At(x, y).RGBA()
			if r != MAX_COL || g != MAX_COL || b != MAX_COL || a != MAX_COL { //its black
				curByte = curByte | PIX_VALUES[bitPos]
			}
		}
		fmt.Fprintf(strBuf, formatByte(curByte))
		if i != fullSize {
			fmt.Fprintf(strBuf, ", ")
		}
	}
	fmt.Fprintf(strBuf, "};\n")

	_, err := io.Copy(w, strBuf)

	return err
}

func formatByte(curByte byte) string {
	str := fmt.Sprintf("%x", curByte)
	if len(str) == 1 {
		str = fmt.Sprintf("0%s", str)
	}
	return fmt.Sprintf("0x%s", str)
}
