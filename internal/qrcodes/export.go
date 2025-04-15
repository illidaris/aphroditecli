package qrcodes

import (
	"bytes"
	"encoding/base64"
	"fmt"
	"os"
	"path"

	apqrcodes "github.com/illidaris/aphrodite/pkg/qrcodes"
	"github.com/illidaris/aphroditecli/pkg/exporter"
	qrcode "github.com/skip2/go-qrcode"
)

func WriteQrCodeExport(size, logoP int, logo, out string, contents ...string) {
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
		bs, err := WriteQrCode(content, qrcode.Medium, size, logoP, logoBs, destFull)
		rows = append(rows, []string{content, destFull, base64.StdEncoding.EncodeToString(bs), fmt.Sprintf("%v", err)})
	}
	exporter.FmtTable(rows, false)
}

func WriteQrCode(content string, quality qrcode.RecoveryLevel, size, logoP int, logo []byte, dest string) ([]byte, error) {
	// 当指定目标路径时，直接写入文件并返回nil字节数组
	if dest == "" {
		return qrcode.Encode(content, quality, size)
	}
	if len(logo) == 0 || logoP == 0 {
		return nil, qrcode.WriteFile(content, qrcode.Medium, size, dest)
	}
	// 未指定路径时，生成PNG格式字节数组
	bs, err := qrcode.Encode(content, quality, size)
	if err != nil {
		return bs, err
	}
	outFile, err := os.Create(dest)
	if err != nil {
		return bs, err
	}
	defer outFile.Close()
	return bs, apqrcodes.ImageWithLogo(bytes.NewBuffer(bs), bytes.NewBuffer(logo), logoP, outFile)
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
