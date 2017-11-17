//Todo(mhc)
//Go HashSet implemented via map
package set

import (
	"bytes"
	"fmt"
)

//interface
type Set interface {
	Add(e interface{}) bool
	Remove(e interface{})
	Clear()
	Contains(e interface{}) bool
	Len() int
	Same(other Set) bool
	Elements() []interface{}
	String() string
	Copy() Set
}

//Super function----------------------------------------------
//Determine whether one is a super set of other
func IsSuperset(one Set, other Set) bool {
	if one == nil || other == nil {
		return false
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || oneLen == otherLen {
		return false
	}
	if oneLen > 0 && otherLen == 0{
		return true
	}
	for _, v := range other.Elements() {
		if !one.Contains(v){
			return false
		}
	}
	return true
}

//Return union of two Set
func Union(one Set, other Set) Set {
	if other == nil && one == nil {
		return nil
	}
	if other == nil {
		return one.Copy()
	}
	if one == nil {
		return other.Copy()
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0{
		return other.Copy()
	}
	if otherLen == 0{
		return one.Copy()
	}
	copyset := other.Copy()
	for _, key := range one.Elements() {
		copyset.Add(key)
	}
	return copyset
}

//Return Intersect of two Set
func Intersect(one Set, other Set) Set {
	if one == nil || other == nil {
		return nil
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || otherLen == 0{
		return nil
	}
	if oneLen >= otherLen {
		copyset := other.Copy()
		for _, key := range copyset.Elements(){
			if !one.Contains(key){
				copyset.Remove(key)
			}
		}
		return copyset
	}else{
		copyset := one.Copy()
		for _, key := range copyset.Elements(){
			if !other.Contains(key){
				copyset.Remove(key)
			}
		}
		return copyset
	}
}

//Return Difference of Set compared to other
func Difference(one Set, other Set) Set {
	if other == nil && one == nil {
		return nil
	}
	if other == nil {
		return one.Copy()
	}
	if one == nil {
		return nil
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 || otherLen == 0{
		return one.Copy()
	}
	copyset := one.Copy()
	for _, key := range copyset.Elements() {
		if other.Contains(key){
			copyset.Remove(key)
		}
	}
	return copyset
}

//Return Summetric Difference of two Set
func SummetricDifference(one Set, other Set) Set {
	if other == nil && one == nil {
		return nil
	}
	if other == nil {
		return one.Copy()
	}
	if one == nil {
		return other.Copy()
	}
	oneLen := one.Len()
	otherLen := other.Len()
	if oneLen == 0 {
		return other.Copy()
	}
	if otherLen == 0{
		return one.Copy()
	}
	union := Union(one, other)
	intersect := Intersect(one, other)
	return Difference(union, intersect)
}

//HashSet implement Set-----------------------------
type HashSet struct {
	m map[interface{}]bool
}

//Generate new HashSet
func NewHashSet() *HashSet {
	return &HashSet{m: make(map[interface{}]bool)}
}

//Add an element
//If contain, return false else true
func (set *HashSet) Add(e interface{}) bool {
	if !set.m[e] {
		set.m[e] = true
		return true
	}
	return false
}

//Remove an element
func (set *HashSet) Remove(e interface{}) {
	delete(set.m, e)
}

//Clear all elements
func (set *HashSet) Clear() {
	set.m = make(map[interface{}]bool)
}

//Determine whether the element is included in the HashSet
func (set *HashSet) Contains(e interface{}) bool {
	return set.m[e]
}

//Return number of elements in HashSet
func (set *HashSet) Len() int {
	return len(set.m)
}

//Determine whether the HashSet is same to other
func (set *HashSet) Same(other Set) bool {
	if other == nil {
		return false
	}
	if set.Len() != other.Len() {
		return false
	}
	for key := range set.m {
		if !other.Contains(key) {
			return false
		}
	}
	return true
}

//Return an iterable slice including all elements
func (set *HashSet) Elements() []interface{} {
	initialLen := len(set.m)
	snapshot := make([]interface{}, initialLen)
	actualLen := 0
	for key := range set.m {
		if actualLen < initialLen {
			snapshot[actualLen] = key
		} else {
			snapshot = append(snapshot, key)
		}
		actualLen++
	}
	if actualLen < initialLen {
		snapshot = snapshot[:actualLen]
	}
	return snapshot
}

//Return a fromat string including all elements
func (set *HashSet) String() string {
	var buf bytes.Buffer
	buf.WriteString("Set{")
	first := true
	for key := range set.m {
		if first {
			first = false
		} else {
			buf.WriteString(" ")
		}
		buf.WriteString(fmt.Sprintf("%v", key))
	}
	buf.WriteString("}")
	return buf.String()
}

//Return a copy HashSet
func (set *HashSet) Copy() Set {
	copySet := NewHashSet()
	for key := range set.m {
		copySet.Add(key)
	}
	return copySet
}