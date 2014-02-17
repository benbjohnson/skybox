package db

import (
	_assert "github.com/stretchr/testify/assert"
	"sort"
	"testing"
)

// Ensure that ids can be sorted.
func TestIdsSort(t *testing.T) {
	s := ids{7, 2, 0, 5}
	sort.Sort(s)
	_assert.Equal(t, s, ids{0, 2, 5, 7})
}

// Ensure that ids are unique.
func TestIdsInsert(t *testing.T) {
	var ids ids
	ids = ids.insert(5)
	ids = ids.insert(2)
	ids = ids.insert(7)
	ids = ids.insert(5)
	ids = ids.insert(0)
	_assert.Equal(t, []int(ids), []int{0, 2, 5, 7})
}

// Ensure that ids can be removed.
func TestIdsRemove(t *testing.T) {
	var ids = ids{0, 2, 5, 7}
	ids = ids.remove(2)
	ids = ids.remove(7)
	ids = ids.remove(6)
	ids = ids.remove(5)
	ids = ids.remove(0)
	_assert.Equal(t, len(ids), 0)
}
