package images

import (
	"fmt"
	"image"
	"image/png"
	"io/fs"
	"os"
	"path"
	"path/filepath"

	"github.com/nfnt/resize"
)

func Resize(src, target string, width uint, height uint) error {
	if target == "" {
		target = path.Join(src, "thb")
	}
	return filepath.WalkDir(src, func(ph string, d fs.DirEntry, err error) error {
		if err != nil {
			fmt.Printf("访问路径 %q 时出错: %v\n", ph, err)
			return err
		}
		if d.IsDir() {
			return nil
		}
		srcFile := path.Join(src, d.Name())
		targetFile := path.Join(target, d.Name())
		return resized(srcFile, targetFile, width, height)
	})
}

func resized(src, target string, width uint, height uint) error {
	raw, err := os.Open(src)
	if err != nil {
		return err
	}
	defer raw.Close()
	srcImg, _, err := image.Decode(raw)
	if err != nil {
		return err
	}
	var targetImg image.Image
	if srcImg.Bounds().Dx() < int(width) {
		targetImg = srcImg
	} else {
		targetImg = resize.Resize(width, height, srcImg, resize.Lanczos3)
	}
	outFile, err := os.Create(target)
	if err != nil {
		return err
	}
	defer outFile.Close()
	return png.Encode(outFile, targetImg)
}
