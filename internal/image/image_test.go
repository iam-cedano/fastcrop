package includes

import (
	"testing"
)

func TestImage_Crop(t *testing.T) {
	image := Image{Path: "testdata/test.jpg"}
	err := image.Crop(100, 100)
	if err != nil {
		t.Fatalf("Error cropping image: %v", err)
	}
}
