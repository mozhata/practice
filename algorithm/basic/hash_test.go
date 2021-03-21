package basic_test

import (
	"practice/algorithm/basic"
	"testing"

	"github.com/smartystreets/assertions"
)

func TestHashSetHas(t *testing.T) {
	h := basic.NewHashSet()
	Announce(t, h.Contains("key"), assertions.ShouldEqual, false)
	h.Add("key")
	Announce(t, h.Contains("key"), assertions.ShouldEqual, true)
	h.Remove("key")
	Announce(t, h.Contains("key"), assertions.ShouldEqual, false)
}
