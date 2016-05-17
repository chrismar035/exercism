// Package stringset implementas a Set as a collection of unique string values.
//
// API:
//
// New() Set
// NewFromSlice([]string) Set
// (s Set) Add(string)         // modify s
// (s Set) Delete(string)      // modify s
// (s Set) Has(string) bool
// (s Set) IsEmpty() bool
// (s Set) Len() int
// (s Set) Slice() []string
// (s Set) String() string
// Equal(s1, s2 Set) bool
// Subset(s1, s2 Set) bool     // return s1 ⊆ s2
// Disjoint(s1, s2 Set) bool
// Intersection(s1, s2 Set) Set
// Union(s1, s2 Set) Set
// Difference(s1, s2 Set) Set  // return s1 ∖ s2
// SymmetricDifference(s1, s2 Set) Set
//
// For Set.String, use '{' and '}', output elements as double-quoted strings
// safely escaped with Go syntax, and use a comma and a single space between
// elements.  For example {"a", "b"}.
// Format the empty set as {}.
package stringset

import (
	"bytes"
	"crypto/sha1"
	"encoding/base64"
	"strings"
)

// Set represents a collection of unique strings
type Set struct {
	items map[string]string
}

// New creates a new empty Set
func New() Set {
	return Set{items: make(map[string]string)}
}

// NewFromSlice create a new Set with the elements from the slice
func NewFromSlice(initial []string) Set {
	set := New()
	for _, val := range initial {
		set.Add(val)
	}

	return set
}

// Add will insert the provided string into the Set
func (s Set) Add(add string) {
	s.items[shaOfString(add)] = add
}

// Delete removes the provided string from the Set
func (s Set) Delete(del string) {
	delete(s.items, shaOfString(del))
}

// Has returns if the Set has the string already
func (s Set) Has(test string) bool {
	_, ok := s.items[shaOfString(test)]
	return ok
}

// IsEmpty returns true if the Set has no elements
func (s Set) IsEmpty() bool {
	return len(s.items) == 0
}

// Len returns the length of the Set
func (s Set) Len() int {
	return len(s.items)
}

// Slice returns the Set converted to a slice of strings
func (s Set) Slice() []string {
	var items []string

	for _, value := range s.items {
		items = append(items, value)
	}
	return items
}

func (s Set) String() string {
	if s.IsEmpty() {
		return "{}"
	}

	var buffer bytes.Buffer

	_, _ = buffer.Write([]byte("{\""))
	_, _ = buffer.Write([]byte(strings.Join(s.Slice(), "\", \"")))
	_, _ = buffer.Write([]byte("\"}"))

	return buffer.String()
}

// Equal returns true if both given Sets contain exactly the same strings
func Equal(s1, s2 Set) bool {
	for _, val := range s1.items {
		if !s2.Has(val) {
			return false
		}
	}
	for _, val := range s2.items {
		if !s1.Has(val) {
			return false
		}
	}

	return true
}

// Subset returns true if all the elements of s1 are in s2
func Subset(s1, s2 Set) bool {
	for _, val := range s1.items {
		if !s2.Has(val) {
			return false
		}
	}

	return true
}

// Disjoint returns true if none of the elements of s1 are in s2
func Disjoint(s1, s2 Set) bool {
	for _, val := range s1.items {
		if s2.Has(val) {
			return false
		}
	}

	return true
}

// Union returns the set with all of the elements from s1 and s2 combined
func Union(s1, s2 Set) Set {
	var new = New()
	for _, val := range s1.items {
		new.Add(val)
	}
	for _, val := range s2.items {
		new.Add(val)
	}

	return new
}

// Intersection returns a new Set with the items from s1 which also occur in s2
func Intersection(s1, s2 Set) Set {
	var intersection = New()

	for _, val := range s1.items {
		if s2.Has(val) {
			intersection.Add(val)
		}
	}

	return intersection
}

// Difference returns a new Set with the items from s2 which are not in s1
func Difference(s1, s2 Set) Set {
	var difference = New()

	for _, val := range s1.items {
		if !s2.Has(val) {
			difference.Add(val)
		}
	}

	return difference
}

// SymmetricDifference returns a new Set with the items from both s1 and s2
// which are not in the other Set respectively
func SymmetricDifference(s1, s2 Set) Set {
	var difference = New()

	for _, val := range s2.items {
		if !s1.Has(val) {
			difference.Add(val)
		}
	}

	for _, val := range s1.items {
		if !s2.Has(val) {
			difference.Add(val)
		}
	}

	return difference
}

func shaOfString(s string) string {
	hasher := sha1.New()
	_, _ = hasher.Write([]byte(s))
	return base64.URLEncoding.EncodeToString(hasher.Sum(nil))
}
