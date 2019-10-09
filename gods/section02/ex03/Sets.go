// Copyright 2019 readailib.com. All rights reserved.
// Use of this source code is governed by a BSD-style
// license that can be found in the LICENSE file.

package set

// A Set is a linear data structure that has a collection of values that are not repeated.
// A set can store unique values without any particular order.
// In the real world, sets can be used to collect all tags
// for blog posts and conversation participants in a chat.
// The data can be of Boolean, integer, float, characters,
// and other types. Static sets allow only query operations,
// which means operations related to querying the elements.
// Dynamic and mutable sets allow for the insertion and deletion of elements.
// Algebraic operations such as union, intersection,
// difference, and subset can be defined on the sets.
// The following example shows the Set integer with a map integer key and bool as a value:

//Set class
type Set struct {
	integerMap map[int]bool
}

// New method in Set
func (s *Set) New() {
	s.integerMap = make(map[int]bool)
}

// The AddElement, DeleteElement, ContainsElement,
func (s *Set) AddElement(property int) {
	if !s.ContainsElement(property) {
		s.integerMap[property] = true
	}
}

func (s *Set) DeleteElement(property int) {
	delete(s.integerMap, property)
}

func (s *Set) ContainsElement(property int) bool {
	_, exist := s.integerMap[property]
	return exist
}

//Intersect, Union, and main methods are discussed in the following sections.
//the InterSect method on the Set class returns an intersectionSet that consists
//of the intersection of set and anotherSet.
//The set class is traversed through integerMap and checked against another
//Set to see if any elements exist:
func (s *Set) Intersect(otherSet *Set) *Set {
	var intersect = &Set{}
	intersect.New()
	for value := range s.integerMap {
		if otherSet.ContainsElement(value) {
			intersect.AddElement(value)
		}
	}
	return intersect
}

func (s *Set) Union(otherSet *Set) *Set {
	var unionSet = &Set{}
	unionSet.New()
	for value := range s.integerMap {
		unionSet.AddElement(value)
	}

	for value := range otherSet.integerMap {
		unionSet.AddElement(value)
	}
	return unionSet
}
