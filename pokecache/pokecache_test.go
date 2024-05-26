package pokecache

import (
	"testing"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache()
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddCacheEntry(t *testing.T) {
	cache := NewCache()

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{inputKey: "key1", inputVal: []byte("val1")},
	}

	for _, cse := range cases {

		cache.Add(cse.inputKey, cse.inputVal)
		actual, ok := cache.Get(cse.inputKey)
		if !ok {
			t.Errorf("%v not found", cse.inputKey)
			continue
		}
		if string(actual) != string(cse.inputVal) {
			t.Errorf("expected value not found, %v vs %s", string(cse.inputVal), string(actual))
		}
	}
}
