package images

import (
	"image"
	"os"
	"path/filepath"
	"strings"

	"github.com/HugoSmits86/nativewebp"
	"github.com/disintegration/imaging"
	_ "golang.org/x/image/webp"
)

type Image struct {
	Path string
}

func (i *Image) Crop(width, height int) error {
	path, err := filepath.Abs(i.Path)
	if err != nil {
		return err
	}

	img, err := imaging.Open(path)
	if err != nil {
		return err
	}

	img = imaging.Resize(img, width, height, imaging.Lanczos)

	if strings.EqualFold(filepath.Ext(path), ".webp") {
		return saveWebP(img, path)
	}

	return imaging.Save(img, path)
}

func saveWebP(img image.Image, path string) error {
	file, err := os.Create(path)
	if err != nil {
		return err
	}
	defer file.Close()

	return nativewebp.Encode(file, img, nil)
}
