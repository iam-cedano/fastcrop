package includes

import (
	"path/filepath"

	"github.com/disintegration/imaging"
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

	return imaging.Save(img, path)
}
