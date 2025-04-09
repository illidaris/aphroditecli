package download

import (
	"net/url"
	"testing"
)

func TestParseUrl(t *testing.T) {
	raw := "https://siteres.ztgame.com/s4/spaceparty/images/activity/2023-10/tA70LWOYBUdUYnQOUF6Xk9OW73L7XJlkd9ZnY9GH.png"
	u, err := url.Parse(raw)
	if err != nil {
		t.Error(err)
	}

	println(u.Path)
}
