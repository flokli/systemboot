package bls

import (
	"bufio"
	"fmt"
	"regexp"
	"strings"
)

type KVMap map[string][]string

var rWhitespace = regexp.MustCompile("\\s+")
var rLine = regexp.MustCompile(`^(?P<k>[\S]+)\s+(?P<v>[\S]+)\s*$`)

func Parse(scanner *bufio.Scanner) (results KVMap, err error) {
	if scanner == nil {
		return nil, fmt.Errorf("scanner may not be nil")
	}
	results = make(KVMap)

	// TODO: regexp as top-level constant, use find to trim and split instead
	for scanner.Scan() {
		line := strings.TrimSpace(scanner.Text())
		if strings.HasPrefix(line, "#") {
			continue
		}
		match := rLine.FindStringSubmatch(line)
		k := match[1]
		v := match[2]

		results[k] = append(results[k], v)
	}
	return results, scanner.Err()
}
