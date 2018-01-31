package tagpairs

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

var __tagpairRegex *regexp.Regexp

func init() {
	var err error
	__tagpairRegex, err = regexp.Compile(`\[(.+?) "(.+?)"\]`)
	if err != nil {
		panic(err)
	}
}

func Parse(content string) (*Tagpairs, error) {
	tagpairs := NewTagpairs()

	s := bufio.NewScanner(strings.NewReader(content))
	s.Split(bufio.ScanLines)

	for s.Scan() {
		line := s.Text()
		matches := __tagpairRegex.FindStringSubmatch(line)
		if len(matches) < 3 {
			return nil, fmt.Errorf("Error parsing tagpair: %s (%v)", line, matches)
		}

		key := matches[1]
		value := matches[2]

		tagpairs.Data[key] = value
	}

	return tagpairs, nil
}
