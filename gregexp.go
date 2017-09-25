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

// Match checks whether a path pattern matches a byte slice.
//
func Match(pattern string, target []byte) (bool, error) {
	r, err := Convert(pattern)
	if err != nil {
		return false, err
	}
	return r.Match(target), nil
}

// Match checks whether a path pattern matches a string.
//
func MatchString(pattern, target string) (bool, error) {
	r, err := Convert(pattern)
	if err != nil {
		return false, err
	}
	return r.MatchString(target), nil
}

// Compile parses a path pattern and returns a regexp object, if successful.
//
func Convert(pattern string) (*regexp.Regexp, error) {
	pattern = "^" + metaReplacer.Replace(regexp.QuoteMeta(pattern)) + "$"
	return regexp.Compile(pattern)
}
