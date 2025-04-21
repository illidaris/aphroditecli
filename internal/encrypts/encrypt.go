package encrypts

import (
	"bytes"
	"encoding/base64"

	"github.com/illidaris/aphrodite/pkg/encrypter"
	"github.com/illidaris/aphroditecli/pkg/exporter"
)

func Encrypt(secret string, args ...string) [][]string {
	rows := [][]string{
		{"Raw", "Encrypt"},
	}
	for _, v := range args {
		raw := v
		result := encryptString(secret, raw)
		row := []string{
			v,
			result,
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}

func Decrypt(secret string, args ...string) [][]string {
	rows := [][]string{
		{"Raw", "Decrypt"},
	}
	for _, v := range args {
		result := decryptString(secret, v)
		row := []string{
			v,
			result,
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}

func encryptString(secret string, raw string) string {
	if secret == "" || raw == "" {
		return ""
	}
	bs := []byte{}
	w := bytes.NewBuffer(bs)
	err := encrypter.EncryptStream(bytes.NewBufferString(raw), w, []byte(secret))
	if err != nil {
		return err.Error()
	}
	return base64.StdEncoding.EncodeToString(w.Bytes())
}

func decryptString(secret string, raw string) string {
	if secret == "" || raw == "" {
		return ""
	}
	pwdBs := []byte{}
	w := bytes.NewBuffer(pwdBs)
	rawBs, err := base64.StdEncoding.DecodeString(raw)
	if err != nil {
		return err.Error()
	}
	err = encrypter.DecryptStream(bytes.NewBuffer(rawBs), w, []byte(secret))
	if err != nil {
		return err.Error()
	}
	return w.String()
}
