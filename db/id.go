package db

import (
	"sort"
)

// ids is a sorted list of unique identifiers.
type ids []int

func (s ids) Len() int           { return len(s) }
func (s ids) Swap(i, j int)      { s[i], s[j] = s[j], s[i] }
func (s ids) Less(i, j int) bool { return s[i] < s[j] }

// insert inserts an id in sorted order and returns a new list of ids.
// If the id already exists then it returns the original list.
func (s ids) insert(id int) ids {
	index := sort.SearchInts([]int(s), id)
	if index == len(s) {
		s = append(s, id)
	} else if s[index] != id {
		s = append(s, 0)
		copy(s[index+1:], s[index:])
		s[index] = id
	}
	return s
}

// remove removes the id from the list if it exists and returns a new list of ids.
func (s ids) remove(id int) ids {
	index := sort.SearchInts([]int(s), id)
	if index < len(s) && s[index] == id {
		s = append(s[:index], s[index+1:]...)
	}
	return s
}
