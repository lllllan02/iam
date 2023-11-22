package stringutil

import (
	"sort"
	"testing"

	qt "github.com/frankban/quicktest"
)

func TestUnique(t *testing.T) {
	c := qt.New(t)
	c.Helper()

	want := []string{"1", "2", "3", "4"}
	got := Unique(append(want, want...))
	sort.Strings(got)

	c.Assert(got, qt.DeepEquals, want)
}

func TestFindString(t *testing.T) {
	c := qt.New(t)
	c.Helper()

	target := []string{"1", "2", "3", "4"}

	for i := range target {
		c.Assert(FindString(target, target[i]), qt.Equals, i)
	}
	c.Assert(FindString(target, "99"), qt.Equals, -1)
}

func TestStringIn(t *testing.T) {
	c := qt.New(t)
	c.Helper()

	target := []string{"1", "2", "3", "4"}

	for i := range target {
		c.Assert(StringIn(target, target[i]), qt.Equals, true)
	}
	c.Assert(StringIn(target, "99"), qt.Equals, false)
}

func TestReverse(t *testing.T) {
	c := qt.New(t)
	c.Helper()

	c.Assert(Reverse("123456"), qt.Equals, "654321")
}
