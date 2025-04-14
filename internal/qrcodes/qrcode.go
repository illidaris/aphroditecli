package qrcodes

import (
	"bytes"
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"

	"github.com/illidaris/aphroditecli/pkg/exporter"
	qrcode "github.com/skip2/go-qrcode"
	qrcodeReader "github.com/tuotoo/qrcode"
)

func WriteQrCodeExport(size int, dest string, contents ...string) {
	rows := [][]string{
		{"Content", "Path", "Data", "Error"},
	}
	for index, content := range contents {
		destFull := path.Join(dest, fmt.Sprintf("%d.png", index+1))
		bs, err := WriteQrCode(content, qrcode.Medium, size, destFull)
		rows = append(rows, []string{content, destFull, string(bs), fmt.Sprintf("%v", err)})
	}
	exporter.FmtTable(rows, false)
}

func WriteQrCode(content string, quality qrcode.RecoveryLevel, size int, dest string) ([]byte, error) {
	if dest != "" {
		return nil, WriteQrCodeToFile(content, qrcode.Medium, size, dest)
	}
	return qrcode.Encode(content, quality, size)
}

func ParseQrCodeExport(raws ...string) {
	rows := [][]string{
		{"Raw", "Content", "Error"},
	}
	for _, raw := range raws {
		res, err := ParseQrCode(raw)
		rows = append(rows, []string{raw, res, fmt.Sprintf("%v", err)})
	}
	exporter.FmtTable(rows, false)
}
func ParseQrCode(raw string) (string, error) {
	if _, err := url.Parse(raw); err != nil {
		return ReadQRCodeByDisk(raw)
	}
	return ReadQRCodeByUrl(raw)
}

func ReadQRCodeByUrl(fileurl string) (string, error) {
	resp, err := http.Get(fileurl)
	if err != nil {
		return "", err
	}
	defer resp.Body.Close()
	return ReadQRCodeByReader(resp.Body)
}

func ReadQRCodeByDisk(filename string) (string, error) {
	bs, err := os.ReadFile(filename)
	if err != nil {
		return "", err
	}
	return ReadQRCodeByReader(bytes.NewReader(bs))
}

func ReadQRCodeByReader(reader io.Reader) (string, error) {
	qrmatrix, err := qrcodeReader.Decode(reader)
	if err != nil {
		return "", err
	}
	if qrmatrix == nil {
		return "", nil
	}
	return qrmatrix.Content, nil
}

func WriteQrCodeToFile(content string, quality qrcode.RecoveryLevel, size int, dest string) error {
	return qrcode.WriteFile(content, quality, size, dest)
}
