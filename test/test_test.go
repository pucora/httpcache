package test_test

import (
	"testing"

	"github.com/velonetics/httpcache"
	"github.com/velonetics/httpcache/test"
)

func TestMemoryCache(t *testing.T) {
	test.Cache(t, httpcache.NewMemoryCache())
}
