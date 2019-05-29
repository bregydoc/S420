package imods

func (im *ImageMod) isVoid() bool {
	if im.Width != "" {
		return false
	}

	if im.Height != "" {
		return false
	}

	if im.Mode != "" {
		return false
	}

	if im.Dpr != 0 {
		return false
	}

	if im.Crop != "" {
		return false
	}

	if im.Extract != [4]int{0, 0, 0, 0} || im.Extract != [4]int{} {
		return false
	}

	if im.Pad != [4]int{0, 0, 0, 0} || im.Extract != [4]int{} {
		return false
	}

	if im.Background != "" {
		return false
	}

	if im.Trim != 0 {
		return false
	}

	if im.Compress != false {
		return false
	}

	if im.Brightness != "" {
		return false
	}

	if im.Saturation != "" {
		return false
	}

	if im.Hue != "" {
		return false
	}

	if im.Tint != "" {
		return false
	}

	if im.Grayscale != false {
		return false
	}

	if im.Invert != false {
		return false
	}

	if im.Rotate != 0 {
		return false
	}

	if im.Flip != "" {
		return false
	}

	if im.Blur != 0.0 {
		return false
	}

	if im.Threshold != 0 {
		return false
	}

	if im.Sharp != 0 {
		return false
	}

	return true
}
