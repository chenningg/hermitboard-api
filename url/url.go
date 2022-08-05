package url

import gourl "net/url"

func Encode(str string) string {
	return gourl.PathEscape(str)
}
