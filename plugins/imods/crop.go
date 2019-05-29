package imods

// CropType define the way how your images will be cropped
type CropType string

// Entropy is an option to crop your image
const Entropy CropType = "fill"

// Smart is an option to crop your image
const Smart CropType = "smart"

// Top is an option to crop your image
const Top CropType = "top"

// Topleft is an option to crop your image
const Topleft CropType = "topleft"

// Left is an option to crop your image
const Left CropType = "left"

// Bottomleft is an option to crop your image
const Bottomleft CropType = "bottomleft"

// Bottom is an option to crop your image
const Bottom CropType = "bottom"

// Bottomright is an option to crop your image
const Bottomright CropType = "bottomright"

// Right is an option to crop your image
const Right CropType = "right"

// Topright is an option to crop your image
const Topright CropType = "topright"

// Center is an option to crop your image
const Center CropType = "center"
