package gregexp

import (
	"os"
	"regexp"
	"strings"
)

var pathSeparator = regexp.QuoteMeta(string(os.PathSeparator))

var metaReplacer = strings.NewReplacer(
	// Wildcard
	"\\\\\\*\\\\\\*"+pathSeparator, "\\*\\*"+pathSeparator,
	"\\\\\\*", "\\*",
	"\\*\\*"+pathSeparator, "(?:[^"+pathSeparator+"]+"+pathSeparator+")*",
	"\\*", "(?:[^."+pathSeparator+"][^"+pathSeparator+"]*)?",
	// Anychar
	"\\\\\\?", "\\?",
	"\\?", "[^"+pathSeparator+"]",
	// Range
	"\\\\\\[", "\\[",
	"\\\\\\]", "\\]",
	"\\[\\^", "[^",
	"\\[!", "[^",
	"\\[", "[",
	"\\]", "]",
	// OR
	"\\\\\\{", "\\{",
	"\\\\\\}", "\\}",
	"\\{", "(?:",
	"\\}", ")",
	"\\,", ",",
	",", "|",
)

func Match(pattern string, target []byte) (bool, error) {
	r, err := Convert(pattern)
	if err != nil {
		return false, err
	}
	return r.Match(target), nil
}

func MatchString(pattern, target string) (bool, error) {
	r, err := Convert(pattern)
	if err != nil {
		return false, err
	}
	return r.MatchString(target), nil
}

func Convert(pattern string) (*regexp.Regexp, error) {
	pattern = "^" + metaReplacer.Replace(regexp.QuoteMeta(pattern)) + "$"
	return regexp.Compile(pattern)
}
