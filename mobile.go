package goods

import "regexp"

var reg *regexp.Regexp = regexp.MustCompile(`(\+86)?1[3-9][0-9]{9}`)

func IsCnMobile(no string) bool {
	no = TrimAllSpace(no)
	return reg.MatchString(no)
}
