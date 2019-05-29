package imods

// ModeBehavior represents a mode for your image modifier
type ModeBehavior string

// Fit is a valid mode when you modify the image with width or height
const Fit ModeBehavior = "fit"

// Crop is a valid mode when you modify the image with width or height
const Crop ModeBehavior = "crop"

// Stretch is a valid mode when you modify the image with width or height
const Stretch ModeBehavior = "stretch"

// Min is a valid mode when you modify the image with width or height
const Min ModeBehavior = "min"

// Max is a valid mode when you modify the image with width or height
const Max ModeBehavior = "max"

// Fill is a valid mode when you modify the image with width or height
const Fill ModeBehavior = "fill"
