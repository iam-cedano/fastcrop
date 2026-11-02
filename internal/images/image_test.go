package images

import (
	"image"
	"os"
	"path/filepath"
	"testing"
)

func TestImage_Crop(t *testing.T) {
	src := filepath.Join("..", "..", "testdata", "01.webp")
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatalf("read fixture: %v", err)
	}

	dst := filepath.Join(t.TempDir(), "01.webp")
	if err := os.WriteFile(dst, data, 0o644); err != nil {
		t.Fatalf("copy fixture: %v", err)
	}

	img := Image{Path: dst}
	if err := img.Crop(100, 100); err != nil {
		t.Fatalf("Error cropping image: %v", err)
	}

	file, err := os.Open(dst)
	if err != nil {
		t.Fatalf("open cropped image: %v", err)
	}
	defer file.Close()

	decoded, _, err := image.Decode(file)
	if err != nil {
		t.Fatalf("decode cropped image: %v", err)
	}

	bounds := decoded.Bounds()
	if bounds.Dx() != 100 || bounds.Dy() != 100 {
		t.Fatalf("cropped size = %dx%d, want 100x100", bounds.Dx(), bounds.Dy())
	}
}
