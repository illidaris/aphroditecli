package download

import (
	"fmt"
	"io"
	"net/http"
	"net/url"
	"os"
	"path"
	"strings"

	"github.com/illidaris/aphrodite/pkg/group"
	fileex "github.com/illidaris/file/path"
	"github.com/schollz/progressbar/v3"
)

func Download(out string, needdir bool, args ...string) {
	urls := []string{}
	for _, v := range args {
		raw := UrlsFrmFile(v)
		if len(raw) == 0 {
			continue
		}
		for _, url := range raw {
			if url != "" {
				urls = append(urls, url)
			}
		}
	}
	bar := progressbar.Default(int64(len(urls)))
	_, _ = group.GroupBaseFunc(func(vs ...string) (int64, error) {
		for _, v := range vs {
			defer bar.Add(1)
			err := SaveAsFile(out, needdir)(v)
			if err != nil {
				fmt.Printf("%v 保存错误: %v\n", v, err)
			}
		}
		return 1, nil
	}, 1, urls...)
}

func SaveAsFile(out string, needdir bool) func(urlstr string) error {
	return func(urlstr string) error {
		keys, name := parseKeys(urlstr)
		fDirKeys := []string{out}
		if needdir {
			fDirKeys = append(fDirKeys, keys...)
		}
		fDirName := path.Join(fDirKeys...)
		fileName := path.Join(fDirName, name)
		_ = fileex.MkdirIfNotExist(fDirName)
		resp, err := http.Get(urlstr)
		if err != nil {
			return err
		}
		defer resp.Body.Close()
		f, err := os.Create(fileName)
		if err != nil {
			return err
		}
		_, err = io.Copy(f, resp.Body)
		if err != nil {
			return err
		}
		return nil
	}
}

func parseKeys(urlstr string) ([]string, string) {
	u, err := url.Parse(urlstr)
	if err != nil {
		return nil, ""
	}
	keys := strings.Split(u.Path, "/")
	if len(keys) == 0 {
		return keys, ""
	}
	return keys[:len(keys)-1], keys[len(keys)-1]
}

func UrlsFrmFile(file string) []string {
	if parsedURL, err := url.Parse(file); err != nil {
		return []string{file}
	} else if parsedURL.Scheme != "http" && parsedURL.Scheme != "https" {
		return []string{file}
	}
	_, err := os.Stat(file)
	if os.IsNotExist(err) {
		return nil
	}
	bs, err := os.ReadFile(file)
	if err != nil {
		return nil
	}
	content := string(bs)
	content = strings.ReplaceAll(content, "\r\n", "\n")
	return strings.Split(content, "\n")
}
