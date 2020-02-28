package models

// Image Data transfer object for an image.
type Image struct {

	// The captured image data in PNG format, encoded as a base64 string. The data string shall not exceed 10MB.
	ImageData string
}
