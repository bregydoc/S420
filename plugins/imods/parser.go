package imods

import (
	"net/url"
	"strconv"
	"strings"
)

func parseURLObject(url *url.URL) *ImageMod {
	return parseQueryParamToImageModifiers(url.RawQuery)
}

func parseURLToImageModifiers(u string) *ImageMod {
	parsed, err := url.Parse(u)
	if err != nil {
		return nil
	}

	return parseQueryParamToImageModifiers(parsed.RawQuery)
}

func parseQueryParamToImageModifiers(query string) *ImageMod {
	values, err := url.ParseQuery(query)
	if err != nil {
		return nil
	}

	modifiers := new(ImageMod)

	w := values.Get("w")
	if w == "" {
		w = values.Get("width")
	}

	h := values.Get("h")
	if h == "" {
		h = values.Get("height")
	}

	mode := values.Get("mode")
	dpr := values.Get("dpr")
	crop := values.Get("crop")
	extract := values.Get("extract")
	pad := values.Get("pad")
	background := values.Get("background")
	trim := values.Get("trim")
	compress := values.Get("compress")
	brightness := values.Get("brightness")
	saturation := values.Get("saturation")
	hue := values.Get("hue")
	tint := values.Get("tint")
	grayscale := values.Get("grayscale")
	invert := values.Get("invert")
	rotate := values.Get("rotate")
	flip := values.Get("flip")
	blur := values.Get("blur")
	threshold := values.Get("threshold")
	sharp := values.Get("sharp")

	modifiers.Width = w
	modifiers.Height = h

	modifiers.Mode = ModeBehavior(mode)
	modifiers.Dpr, _ = strconv.Atoi(dpr)

	modifiers.Crop = CropType(crop)

	extractArr := [4]int{}
	cks := strings.Split(strings.ReplaceAll(extract, " ", ""), ",")
	if len(cks) == 4 {
		for i := range cks {
			extractArr[i], _ = strconv.Atoi(cks[i])
		}
	}
	modifiers.Extract = extractArr

	padArr := [4]int{}
	cks = strings.Split(strings.ReplaceAll(pad, " ", ""), ",")
	if len(cks) == 4 {
		for i := range cks {
			padArr[i], _ = strconv.Atoi(cks[i])
		}
	}

	modifiers.Pad = padArr
	modifiers.Background = background
	modifiers.Trim, _ = strconv.Atoi(trim)
	modifiers.Compress, _ = strconv.ParseBool(compress)
	modifiers.Brightness = brightness
	modifiers.Saturation = saturation
	modifiers.Hue = hue
	modifiers.Tint = tint
	modifiers.Grayscale, _ = strconv.ParseBool(grayscale)
	modifiers.Invert, _ = strconv.ParseBool(invert)
	modifiers.Rotate, _ = strconv.Atoi(rotate)
	modifiers.Flip = flip
	modifiers.Blur, _ = strconv.ParseFloat(blur, 64)
	modifiers.Threshold, _ = strconv.Atoi(threshold)
	modifiers.Sharp, _ = strconv.Atoi(sharp)

	if modifiers.isVoid() {
		return nil
	}
	return modifiers
}
