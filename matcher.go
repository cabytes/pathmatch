package pathmatch

import (
	"errors"
	"regexp"
	"strings"
)

var (
	RegMatchVar = regexp.MustCompile(`{(?P<var>[a-z]*)(?:\:(?P<reg>.*))?}`)
)

type Match interface {
	Var(name string) string
	Has(name string) bool
}

type Matcher interface {
	Match(url string) (Match, error)
}

type matcher struct {
	pattern string
}

type match struct {
	pattern string
	url     string
	vars    map[string]string
}

var ErrNotMatched = errors.New("not matched")

func NewMatcher(pattern string) Matcher {
	return &matcher{pattern: pattern}
}

func (matcher *matcher) Match(url string) (Match, error) {

	if matcher.pattern == url {
		return &match{
			pattern: matcher.pattern,
			url:     url,
			vars:    make(map[string]string),
		}, nil
	}

	parts := strings.Split(matcher.pattern, "/")[1:]

	mappings := strings.Split(url, "/")[1:]

	var m match = match{
		pattern: matcher.pattern,
		url:     url,
		vars:    make(map[string]string),
	}

	for pos, part := range parts {

		if RegMatchVar.Match([]byte(part)) {

			matches := RegMatchVar.FindStringSubmatch(part)

			if matches[2] != "" {

				if false == regexp.MustCompile("^"+matches[2]+"$").Match([]byte(mappings[pos])) {
					return nil, errors.New("Variable does not match regexp")
				}

				m.vars[matches[1]] = mappings[pos]

				continue

			} else {
				m.vars[matches[1]] = mappings[pos]
				continue
			}
		} else if len(mappings) > pos && part == mappings[pos] {
			continue
		}

		if part == "*" && len(parts) < len(mappings) {
			return &m, nil
		}

		if part == "*" {
			continue
		}

		if len(mappings) < pos && part == mappings[pos] {
			continue
		}

		return nil, ErrNotMatched
	}

	if len(mappings) > len(parts) {
		return nil, ErrNotMatched
	}

	return &m, nil
}

func (m *match) Has(name string) bool {
	_, has := m.vars[name]
	return has
}

func (m *match) Var(name string) string {
	v, _ := m.vars[name]
	return v
}
