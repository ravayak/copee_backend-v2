package url

import "regexp"

func GetResourceName(url string) string {
	re := regexp.MustCompile(`^/api/v1/(?P<resource>[^/]+)/?`)
	match := re.FindStringSubmatch(url)

	if len(match) > 0 {
		return match[1]
	} else {
		return ""
	}
}
