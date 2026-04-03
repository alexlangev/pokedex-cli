package pokecache

import (
	"bytes"
	"testing"
	"time"
)

func TestCacheAdd(t *testing.T) {
	cases := []struct {
		name     string
		key      string
		value    []byte
		expected []byte
	}{
		{
			name:     "empty",
			key:      "empty",
			value:    []byte{},
			expected: []byte{},
		},
		{
			name:     "simple",
			key:      "foo",
			value:    []byte("bar"),
			expected: []byte("bar"),
		},
		{
			name:     "overwrite existing entry",
			key:      "foo",
			value:    []byte("different bar"),
			expected: []byte("different bar"),
		},
	}

	testCache := NewCache(10 * time.Second)

	for _, c := range cases {
		testCache.Add(c.key, c.value)

		if !bytes.Equal(c.expected, c.value) {
			t.Errorf("[%s] expected val %v, got %v", c.name, c.expected, c.value)
		}
	}
}
