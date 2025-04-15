package qrcodes

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"image"
	"image/color"
	"image/draw"
	"image/png"
	"io"
	"os"
	"path"

	apqrcodes "github.com/illidaris/aphrodite/pkg/qrcodes"
	"github.com/illidaris/aphroditecli/pkg/exporter"
	"github.com/nfnt/resize"
	qrcode "github.com/skip2/go-qrcode"
)

func WriteQrCodeExport(size, logoP, zoom int, logo, out string, contents ...string) {
	rows := [][]string{
		{"Content", "Path", "Data", "Error"},
	}
	var (
		logoBs []byte
	)
	if logo != "" {
		logoBs, _ = apqrcodes.ReadFile(logo)
	}
	for index, content := range contents {
		var destFull string
		if out != "" {
			destFull = path.Join(out, fmt.Sprintf("%d.png", index+1))
		}
		bs, err := WriteQrCode(content, qrcode.Medium, size, logoP, zoom, logoBs, destFull)
		rows = append(rows, []string{content, destFull, base64.StdEncoding.EncodeToString(bs), fmt.Sprintf("%v", err)})
	}
	exporter.FmtTable(rows, false)
}

func WriteQrCode(content string, quality qrcode.RecoveryLevel, size, logoP, zoom int, logo []byte, dest string) ([]byte, error) {
	// 当指定目标路径时，直接写入文件并返回nil字节数组
	if dest == "" {
		return qrcode.Encode(content, quality, size)
	}
	if len(logo) == 0 || logoP == 0 {
		return nil, qrcode.WriteFile(content, qrcode.Medium, size, dest)
	}
	// 未指定路径时，生成PNG格式字节数组
	q, err := qrcode.New(content, quality)
	if err != nil {
		return nil, err
	}
	q.DisableBorder = true
	bs, err := q.PNG(size)
	if err != nil {
		return bs, err
	}
	outFile, err := os.Create(dest)
	if err != nil {
		return bs, err
	}
	defer outFile.Close()
	return bs, ImageWithLogo(bytes.NewBuffer(bs), bytes.NewBuffer(logo), logoP, zoom, outFile)
}

func ImageWithLogo(src, logo io.Reader, logoP, zoom int, out io.Writer) error {
	if logoP > 10 || logoP < 0 {
		return fmt.Errorf("logo占比必须在0-10之间")
	}
	// 将二维码文件接码成图片 370 20  370 20
	srcImg, srcName, srcErr := image.Decode(src)
	if srcErr != nil {
		return srcErr
	}
	if srcImg == nil {
		return fmt.Errorf("source %v %v", srcName, srcErr)
	}
	// 调整原图大小（二维码大小的1/18.5）
	srcImgSize := srcImg.Bounds().Dx() / 10 * zoom
	resizedSrc := resize.Resize(uint(srcImgSize), 0, srcImg, resize.Lanczos3)
	// 计算Logo位置（居中）
	srcOffset := image.Pt(
		(srcImg.Bounds().Dx()-resizedSrc.Bounds().Dx())/2,
		(srcImg.Bounds().Dy()-resizedSrc.Bounds().Dy())/2,
	)
	// 将填充图解码成png图片
	logoImg, logoName, logoErr := image.Decode(logo)
	if logoErr != nil {
		return logoErr
	}
	if logoImg == nil {
		return fmt.Errorf("logo %v %v", logoName, srcErr)
	}
	// 调整Logo大小（二维码大小的1/5）
	logoSize := srcImg.Bounds().Dx() / logoP
	resizedLogo := resize.Resize(uint(logoSize), 0, logoImg, resize.Lanczos3)
	// 计算Logo位置（居中）
	offset := image.Pt(
		(srcImg.Bounds().Dx()-resizedLogo.Bounds().Dx())/2,
		(srcImg.Bounds().Dy()-resizedLogo.Bounds().Dy())/2,
	)
	// 创建画布并合并图片
	canvas := image.NewRGBA(srcImg.Bounds())
	white := image.NewUniform(color.White) // 使用draw包中的Draw方法快速填充
	draw.Draw(canvas, srcImg.Bounds(), white, image.Point{}, draw.Src)
	draw.Draw(canvas, resizedSrc.Bounds().Add(srcOffset), resizedSrc, image.Point{}, draw.Over)
	draw.Draw(canvas, resizedLogo.Bounds().Add(offset), resizedLogo, image.Point{}, draw.Over)
	return png.Encode(out, canvas)
}

func ParseQrCodeExport(raws ...string) {
	rows := [][]string{
		{"Raw", "Content", "Error"},
	}
	for _, raw := range raws {
		res, err := apqrcodes.ParseQrCode(raw)
		rows = append(rows, []string{raw, res, fmt.Sprintf("%v", err)})
	}
	exporter.FmtTable(rows, false)
}
