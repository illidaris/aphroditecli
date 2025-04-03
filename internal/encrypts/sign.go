package encrypts

import (
	"crypto/hmac"
	"crypto/md5"
	"crypto/sha1"
	"crypto/sha256"
	"fmt"
	"hash"

	"github.com/illidaris/aphroditecli/pkg/exporter"
)

func Sign(secret string, args ...string) [][]string {
	rows := [][]string{}
	if secret == "" {
		row := []string{"Raw", "MD5", "SHA1", "SHA256"}
		rows = append(rows, row)
	} else {
		row := []string{"Raw", "hmac-MD5", "hmac-SHA1", "hmac-SHA256"}
		rows = append(rows, row)
	}
	for _, v := range args {
		row := []string{
			v,
			HashSign(md5.New, secret, v),
			HashSign(sha1.New, secret, v),
			HashSign(sha256.New, secret, v),
		}
		rows = append(rows, row)
	}
	exporter.FmtTable(rows, false)
	return rows
}

func HashSign(f func() hash.Hash, secret string, raw string) string {
	var hashSign hash.Hash
	if secret == "" {
		hashSign = f()
	} else {
		hashSign = hmac.New(f, []byte(secret))
	}
	hashSign.Write([]byte(raw))
	bs := hashSign.Sum(nil)
	return fmt.Sprintf("%v\n%v", fmt.Sprintf("%X", bs), fmt.Sprintf("%x", bs))
}
