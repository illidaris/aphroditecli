package qrcodes

import (
	"bytes"
	"encoding/base64"
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

func WriteQrCodeExport(size int, out string, contents ...string) {
	rows := [][]string{
		{"Content", "Path", "Data", "Error"},
	}
	for index, content := range contents {
		var destFull string
		if out != "" {
			destFull = path.Join(out, fmt.Sprintf("%d.png", index+1))
		}
		res, err := WriteQrCode(content, qrcode.Medium, size, destFull)
		rows = append(rows, []string{content, destFull, res, fmt.Sprintf("%v", err)})
	}
	exporter.FmtTable(rows, false)
}

func WriteQrCode(content string, quality qrcode.RecoveryLevel, size int, dest string) (string, error) {
	if dest != "" {
		return "", WriteQrCodeToFile(content, qrcode.Medium, size, dest)
	}
	bs, err := qrcode.Encode(content, quality, size)
	if err != nil {
		return "", err
	}
	return base64.URLEncoding.EncodeToString(bs), nil
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
		if res, bs64Err := ReadQRCodeByBase64(raw); bs64Err == nil {
			return res, nil
		}
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

func ReadQRCodeByBase64(bs64 string) (string, error) {
	bs, err := base64.URLEncoding.DecodeString(bs64)
	if err != nil {
		bs, err = base64.StdEncoding.DecodeString(bs64)
		if err != nil {
			return "", err
		}
	}
	return ReadQRCodeByReader(bytes.NewReader(bs))
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
