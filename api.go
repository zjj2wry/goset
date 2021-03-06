/*
Copyright 2017 Jim Zhang (jim.zoumo@gmail.com)

Licensed under the Apache License, Version 2.0 (the "License");
you may not use this file except in compliance with the License.
You may obtain a copy of the License at

    http://www.apache.org/licenses/LICENSE-2.0

Unless required by applicable law or agreed to in writing, software
distributed under the License is distributed on an "AS IS" BASIS,
WITHOUT WARRANTIES OR CONDITIONS OF ANY KIND, either express or implied.
See the License for the specific language governing permissions and
limitations under the License.
*/

package goset

// Set provides a collection of oprations for sets
//
// The implementation of Set is base on hash table.
// So the elements must be hashable. functions, maps,
// slices are unhashable type, adding these elements
// will cause panic.
//
// There are two implementations of Set:
// 1. default is unsafe based on hash table
// 2. thread safe based on sync.RWMutex (maybe change to sync.Map in GO1.9)
//
// The two kinds of sets can easily convert to the
// other one. But you must know exactly what you are doing
// to avoid the concurrent race
type Set interface {
	// Add adds an element to the set.
	// If the elem is already exists, it will not return an ErrAlreadyExisted by default.
	// You can set package level var RaiseErrAlreadyExisted=true to raise it.
	Add(elem interface{}) error

	// Extend adds all elements in the given interface b to this set
	// the given interface must be array, slice or set.
	// Extend will ignore elements already existed error.
	Extend(b interface{}) error

	// Remove removes an element from the set.
	Remove(elem interface{})

	// Clear removes all elevemnts from the set.
	Clear()

	// Copy clones the set.
	Copy() Set

	// Len returns the size of set. aka Cardinality.
	Len() int

	// Elements returns all elements in this set.
	Elements() []interface{}

	// Contains checks whether the given item is in the set.
	Contains(item interface{}) bool

	// Equal checks whether this set is equal to the given one.
	// There are two constraints if set a is equal to set b.
	// the two set must have the same size and contain the same elements.
	Equal(b Set) bool

	// IsSubsetOf checks whether this set is the subset of the given set
	// In other words, all elements in this set are also the elements
	// of the given set.
	IsSubsetOf(b Set) bool

	// IsSupersetOf checks whether this set is the superset of the given set
	// In other words, all elements in the given set are also the elements
	// of this set.
	IsSupersetOf(b Set) bool

	// String returns the string representation of the set.
	String() string

	// ToThreadUnsafe returns a thread unsafe set.
	// Carefully use the method.
	ToThreadUnsafe() Set

	// ToThreadSafe returns a thread safe set.
	// Carefully use the method.
	ToThreadSafe() Set

	// Range calls f sequentially for each element present in the set.
	// If f returns false, range stops the iteration.
	//
	// Note: the iteration order is not specified and is not guaranteed
	// to be the same from one iteration to the next. The index only
	// means how many elements have been visited in the iteration, it not
	// specifies the index of an element in the set
	Range(foreach func(index int, elem interface{}) bool)

	// ---------------------------------------------------------------------
	// Set Oprations

	// Diff returns the difference between the set and this given
	// one, aka Difference Set
	// math formula: a - b
	Diff(b Set) Set

	// SymmetricDiff returns the symmetric difference between this set
	// and the given one. aka Symmetric Difference Set
	// math formula: (a - b) ∪ (b - a)
	SymmetricDiff(b Set) Set

	// Unite combines two sets into a new one, aka Union Set
	// math formula: a ∪ b
	Unite(b Set) Set

	// Intersect returns the intersection of two set, aka Intersection Set
	// math formula: a ∩ b
	Intersect(b Set) Set
}

// NewSetFrom returns a new Set from the given collection
// the collection must be array, slice or Set,
// otherwise it will panic
func NewSetFrom(i interface{}) Set {
	s := newSet()
	err := s.Extend(i)
	if err != nil {
		panic(err)
	}
	return s
}

// NewSet returns a new Set which contains the
// given elements
func NewSet(elems ...interface{}) Set {
	return NewSetFrom(elems)
}

// NewSetFromInts returns a new Set containing
// all the elements in the int slice
func NewSetFromInts(e []int) Set {
	return NewSetFrom(e)
}

// NewSetFromStrings returns a new Set containing
// all the elements in the string slice
func NewSetFromStrings(e []string) Set {
	return NewSetFrom(e)
}

// NewSetFromFloats returns a new Set containing
// all the elements in the float64 slice
func NewSetFromFloats(e []float64) Set {
	return NewSetFrom(e)
}

// NewSafeSet returns a new thread-safe Set
// which contains the given elements
func NewSafeSet(elems ...interface{}) Set {
	return NewSafeSetFrom(elems)
}

// NewSafeSetFrom returns a new thread-safe Set
// from the given collection. The collection must be
// array, slice or Set, otherwise it will panic.
func NewSafeSetFrom(i interface{}) Set {
	s := newThreadSafeSet()
	err := s.Extend(i)
	if err != nil {
		panic(err)
	}
	return s
}

// NewSafeSetFromInts returns a new thread-safe Set containing
// all the elements in the int slice
func NewSafeSetFromInts(e []int) Set {
	return NewSafeSetFrom(e)
}

// NewSafeSetFromStrings returns a new thread-safe Set containing
// all the elements in the string slice
func NewSafeSetFromStrings(e []string) Set {
	return NewSafeSetFrom(e)
}

// NewSafeSetFromFloats returns a new thread-safe Set containing
// all the elements in the float64 slice
func NewSafeSetFromFloats(e []float64) Set {
	return NewSafeSetFrom(e)
}
