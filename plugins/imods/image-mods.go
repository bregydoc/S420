package imods

// ImageMod is a list of modifiers for your image
type ImageMod struct {
	Width      string       `json:"width"`
	Height     string       `json:"height"`
	Mode       ModeBehavior `json:"mode"`
	Dpr        int          `json:"dpr"`
	Crop       CropType     `json:"crop"`
	Extract    [4]int       `json:"extract"`
	Pad        [4]int       `json:"pad"`
	Background string       `json:"background"`
	Trim       int          `json:"trim"`
	Compress   bool         `json:"compress"`
	Brightness string       `json:"brightness"`
	Saturation string       `json:"saturation"`
	Hue        string       `json:"hue"`
	Tint       string       `json:"tint"`
	Grayscale  bool         `json:"grayscale"`
	Invert     bool         `json:"invert"`
	Rotate     int          `json:"rotate"`
	Flip       string       `json:"flip"`
	Blur       float64      `json:"blur"`
	Threshold  int          `json:"threshold"`
	Sharp      int          `json:"sharp"`
}
