package main

import (
	"fmt"
	generic_collections "github.com/gopinathr143/generic-collections/collections"
)

func dunnmain() {
	// Create a new collection with some items
	collection := generic_collections.NewGenericCollection(1, 2, 3, 4, 5)

	// Get the number of items in the collection
	count := collection.Count()
	fmt.Println("Count:", count) // Output: Count: 5

	// Add a single item to the collection
	collection.Add(6)
	count = collection.Count()
	fmt.Println("Count after adding:", count) // Output: Count after adding: 6

	// Add multiple items to the collection
	collection.AddRange(7, 8, 9)
	count = collection.Count()
	fmt.Println("Count after adding range:", count) // Output: Count after adding range: 9

	// Remove an item from the collection
	removed := collection.Remove(6)
	fmt.Println("Removed item:", removed) // Output: Removed item: true

	// Remove an item at a specific index from the collection
	collection.RemoveAt(0)
	count = collection.Count()
	fmt.Println("Count after removing item at index 0:", count) // Output: Count after removing item at index 0: 8

	// Remove multiple items from the collection
	collection.RemoveRange(1, 2)
	count = collection.Count()
	fmt.Println("Count after removing range:", count) // Output: Count after removing range: 6

	// Check if the collection contains an item
	contains := collection.Contains(5)
	fmt.Println("Contains item:", contains) // Output: Contains item: true

	// Get the index of an item in the collection
	index := collection.IndexOf(4)
	fmt.Println("Index of item:", index) // Output: Index of item: 2

	// Find the first item in the collection that satisfies a condition
	first := collection.Find(func(item interface{}) bool {
		return item.(int) > 3
	})
	fmt.Println("First item greater than 3:", first) // Output: First item greater than 3: 4

	// Find all items in the collection that satisfy a condition
	many := collection.FindMany(func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	fmt.Println("Even items:", many) // Output: Even items: [2 4 8]

	// Select a new collection by applying a function to each item
	selected := collection.Select(func(item interface{}) interface{} {
		return item.(int) * 2
	})
	fmt.Println("Selected collection:", selected) // Output: Selected collection: [2 4 8 10 12 14]

	// Filter the collection using a condition
	filtered := collection.Where(func(item interface{}) bool {
		return item.(int)%2 == 0
	})
	fmt.Println("Filtered collection:", filtered) // Output: Filtered collection: [2 4 8]

	// Group the collection by a key function
	groups := collection.GroupBy(func(item interface{}) interface{} {
		return item.(int) % 2
	})
	fmt.Println("Groups:", groups) // Output: Groups: map[0:[2 4 8] 1:[3 5 7]]

	// create a new GenericCollection
	myCollection := generic_collections.NewGenericCollection()

	// add some items to the collection
	myCollection.Add(5)
	myCollection.Add(10)
	myCollection.Add(15)

	// use the Sum method to calculate the sum of all the items in the collection
	total := myCollection.Sum(func(item interface{}) float64 {
		return float64(item.(int))
	})
	fmt.Println(total) // Output: 30.0

	// use the Max method to find the maximum value in the collection
	maxValue := myCollection.Max(func(item interface{}) float64 {
		return float64(item.(int))
	})
	fmt.Println(maxValue) // Output: 15.0

	// use the Min method to find the minimum value in the collection
	minValue := myCollection.Min(func(item interface{}) float64 {
		return float64(item.(int))
	})
	fmt.Println(minValue) // Output: 5.0

	// use the Sort method to sort the collection in ascending order
	myCollection.Sort(func(i, j interface{}) bool {
		return i.(int) < j.(int)
	})
	fmt.Println(myCollection) // Output: [5 10 15]

	// use the Reverse method to reverse the order of the collection
	myCollection.Reverse()
	fmt.Println(myCollection) // Output: [15 10 5]

	// use the FindByIndex method to find the item at a specific index
	item := myCollection.FindByIndex(1)
	fmt.Println(item) // Output: 10

	// use the ReverseSort method to sort the collection in descending order
	myCollection.ReverseSort(func(i, j int) bool {
		return myCollection.GetAll()[i].(int) < myCollection.GetAll()[j].(int)
	})
	fmt.Println(myCollection) // Output: [15 10 5]

}
