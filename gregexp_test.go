package gregexp

import (
	"testing"
)

type testCase map[string]bool

var testSet = map[string]testCase{
	// Wildcard match
	"lib/**/*_test.go": testCase{
		"lib/_test.go":         true,
		"lib/hoge_test.go":     true,
		"lib/sub/dir/_test.go": true,
		"lib/dir/path_test.go": true,
		"lib/dir/path.go":      false,
		"lib/.ignore_test.go":  false,
	},
	// Escaped wildcard
	"lib/\\*_test.go": testCase{
		"lib/*_test.go": true,
		"lib/n_test.go": false,
	},
	"lib/\\*\\*/test.go": testCase{
		"lib/test.go":     false,
		"lib/dir/test.go": false,
	},
	// Anychar
	"lib/?.go": testCase{
		"lib/a.go":   true,
		"lib/.go":    false,
		"lib/abc.go": false,
	},
	// Range
	"lib/[abc].go": testCase{
		"lib/a.go": true,
		"lib/b.go": true,
		"lib/c.go": true,
		"lib/d.go": false,
	},
	// Range not(w/ ^)
	"lib/[^abc].go": testCase{
		"lib/a.go": false,
		"lib/b.go": false,
		"lib/c.go": false,
		"lib/d.go": true,
	},
	// Range not(w/ !)
	"lib/[!abc].go": testCase{
		"lib/a.go": false,
		"lib/b.go": false,
		"lib/c.go": false,
		"lib/d.go": true,
	},
	// Range
	"lib/[a-c0-9].go": testCase{
		"lib/a.go": true,
		"lib/b.go": true,
		"lib/c.go": true,
		"lib/1.go": true,
		"lib/2.go": true,
		"lib/8.go": true,
		"lib/d.go": false,
	},
	// OR
	"lib/{abc,012}.go": testCase{
		"lib/abc.go":       true,
		"lib/012.go":       true,
		"lib/{abc,012}.go": false,
	},
	"lib/{abc}.go": testCase{
		"lib/abc.go":       true,
		"lib/012.go":       false,
		"lib/{abc,012}.go": false,
	},
	// Escaped OR
	"lib/\\{abc\\,012\\}.go": testCase{
		"lib/abc.go":       false,
		"lib/012.go":       false,
		"lib/{abc,012}.go": true,
	},
}

var crashCases = []string{
	"lib/**/{.go",
}

func TestConvert(t *testing.T) {
	t.Parallel()

	for pattern, cases := range testSet {
		func(patter string, cases testCase) {
			t.Run(pattern, func(t *testing.T) {
				for target, match := range cases {
					func(pattern, target string, match bool) {
						t.Run("with "+target, func(t *testing.T) {
							t.Logf("Match: %s <=> %s(%v)", pattern, target, match)
							r, err := Convert(pattern)
							if err != nil {
								t.Fatal(err)
							}
							m := r.MatchString(target)
							if m != match {
								t.Logf("Compiled: %s", r.String())
								t.Errorf("%s <=> %s(expected: %v, actual: %v)", pattern, target, match, m)
							}
						})
					}(pattern, target, match)
				}
			})
		}(pattern, cases)
	}
}

func TestMatch(t *testing.T) {
	t.Parallel()

	for pattern, cases := range testSet {
		func(patter string, cases testCase) {
			t.Run(pattern, func(t *testing.T) {
				for target, match := range cases {
					func(pattern, target string, match bool) {
						t.Run("with "+target, func(t *testing.T) {
							t.Logf("Match: %s <=> %s(%v)", pattern, target, match)
							m, err := Match(pattern, []byte(target))
							if err != nil {
								t.Fatal(err)
							}
							if m != match {
								t.Errorf("%s <=> %s(expected: %v, actual: %v)", pattern, target, match, m)
							}
						})
					}(pattern, target, match)
				}
			})
		}(pattern, cases)
	}
}

func TestMatchString(t *testing.T) {
	t.Parallel()

	for pattern, cases := range testSet {
		func(patter string, cases testCase) {
			t.Run(pattern, func(t *testing.T) {
				for target, match := range cases {
					func(pattern, target string, match bool) {
						t.Run("with "+target, func(t *testing.T) {
							t.Logf("Match: %s <=> %s(%v)", pattern, target, match)
							m, err := MatchString(pattern, target)
							if err != nil {
								t.Fatal(err)
							}
							if m != match {
								t.Errorf("%s <=> %s(expected: %v, actual: %v)", pattern, target, match, m)
							}
						})
					}(pattern, target, match)
				}
			})
		}(pattern, cases)
	}
}

func TestCrash(t *testing.T) {
	var err error
	for _, pattern := range crashCases {
		_, err = MatchString(pattern, "lib/test.go")
		if err == nil {
			t.Fatal("Required crash")
		} else {
			t.Log(pattern + ": " + err.Error())
		}

		_, err = Match(pattern, []byte("lib/test.go"))
		if err == nil {
			t.Fatal("Required crash")
		} else {
			t.Log(pattern + ": " + err.Error())
		}
	}
}
