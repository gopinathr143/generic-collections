package collections

import "sort"

type Number interface {
	int64 | float64 | int32 | int16 | int8 | int | float32
}

// GenericCollection represents a collection of generic items.
type GenericCollection struct {
	items []interface{}
}

// NewGenericCollection creates a new GenericCollection with the specified items.
func NewGenericCollection(items ...interface{}) *GenericCollection {
	return &GenericCollection{items: items}
}

// Count returns the number of items in the collection.
func (c *GenericCollection) Count() int {
	return len(c.items)
}

// Add adds a single item to the collection.
func (c *GenericCollection) Add(item interface{}) {
	c.items = append(c.items, item)
}

// AddRange adds multiple items to the collection.
func (c *GenericCollection) AddRange(items ...interface{}) {
	c.items = append(c.items, items...)
}

// Remove removes a single item from the collection.
func (c *GenericCollection) Remove(item interface{}) bool {
	for i, v := range c.items {
		if v == item {
			c.items = append(c.items[:i], c.items[i+1:]...)
			return true
		}
	}
	return false
}

// RemoveAt removes an item at the specified index from the collection.
func (c *GenericCollection) RemoveAt(index int) {
	c.items = append(c.items[:index], c.items[index+1:]...)
}

// RemoveRange removes multiple items from the collection.
func (c *GenericCollection) RemoveRange(index int, count int) {
	c.items = append(c.items[:index], c.items[index+count:]...)
}

// Contains returns true if the collection contains the specified item, false otherwise.
func (c *GenericCollection) Contains(item interface{}) bool {
	for _, v := range c.items {
		if v == item {
			return true
		}
	}
	return false
}

// IndexOf returns the index of the specified item in the collection, or -1 if it is not found.
func (c *GenericCollection) IndexOf(item interface{}) int {
	for i, v := range c.items {
		if v == item {
			return i
		}
	}
	return -1
}

// Find returns the first item in the collection that satisfies the specified predicate function,
// or nil if no such item is found.
func (c *GenericCollection) Find(predicate func(interface{}) bool) interface{} {
	for _, v := range c.items {
		if predicate(v) {
			return v
		}
	}
	return nil
}

// FindMany returns all items in the collection that satisfy the specified predicate function,
// or an empty slice if no such item is found.
func (c *GenericCollection) FindMany(predicate func(interface{}) bool) []interface{} {
	var result []interface{}
	for _, v := range c.items {
		if predicate(v) {
			result = append(result, v)
		}
	}
	return result
}

// Select returns a new collection containing the results of applying the specified selector function to each item in the collection.
func (c *GenericCollection) Select(selector func(interface{}) interface{}) *GenericCollection {
	result := make([]interface{}, len(c.items))
	for i, v := range c.items {
		result[i] = selector(v)
	}
	return NewGenericCollection(result...)
}

// Where applies a filtering operation on the items in the collection by applying a given predicate function. It returns a new GenericCollection containing only the items that satisfy the predicate.
func (gc *GenericCollection) Where(predicate func(interface{}) bool) *GenericCollection {
	result := NewGenericCollection()
	for _, item := range gc.items {
		if predicate(item) {
			result.Add(item)
		}
	}
	return result
}

// GroupBy groups the items in the collection based on a key obtained by applying a given keySelector function to each item. It returns a map where each key corresponds to a GenericCollection containing the items with that key.
func (gc *GenericCollection) GroupBy(keySelector func(interface{}) interface{}) map[interface{}]*GenericCollection {
	groups := make(map[interface{}]*GenericCollection)
	for _, item := range gc.items {
		key := keySelector(item)
		if _, ok := groups[key]; !ok {
			groups[key] = NewGenericCollection()
		}
		groups[key].Add(item)
	}
	return groups
}

// Sum calculates the sum of a numerical property of each item in the collection, as specified by the given selector function.
func (gc *GenericCollection) Sum(selector func(interface{}) float64) float64 {
	sum := 0.0
	for _, item := range gc.items {
		sum = sum + selector(item)
	}
	return sum
}

// Max returns the maximum value of a numerical property of each item in the collection, as specified by the given selector function.
func (gc *GenericCollection) Max(selector func(interface{}) float64) float64 {
	if len(gc.items) == 0 {
		return 0
	}
	max := selector(gc.items[0])
	for i := 1; i < len(gc.items); i++ {
		if current := selector(gc.items[i]); current > max {
			max = current
		}
	}
	return max
}

// Min returns the minimum value of a numerical property of each item in the collection, as specified by the given selector function.
func (gc *GenericCollection) Min(selector func(interface{}) float64) float64 {
	if len(gc.items) == 0 {
		return 0
	}
	min := selector(gc.items[0])
	for i := 1; i < len(gc.items); i++ {
		if current := selector(gc.items[i]); current < min {
			min = current
		}
	}
	return min
}

// Sort sorts the items in the collection based on a given comparison function that takes two items as arguments and returns a boolean indicating whether the first item should come before the second item in the sorted result.
func (gc *GenericCollection) Sort(less func(i, j interface{}) bool) {
	sort.Slice(gc.items, func(i, j int) bool {
		return less(gc.items[i], gc.items[j])
	})
}

// Reverse reverses the order of the items in the collection
func (gc *GenericCollection) Reverse() {
	for i, j := 0, len(gc.items)-1; i < j; i, j = i+1, j-1 {
		gc.items[i], gc.items[j] = gc.items[j], gc.items[i]
	}
}

// FindByIndex returns the item in the collection at a given index. If the index is out of range, it returns nil.
func (gc *GenericCollection) FindByIndex(index int) interface{} {
	if index < 0 || index >= len(gc.items) {
		return nil
	}
	return gc.items[index]
}

// ReverseSort is a method of the GenericCollection struct that sorts the items in the collection in descending order.
func (gc *GenericCollection) ReverseSort(less func(i, j int) bool) {
	sort.SliceStable(gc.items, func(i, j int) bool {
		return less(j, i)
	})
}

// GetAll returns all the items in the collection.
func (gc *GenericCollection) GetAll() []interface{} {
	return gc.items
}

func (gc *GenericCollection) indexOf(item interface{}) int {
	for i, val := range gc.items {
		if val == item {
			return i
		}
	}
	return -1
}
