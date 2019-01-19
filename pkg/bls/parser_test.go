package bls

import (
	"bufio"
	"github.com/stretchr/testify/assert"
	"strings"
	"testing"
)

func str2Scanner(str string) *bufio.Scanner {
	return bufio.NewScanner(strings.NewReader(str))
}
func parseAndAssertResult(t *testing.T, str string, expected KVMap) {
	scanner := str2Scanner(str)
	results, err := Parse(scanner)
	assert.NoError(t, err)
	assert.Equal(t, expected, results)
}

func TestParserSimple(t *testing.T) {
	parseAndAssertResult(t, "foo bar", KVMap{"foo": []string{"bar"}})
}
func TestParserIgnoreComment(t *testing.T) {
	parseAndAssertResult(t, "foo bar\n#bla", KVMap{"foo": []string{"bar"}})
}
func TestParserMultiline(t *testing.T) {
	parseAndAssertResult(t, "foo bar\nbaz baf", KVMap{"foo": []string{"bar"}, "baz": []string{"baf"}})
}
func TestParserWhitespaceAtEnd(t *testing.T) {
	parseAndAssertResult(t, "foo bar\nfoo baz    ", KVMap{"foo": []string{"bar", "baz"}})
}
func TestParserMultilineReverse(t *testing.T) {
	parseAndAssertResult(t, "baz baf\nfoo bar", KVMap{"foo": []string{"bar"}, "baz": []string{"baf"}})
}
func TestParserKeysTwice(t *testing.T) {
	parseAndAssertResult(t, "foo bar\nfoo baz", KVMap{"foo": []string{"bar", "baz"}})
}
func TestParserMoreWhitespace(t *testing.T) {
	parseAndAssertResult(t, "foo bar\nfoo  \tbaz", KVMap{"foo": []string{"bar", "baz"}})
}

func TestParserNotNil(t *testing.T) {
	_, err := Parse(nil)
	assert.Error(t, err)
}
