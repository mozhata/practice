package basic_test

import (
	"fmt"
	"practice/algorithm/basic"
	"runtime"
	"testing"

	"github.com/smartystreets/assertions"
)

func fullStackTrace() string {
	buffer := make([]byte, 1024*64)
	n := runtime.Stack(buffer, true)
	return string(buffer[:n])
}

type assertion func(actual interface{}, expected ...interface{}) string

func Announce(t *testing.T, actual interface{}, assert assertion, expect interface{}) {
	msg := assert(actual, expect)
	if msg != "" {
		t.Fatal(fmt.Sprintf("%s\n%s", msg, fullStackTrace()))
		t.SkipNow()
	}
}

func TestSetLen(t *testing.T) {
	s := basic.IntSet{}
	Announce(t, s.Len(), assertions.ShouldEqual, 0)
	s.Add(0)
	Announce(t, s.Len(), assertions.ShouldEqual, 1)
	s.Add(2)
	Announce(t, s.Len(), assertions.ShouldEqual, 2)
	s.Add(3)
	Announce(t, s.Len(), assertions.ShouldEqual, 3)
	s.AddAll(4, 5)
	Announce(t, s.Len(), assertions.ShouldEqual, 5)

}

func TestSetAddHas(t *testing.T) {
	s := &basic.IntSet{}
	s.Add(1)
	Announce(t, s.Len(), assertions.ShouldEqual, 1)
	s.AddAll(2, 2, 3)
	Announce(t, s.Len(), assertions.ShouldEqual, 3)
	Announce(t, s.Has(1), assertions.ShouldEqual, true)
	Announce(t, s.Has(2), assertions.ShouldEqual, true)
	Announce(t, s.Has(3), assertions.ShouldEqual, true)
	Announce(t, s.Has(4), assertions.ShouldEqual, false)
}

func TestSetRemove(t *testing.T) {
	s := &basic.IntSet{}
	s.AddAll(1, 2, 3)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 3})
	s.Remove(1)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{2, 3})
	s.Remove(1)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{2, 3})
}
func TestSetElems(t *testing.T) {
	s := &basic.IntSet{}
	Announce(t, len(s.Elems()), assertions.ShouldResemble, 0)
	s.Add(0)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{0})
	s.AddAll(1, 2, 3)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{0, 1, 2, 3})
}
func TestSetString(t *testing.T) {
	s := &basic.IntSet{}
	Announce(t, s.String(), assertions.ShouldEqual, "{}")
	s.Add(0)
	Announce(t, s.String(), assertions.ShouldEqual, "{0}")
	s.AddAll(1, 2, 3)
	Announce(t, s.String(), assertions.ShouldEqual, "{0 1 2 3}")
}

func TestSetClear(t *testing.T) {
	s := &basic.IntSet{}
	Announce(t, len(s.Elems()), assertions.ShouldResemble, 0)
	Announce(t, s.Len(), assertions.ShouldResemble, 0)
	s.Add(0)
	Announce(t, len(s.Elems()), assertions.ShouldResemble, 1)
	Announce(t, s.Len(), assertions.ShouldResemble, 1)
	s.Clear()
	Announce(t, len(s.Elems()), assertions.ShouldResemble, 0)
	Announce(t, s.Len(), assertions.ShouldResemble, 0)
}
func TestSetCopy(t *testing.T) {
	s := &basic.IntSet{}
	s.AddAll(1, 2, 3)
	s2 := s.Copy()
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 3})
	Announce(t, s2.Elems(), assertions.ShouldResemble, []uint64{1, 2, 3})
	s2.Remove(2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 3})
	Announce(t, s2.Elems(), assertions.ShouldResemble, []uint64{1, 3})
}

func TestSetUnionWith(t *testing.T) {
	s, s2 := &basic.IntSet{}, &basic.IntSet{}
	s.AddAll(1, 2, 3)
	s2.AddAll(3, 4)
	s.UnionWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 3, 4})
	s.Clear()
	s2.Clear()
	s.AddAll(1, 2, 3)
	s.UnionWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 3})
}
func TestSetIntersectionWith(t *testing.T) {
	s, s2 := &basic.IntSet{}, &basic.IntSet{}
	s.AddAll(1, 2, 3)
	s2.AddAll(3, 4)
	s.IntersectionWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{3})
	s.Clear()
	s2.Clear()
	s.AddAll(1, 2, 3, 66)
	s2.AddAll(3, 4)
	s.IntersectionWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{3})
	s.IntersectionWith(basic.IntSet{})
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64(nil))
}
func TestSetDifferenceWith(t *testing.T) {
	s, s2 := &basic.IntSet{}, &basic.IntSet{}
	s.AddAll(1, 2, 3)
	s2.AddAll(3, 4)
	s.DifferenceWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2})
	s.Clear()
	s2.Clear()
	s.AddAll(1, 2, 3, 66)
	s2.AddAll(3, 4)
	s.DifferenceWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 66})
	s.DifferenceWith(basic.IntSet{})
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 66})
}
func TestSetSymmetricDiffWith(t *testing.T) {
	s, s2 := &basic.IntSet{}, &basic.IntSet{}
	s.AddAll(1, 2, 3)
	s2.AddAll(3, 4)
	s.SymmetricDiffWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 4})
	s.Clear()
	s2.Clear()
	s.AddAll(1, 2, 3, 66)
	s2.AddAll(3, 4)
	s.SymmetricDiffWith(*s2)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 4, 66})
	s.SymmetricDiffWith(basic.IntSet{})
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 4, 66})
	// 集合运算
	s.Clear()
	s2.Clear()
	s.AddAll(1, 2, 3)
	s3 := s.Copy()
	s2.AddAll(3, 4)
	s.UnionWith(*s2)
	s3.IntersectionWith(*s2)
	s.DifferenceWith(*s3)
	Announce(t, s.Elems(), assertions.ShouldResemble, []uint64{1, 2, 4})
}
