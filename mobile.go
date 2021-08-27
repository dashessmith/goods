package goods

import "regexp"

var reg *regexp.Regexp

func init() {
	var err error
	reg, err = regexp.Compile(`(\+86)?1[3-9][0-9]{9}`)
	if err != nil {
		panic(err)
	}
}

func IsCnMobile(no string) bool {
	no = TrimAllSpace(no)
	return reg.MatchString(no)
}
