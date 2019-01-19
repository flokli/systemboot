package bls

import (
	"github.com/stretchr/testify/assert"
	"testing"
)

type magic struct {
	Foo []string `key:"foo"`
	Baz string   `key:"baz"`
}

func TestMagic(t *testing.T) {
	result := new(magic)
	kvMap := KVMap{"foo": []string{"bar"}, "baz": []string{"blub"}}
	err := Deserialize(kvMap, result)
	assert.NoError(t, err)
	assert.Equal(t, "blub", result.Baz)
	assert.Equal(t, []string{"bar"}, result.Foo)
	// TODO: this should probably contain more tests which exercise the other branches in Deserialize, too.
}

func TestMagicWithIncompatibleConfig(t *testing.T) {
	result := new(magic)
	kvMap := KVMap{"baz": []string{"blub", "I try to be a list, boo!"}}
	err := Deserialize(kvMap, result)
	assert.Error(t, err)
}
