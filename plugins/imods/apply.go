package imods

import (
	"image"
	"strconv"
	"strings"

	"github.com/disintegration/imaging"
)

func (im *ImageMod) apply(img image.Image) image.Image {
	if im.isVoid() {
		return img
	}

	result := img

	w, wOk := strconv.ParseInt(strings.ReplaceAll(im.Width, "px", ""), 10, 64)
	h, hOk := strconv.ParseInt(strings.ReplaceAll(im.Height, "px", ""), 10, 64)

	if wOk == nil && hOk == nil {
		result = imaging.Resize(img, int(w), int(h), imaging.Lanczos)
	}

	return result
}
