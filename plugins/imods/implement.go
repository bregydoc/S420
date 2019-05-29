package imods

import (
	"bytes"
	"image"
	"image/jpeg"
	"image/png"
	"net/http"
	"strings"

	s420 "github.com/bregydoc/S420"
)

// Plugin is a frame struct to create a new image plugin
type Plugin struct {
}

// NewImagePlugin returns a new pointer instance struct
func NewImagePlugin() *Plugin {
	return &Plugin{}
}

// ProcessFile implements the plugin interface of s420
func (p *Plugin) ProcessFile(r *http.Request, cType s420.ContentType, in []byte) []byte {

	if !strings.Contains(string(cType), "image") {
		return in
	}

	inBuf := bytes.NewBuffer(in)
	var inImage, outImage image.Image
	var err error

	if strings.Contains(string(cType), "png") { // * PNG
		inImage, err = png.Decode(inBuf)
		if err != nil {
			return in
		}
	} else if strings.Contains(string(cType), "jpg") || strings.Contains(string(cType), "jpeg") { // * JPG, JPEG
		inImage, err = jpeg.Decode(inBuf)
		if err != nil {
			return in
		}
	} else {
		inImage, err = png.Decode(inBuf) // * DEFAULT (PNG)
		if err != nil {
			return in
		}
	}

	im := parseURLObject(r.URL)
	outImage = im.apply(inImage)

	outBuf := bytes.NewBuffer([]byte{})

	if strings.Contains(string(cType), "png") { // * PNG
		err = png.Encode(outBuf, outImage)
		if err != nil {
			return in
		}
	} else if strings.Contains(string(cType), "jpg") || strings.Contains(string(cType), "jpeg") { // * JPG, JPEG
		err = jpeg.Encode(outBuf, outImage, nil)
		if err != nil {
			return in
		}
	} else {
		err = png.Encode(outBuf, outImage) // * DEFAULT (PNG)
		if err != nil {
			return in
		}
	}

	return outBuf.Bytes()
}
