package pokecache

import (
	"testing"
	"time"
)

func TestCreateCache(t *testing.T) {
	cache := NewCache(time.Millisecond)
	if cache.cache == nil {
		t.Error("cache is nil")
	}
}

func TestAddCacheEntry(t *testing.T) {
	cache := NewCache(time.Millisecond)

	cases := []struct {
		inputKey string
		inputVal []byte
	}{
		{inputKey: "key1", inputVal: []byte("val1")},
		{inputKey: "key2", inputVal: []byte("val2")},
		{inputKey: "", inputVal: []byte("val3")},
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

func TestReapCache(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "key1"
	cache.Add(key, []byte("val1"))

	time.Sleep(interval + time.Millisecond)

	_, ok := cache.Get(key)
	if ok {
		t.Errorf("%s should have been reaped", key)
	}
}

func TestReapCacheFail(t *testing.T) {
	interval := time.Millisecond * 10
	cache := NewCache(interval)

	key := "key1"
	cache.Add(key, []byte("val1"))

	time.Sleep(interval / 2)

	_, ok := cache.Get(key)
	if !ok {
		t.Errorf("%s should have been reaped", key)
	}
}
